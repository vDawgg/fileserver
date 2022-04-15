package backend;

import com.google.protobuf.ByteString;
import io.grpc.Server;
import io.grpc.ServerBuilder;
import io.grpc.stub.StreamObserver;

import java.io.ByteArrayInputStream;
import java.io.InputStream;
import java.util.*;
import java.util.logging.Level;
import java.util.logging.Logger;

import io.minio.*;
import io.minio.messages.Item;

import static java.lang.System.getenv;

public class Retriever {

    //TODO: Look into how to set up health checks

    //TODO: Look at java grpc authentication example (uses tokens, so might be interesting)
    // https://github.com/grpc/grpc-java/tree/master/examples/example-jwt-auth

    private static final Logger logger = Logger.getLogger("Retriever");

    private final int port;
    private final Server server;

    private static MinioClient minioClient;

    //Stub for testing
    public Retriever(ServerBuilder<?> serverBuilder, int port) {
        this.port = port;
        this.server = serverBuilder
                .addService(new RetrieverImpl())
                .build();
    }

    void start() {
        int port = Integer.parseInt(getenv().getOrDefault("PORT", "9390"));

        //TODO: Set up connection to minio

        try {
            server.start();
            minioClient = MinioClient.builder()
                    .endpoint("http://127.0.0.1:9000") //Environment variables?
                    .credentials("minioadmin", "minioadmin") //Should be gotten via environment variables
                    .build();
        } catch (Exception e) {
            e.printStackTrace();
        }
        Runtime.getRuntime()
                .addShutdownHook(
                        new Thread(
                                () -> {
                                    System.err.println("*** Shutting down since JVM is shutting down");
                                    Retriever.this.stop();
                                    System.err.println("*** Server shutting down");
                                }
                        )
                );
    }

    //Does the healthmanager need to be stopped here?
    void stop() {
        if (server!=null) {
            server.shutdown();
        }
    }

    //TODO: Rename this for something better
    public static String getTypeFromPathName(String pathName) {

        String[] s = pathName.split("/");

        if(!s[0].contains(".")) return "directory";

        String[] result = s[0].split("\\.");

        return result[1];
    }


    static class RetrieverImpl extends RetrieverGrpc.RetrieverImplBase {

        //TODO: Implement functionality for sending multiple files
        @Override
        public StreamObserver<Chunk> saveFiles(StreamObserver<UploadStatus> responseObserver) {
            return new StreamObserver<Chunk>() {
                String filename;
                String bucket;
                ByteString bs;

                @Override
                public void onNext(Chunk value) {
                    if(filename==null & bucket==null) {
                        filename = value.getFileDescription().getFileName();
                        bucket = value.getFileDescription().getBucket();
                        logger.log(Level.INFO, "Receiving file with name: "+filename);
                    }
                    if(bs==null) {
                        bs = (ByteString) value.getContent();
                    } else {
                        bs = bs.concat((ByteString) value.getContent());
                    }
                }

                @Override
                public void onError(Throwable t) {
                    responseObserver.onNext(UploadStatus.newBuilder() //Is this the right place??
                                    .setCodeValue(2)
                                    .build());
                    logger.log(Level.WARNING, "An Error occurred while trying to save a File", t);
                }

                @Override
                public void onCompleted() {
                    //TODO: Add the file to the db
                    responseObserver.onNext(UploadStatus.newBuilder()
                                    .setCodeValue(1)
                                    .build());
                    //TODO: Delete file after saving it to the db to avoid write errors
                    try {
                        //File file = new File(filename);
                        //logger.log(Level.INFO, ""+bs.toByteArray().length);

                        if (!minioClient.bucketExists(BucketExistsArgs.builder().bucket(bucket).build())) {
                            minioClient.makeBucket(MakeBucketArgs.builder().bucket(bucket).build());
                            logger.log(Level.INFO, "Created bucket: "+bucket);
                        }

                        Map<String, String> map = Map.of("filename", filename); //Is this even necessary?

                        //TODO: Think of meaningful headers and maybe use user metadata
                        minioClient.putObject(
                                PutObjectArgs.builder()
                                        .bucket(bucket)
                                        .object(filename)
                                        .headers(map)
                                        .stream(new ByteArrayInputStream(bs.toByteArray()), bs.size(), -1) //What does partsize do?
                                        .build()
                        );
                        logger.log(Level.INFO, "Added file: "+filename);
                    } catch (Exception e) {
                        logger.log(Level.WARNING, "An error occured while trying to create a file"+e);
                    }
                }
            };
        }

        //frontend should always start in home directory!
        @Override
        public void getStructure(StructureRequest request, StreamObserver<Structure> responseObserver) {
            String bucket = request.getBucket();
            String directory = request.getDirectory();
            Structure.Builder structure = Structure.newBuilder();

            Iterable<Result<Item>> it = minioClient.listObjects(ListObjectsArgs.builder()
                            .prefix(directory)
                            .bucket(bucket)
                            .build());

            try {
                for (Result<Item> result : it) {
                    String s = result.get().objectName();
                    logger.log(Level.INFO, s);
                    if(s.startsWith(directory)) {
                        String subSequence = (String) s.subSequence(directory.length(), s.length());
                        structure.addObject(Object.newBuilder()
                                        .setName(subSequence)
                                        .setType(getTypeFromPathName(subSequence))
                                        .build());
                    }
                }
                responseObserver.onNext(structure.build());
            } catch (Exception e) {
                logger.log(Level.WARNING, "Unable to retrieve filestructure for bucket: "+bucket);
            }
        }

        @Override
        public void getFiles(DownloadRequest request, StreamObserver<Chunk> responseObserver) {

            String bucket = request.getFileDescription().getBucket();
            String filename = request.getFileDescription().getFileName();

            try {
                InputStream stream = minioClient.getObject(
                        GetObjectArgs.builder()
                                .bucket(bucket)
                                .object(filename)
                                .build());

                ByteString bs = ByteString.readFrom(stream);

                int i = 0;
                int j = 1000;

                while(j<bs.size()) {
                    ByteString b = bs.substring(i, j);
                    Chunk chunk = Chunk.newBuilder()
                            .setFileDescription(
                                    FileDescription.newBuilder()
                                            .setFileName(filename)
                                            .build())
                            .setContent(b)
                            .build();
                    responseObserver.onNext(chunk);
                    i = j;
                    j += 1000;
                }

                //TODO: Should be able to make this code look a lot nicer!
                ByteString b = bs.substring(i, bs.size());
                Chunk chunk = Chunk.newBuilder()
                        .setFileDescription(
                                FileDescription.newBuilder()
                                        .setFileName(filename)
                                        .build())
                        .setContent(b)
                        .build();

                responseObserver.onNext(chunk);

                responseObserver.onCompleted();

                logger.log(Level.INFO, "Sent file from bucket: "+bucket);

            } catch (Exception e) {
                logger.log(Level.WARNING, "Could not retrieve file of name: "+filename+", "+ Arrays.toString(e.getStackTrace())); //Is this a good idea for privacy???
                e.printStackTrace();
            }
        }
    }

    public static void main(String[] args) {

    }
}

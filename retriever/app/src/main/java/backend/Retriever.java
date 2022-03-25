package backend;

import com.google.protobuf.ByteString;
import io.grpc.Server;
import io.grpc.ServerBuilder;
import io.grpc.stub.StreamObserver;

import java.io.ByteArrayInputStream;
import java.io.File;
import java.io.FileOutputStream;
import java.util.logging.Level;
import java.util.logging.Logger;

import io.minio.*;

import static java.lang.System.getenv;

public class Retriever {

    //TODO: Look into how to set up health checks

    //TODO: Look at java grpc authentication example (uses tokens, so might be interesting)
    // https://github.com/grpc/grpc-java/tree/master/examples/example-jwt-auth
    // if feeling like it google auth can be added as well:
    // https://github.com/grpc/grpc-java/tree/master/examples/example-gauth

    private static Logger logger = Logger.getLogger("Retriever");

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


    static class RetrieverImpl extends RetrieverGrpc.RetrieverImplBase {

        //TODO: Implement functionality for sending multiple files
        @Override
        public StreamObserver<Chunk> saveFiles(StreamObserver<UploadStatus> responseObserver) {
            return new StreamObserver<Chunk>() {
                String filename;
                String directory;
                User user; //TODO: Implement users
                ByteString bs;

                @Override
                public void onNext(Chunk value) {
                    if(filename==null & directory==null) {
                        filename = value.getFileDescription().getFileName();
                        directory = value.getFileDescription().getDirectory();
                        logger.log(Level.INFO, "Receiving file with name: "+filename);
                    }
                    if(bs==null) {
                        bs = (ByteString) value.getContent();
                    } else {
                        bs = bs.concat((ByteString) value.getContent()); //Should work, but definitely needs to be checked!
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

                        if (!minioClient.bucketExists(BucketExistsArgs.builder().bucket(directory).build())) {
                            minioClient.makeBucket(MakeBucketArgs.builder().bucket(directory).build());
                            logger.log(Level.INFO, "Created bucket: "+directory);
                        }

                        File file = new File(filename);

                        FileOutputStream fs = new FileOutputStream(filename);
                        fs.write(bs.toByteArray());

                        //TODO: Think of meaningful headers and maybe use user metadata
                        minioClient.putObject(
                                PutObjectArgs.builder()
                                        .bucket(directory)
                                        .object(filename)
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

        //Probably needs to be changed if thumbnails are to be shown
        @Override
        public void sendStructure(StructureRequest request, StreamObserver<Structure> responseObserver) {

        }

        @Override
        public void getFiles(DownloadRequest request, StreamObserver<Chunk> responseObserver) {

        }

        @Override
        public void authenticate(User user, StreamObserver<AuthenticationStatus> responseObserver) {

        }

    }

    public static void main(String[] args) {

    }
}

Services:

- frontent -> interprets filetypes and directories; can send files over grpc -> react
- retriever -> retrieves and adds files to object storage (minio), also has to be able to add metadata (maybe seperate db if minio doesnt support), gets requests from frontend via grpc -> Java, gradle
- authenticator -> authenticates user using mongodb for storage and retrieves information on what files to show him (is supported by minio add metadata if not use same mongodb as in retriever) -> Use authentication to retrieve ID with which the file content can be fetched (doesnt allow for authentication in minio though (minio has the possibility to lock every bucket with authentication)) -> Maybe get userid to see all the buckets associated with that id and once the user tries to get information from that bucket get the according passwords from the auth server (with userId and password)

Notes:
- According to website that tries filetransfer over grpc, grpc takes twice as long as normal http2 to transfer files -> very big impact; does it make any sense to do this in grpc?, or should grpc just be used for authentication?
- use docker compose for container orchestration

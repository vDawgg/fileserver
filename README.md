Services:

- frontent -> interprets filetypes and directories; can send files over grpc -> react
- retriever -> retrieves and adds files to object storage (minio), also has to be able to add metadata (maybe seperate db if minio doesnt support), gets requests from frontend via grpc -> Java, gradle
- authenticator -> authenticates user using mongodb for storage and retrieves information on what files to show him (is supported by minio add metadata if not use same mongodb as in retriever)

Notes:
- use docker compose for container orchestration

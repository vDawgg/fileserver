import datetime
import hashlib
import logging
import os
from base64 import b64encode
from os.path import exists
import jwt
import grpc
from Crypto.PublicKey import RSA
from Crypto.Cipher import PKCS1_OAEP
from proto import authenticator_pb2_grpc
from proto import authenticator_pb2
from concurrent import futures
from pymongo import MongoClient
from pprint import pprint

# TODO: Change to ssl


def init_db(db):
    salt = os.urandom(32)
    random = os.urandom(64)
    uid = b64encode(random).decode("utf-8")  # This is probably pretty inefficient
    pwd = hashlib.pbkdf2_hmac(
        'sha256',
        b'fileserver',
        salt,
        150000
    )
    admin = {
        'name': 'admin',
        'password': pwd,
        'salt': salt,
        'id': uid  # TODO: Make sure that the same id cannot be generated!
    }
    pprint(db.password.insert_one(admin))


# TODO: sign this with pub key from authorization service once that is up and running
def token(name, uid):
    user = {
        'name': name,
        'id': uid
    }
    t = jwt.encode({
        'user': user,
        'exp': datetime.datetime.utcnow() + datetime.timedelta(hours=24)
    }, "secret",
        algorithm="HS256")
    print("Token:", t)
    return t  # Need to add encryption here


class authenticatorServicer(authenticator_pb2_grpc.authenticatorServicer):

    def getKeys(self, request, context):
        if not exists("keys/private.pem"):
            print("Creating keys")
            key = RSA.generate(2048)
            private_key = key.export_key('PEM', '3b1j873bhj')  # What should the passphrase be?
            file_out = open("keys/private.pem", "wb")
            file_out.write(private_key)
            file_out.close()
            public_key = key.publickey().export_key('PEM')
            file_out = open("keys/public.pem", "wb")
            file_out.write(public_key)
            file_out.close()
        public = open("keys/public.pem").read()
        return authenticator_pb2.Keys(keys=public)  # keys needs to be specified to not throw an error

    def login(self, request, context):

        # TODO: Change to env variables
        client = MongoClient(port=27017)
        try:
            print(client.server_info())
        except Exception:
            print("Connection to mongodb failed!")
        db = client.passwords
        if not client.list_database_names().__contains__('passwords'):
            print("yup")
            init_db(db)

        # TODO: Change this to also work with jwt
        private = RSA.importKey(open("keys/private.pem").read(), "3b1j873bhj")
        cipher = PKCS1_OAEP.new(private)
        name_crypt = request.name
        password_crypt = request.password
        name = cipher.decrypt(name_crypt).decode("utf-8")
        password = cipher.decrypt(password_crypt)

        orig = db.password.find_one({'name': name})
        if orig is None:
            return authenticator_pb2.Token(
                token=None,
                status=authenticator_pb2.AuthStatus.FAILED
            )

        check = hashlib.pbkdf2_hmac(
            'sha256',
            password,
            orig['salt'],
            150000
        )
        print("Check:", orig['password'] == check, orig['password'], check)

        if orig['password'] == check:
            print("Password was correct!")
            t = token(name, orig['id'])
            pprint(t)
            print(token(name, orig['id']))
            return authenticator_pb2.Token(
                token=token(name, orig['id']),
                status=authenticator_pb2.AuthStatus.OK
            )
        else:
            print('Password incorrect')
            return authenticator_pb2.Token(
                token=None,
                status=authenticator_pb2.AuthStatus.FAILED
            )


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(10))  # is this necessary?
    authenticator_pb2_grpc.add_authenticatorServicer_to_server(authenticatorServicer(), server)
    server.add_insecure_port('[::]:50051')  # TODO: Change to secure
    server.start()
    print("Started Server")
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()

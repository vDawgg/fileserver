import datetime
import logging
from pprint import pprint

import grpc
import jwt
from Crypto.PublicKey import RSA
from Crypto.Cipher import PKCS1_OAEP
from proto import unified_pb2_grpc
from proto import unified_pb2


def get_keys(stub):
    return stub.getKeys(unified_pb2.Empty())


def log_in(name, password, stub):
    key = RSA.importKey(open("keys/public.pem").read())
    cipher = PKCS1_OAEP.new(key)
    name = cipher.encrypt(name)
    password = cipher.encrypt(password)
    # f1 = open("nameEncrypted", "wb")
    # f2 = open("passEncrypted", "wb")
    # f1.write(name)
    # f2.write(password)
    request = unified_pb2.User(name=name, password=password)
    answer = stub.login(request)
    print("Answer:", answer.token)
    pprint(jwt.decode(answer.token, "secret", algorithms=["HS256"]))


def run():
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = unified_pb2_grpc.authenticatorStub(channel)
        print("----------Getting keys--------------")
        pubkey = get_keys(stub)
        print(pubkey)
        log_in(b'admin', b'fileserver', stub)


if __name__ == '__main__':
    logging.basicConfig()
    run()

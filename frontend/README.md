# Intro

This is the frontend for the fileserver. It is written in go and connected to all other services in this repository.

This frontend consists of a simple http-server, go-templates and grpc-helper functions to fetch data from the other services.

# Testing and manual execution

To generate the proto functions for the frontend run the helper-script inside this directory:
'''
./genproto.sh
'''

To start the frontend on its own run 
'''
go run frontend.go
'''
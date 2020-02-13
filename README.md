# Multicast Example in Golang

This is an example of doing multicast with golang. This repo is created as I need a multicast implementation to do some POC of multicast usage in K8S deployment. The binary can be run directly or within docker. 

## Install

After clone the repo to your local box, change to the root directory of it. You can build the docker image like this:

`make all`

With a golang installed on you box, you can build binary like this:

`make build-local`

For other method to build and install the binary, please check the Makefile.

## Usage

The binary can be stared as a server or a receiver. By default, it's started as receiver. And you can make it's stared as server with flag `-s`. If it's not specified, both receriver and server is started and join multicast group 239.0.0.1:12345. You can choose to specify your own group address by flag `-g`. For other usage, you can start and binary with `-h` flag for detail.

Start receiver as docker container:

`docker run -d multicast`

Start server as docker container: 

`docker run -d multicast -s`

## License

[MIT](LICENSE)

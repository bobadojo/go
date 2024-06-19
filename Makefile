APIS=$(shell find ../apis/bobadojo -name "*.proto")

all:	rpc grpc

rpc:
	protoc ${APIS} \
		--proto_path='../apis' \
		--go_opt='module=github.com/bobadojo/go' \
		--go_out='pkg'

grpc:
	protoc ${APIS} \
		--proto_path='../apis' \
		--go-grpc_opt='module=github.com/bobadojo/go' \
		--go-grpc_out='pkg'

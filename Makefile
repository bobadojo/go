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

gapic:
	protoc ${APIS} \
		--proto_path='../apis' \
        	--go_gapic_opt='go-gapic-package=github.com/bobadojo/go/pkg/gapic;gapic' \
        	--go_gapic_opt='grpc-service-config=grpc_service_config.json' \
        	--go_gapic_opt='module=github.com/bobadojo/go' \
        	--go_gapic_out='.'

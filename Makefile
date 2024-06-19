APIS=$(shell find ../apis/bobadojo -name "*.proto")

build:
	go install ./...

clean:
	rm -rf cmd pkg

all:	rpc grpc gapic cli

rpc:
	mkdir -p pkg
	protoc ${APIS} \
		--proto_path='../apis' \
		--go_opt='module=github.com/bobadojo/go/pkg' \
		--go_out='pkg'

grpc:
	mkdir -p pkg
	protoc ${APIS} \
		--proto_path='../apis' \
		--go-grpc_opt='module=github.com/bobadojo/go/pkg' \
		--go-grpc_out='pkg'

gapic:
	go install github.com/googleapis/gapic-generator-go/cmd/protoc-gen-go_gapic@latest
	mkdir -p pkg
	protoc ${APIS} \
		--proto_path='../apis' \
        	--go_gapic_opt='go-gapic-package=github.com/bobadojo/go/pkg/gapic;gapic' \
        	--go_gapic_opt='grpc-service-config=grpc_service_config.json' \
        	--go_gapic_opt='module=github.com/bobadojo/go/pkg' \
        	--go_gapic_out='pkg'

cli:
	mkdir -p cmd/bobatool
	protoc ${APIS} \
		--proto_path='../apis' \
        	--go_cli_opt='root=bobatool' \
        	--go_cli_opt='gapic=github.com/bobadojo/go/pkg/gapic' \
        	--go_cli_out=cmd/bobatool

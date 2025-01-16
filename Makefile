APIS=$(shell find apis/bobadojo -name "*.proto")

build:
	go install ./...

clean:
	rm -rf cmd pkg

all:	rpc grpc gapic cli build

descriptor:
	git submodule init
	git submodule update
	mkdir -p pkg
	protoc ${APIS} \
		--proto_path='apis' \
		--include_imports \
		--descriptor_set_out=descriptor.pb

rpc:
	git submodule init
	git submodule update
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	mkdir -p pkg
	protoc ${APIS} \
		--proto_path='apis' \
		--go_opt='module=github.com/bobadojo/go/pkg' \
		--go_out='pkg'

grpc:
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	mkdir -p pkg
	protoc ${APIS} \
		--proto_path='apis' \
		--go-grpc_opt='module=github.com/bobadojo/go/pkg' \
		--go-grpc_out='pkg'

gapic:
	go install github.com/googleapis/gapic-generator-go/cmd/protoc-gen-go_gapic@v0.42.0
	mkdir -p pkg
	protoc ${APIS} \
		--proto_path='apis' \
        	--go_gapic_opt='go-gapic-package=github.com/bobadojo/go/pkg/gapic;gapic' \
        	--go_gapic_opt='grpc-service-config=grpc_service_config.json' \
        	--go_gapic_opt='module=github.com/bobadojo/go/pkg' \
        	--go_gapic_out='pkg'

cli:
	go install github.com/googleapis/gapic-generator-go/cmd/protoc-gen-go_cli@v0.42.0
	mkdir -p cmd/bobatool
	protoc ${APIS} \
		--proto_path='apis' \
        	--go_cli_opt='root=bobatool' \
        	--go_cli_opt='gapic=github.com/bobadojo/go/pkg/gapic' \
        	--go_cli_out=cmd/bobatool

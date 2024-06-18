rpc:
	protoc \
		../apis/bobadojo/stores/v1/stores.proto \
		--proto_path='../apis' \
		--go_opt='module=github.com/bobadojo/go' \
		--go_out='.'

grpc:
	protoc \
		../apis/bobadojo/stores/v1/stores.proto \
		 --proto_path='../apis' \
		 --go-grpc_opt='module=github.com/bobadojo/go' \
		 --go-grpc_out='.'

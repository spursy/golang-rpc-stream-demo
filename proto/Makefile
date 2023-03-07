.PHONY: stream
stream:
	protoc \
     -I . \
     --go_out=.  --go_opt=paths=source_relative\
     --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative \
     ./stream_demo/stream_demo.proto
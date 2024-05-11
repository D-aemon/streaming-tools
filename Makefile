build: build_proto build_image

build_proto:
	protoc -I ./backend/proto \
           -I $GOPATH/src/google/api \
           --go_out ./backend/proto --go_opt paths=source_relative \
           --go-grpc_out ./backend/proto --go-grpc_opt paths=source_relative \
           --openapiv2_out ./backend/proto \
           --grpc-gateway_out ./backend/proto --grpc-gateway_opt paths=source_relative ./backend/proto/service.proto
build_image:


push_image:


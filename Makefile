test_env:
  export DB_USER="root"
  export DB_PASSWD="123456"
  export DB_IP="xxxx"
  export DB_PORT="3306"
  export DB_DBNAME="st"
  export LOG_LEVEL="debug"
  export HTTP_PORT="9000"
  export GRPC_PORT="9001"

build: build_proto build_backend build_frontend

build_backend:

build_frontend:

build_proto:
	protoc -I ./backend/proto \
           --go_out ./backend/proto --go_opt paths=source_relative \
           --go-grpc_out ./backend/proto --go-grpc_opt paths=source_relative \
           --openapiv2_out ./backend/proto \
           --grpc-gateway_out ./backend/proto --grpc-gateway_opt paths=source_relative ./backend/proto/*.proto

release: push_image push_charts

push_image:

push_charts:

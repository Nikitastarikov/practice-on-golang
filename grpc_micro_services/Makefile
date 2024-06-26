############################################################
# Переменные
############################################################

docker_image_name=grpc_server

############################################################
# Команды
############################################################

# run
run:
	go run cmd/main.go

# unit_test
unit_test:
	go test -v -cover ./...

# grpc
grpc:
	protoc --proto_path=proto reverse.proto \
		--go_out=paths=source_relative:proto \
		--go-grpc_out=paths=source_relative:proto

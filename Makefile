proto-logger-gen:
	protoc --proto_path=. --go_out=plugins=grpc:internal logger.proto
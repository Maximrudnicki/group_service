
build:
	protoc -Iproto --go_opt=module=group_service --go_out=. --go-grpc_opt=module=group_service --go-grpc_out=. proto/*.proto
	go build -o bin/group_service.exe ./cmd/.
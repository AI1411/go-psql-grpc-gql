migrate:
	migrate -path db/migrations -database "postgresql://root:root@127.0.0.1:5432/go_pg?sslmode=disable" -verbose up

gen-user-proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc/test.proto
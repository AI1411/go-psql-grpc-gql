migrate:
	migrate -path db/migrations -database "postgresql://root:root@127.0.0.1:15432/go_pg?sslmode=disable" -verbose up
migrate-for-test:
	migrate -path db/migrations -database "postgresql://postgres:postgres@127.0.0.1:25432/go_pg_test?sslmode=disable" -verbose up

gen-test-proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --doc_out=./grpc/test --doc_opt=html,index.html grpc/test/test.proto

gen-user-proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --doc_out=./grpc/user --doc_opt=html,index.html grpc/user/user.proto

test:
	richgo test -v ./...
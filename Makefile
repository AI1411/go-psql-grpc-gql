migrate:
	migrate -path db/migrations -database "postgresql://root:root@127.0.0.1:15432/go_pg?sslmode=disable" -verbose up
migrate-for-test:
	migrate -path db/migrations -database "postgresql://postgres:postgres@127.0.0.1:25432/go_pg_test?sslmode=disable" -verbose up

gen-test-proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --doc_out=./grpc/doc --doc_opt=html,test.html grpc/test.proto

gen-user-proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --doc_out=./grpc/doc --doc_opt=html,user.html grpc/user.proto

gen-task-proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --doc_out=./grpc/doc --doc_opt=html,task.html grpc/task.proto

gen-product-proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --doc_out=./grpc/doc --doc_opt=html,product.html grpc/product.proto

test:
	richgo test -v ./...

gqlgen:
	go get github.com/99designs/gqlgen@v0.17.17 && go run github.com/99designs/gqlgen && go mod tidy
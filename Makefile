migrate:
	migrate -path db/migrations -database "postgresql://root:root@127.0.0.1:15432/go_pg?sslmode=disable" -verbose up
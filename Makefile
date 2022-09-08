migrate:
	migrate -database 'postgres://root:root@localhost:15432/go_pg?sslmode=disable' -path db/migrations up
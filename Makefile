postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:12-alpine
create-db:
	docker exec -it postgres12 createdb --username=root --owner=root simple-bank
drop-db:
	docker exec -it postgres12  simple-bank
migrate-up:
	 migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple-bank?sslmode=disable" -verbose up
migrate-down:
	 migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple-bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
.PHONY: postgres create-db drop-db migrate-up migrate-down sqlc

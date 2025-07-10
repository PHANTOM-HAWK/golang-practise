postgres:
	docker run --name postgres12 -e POSTGRES_PASSWORD=root -e POSTGRES_USER=root -p 5432:5432 -d postgres:12-alpine

stoppostgres:
	docker stop postgres12

removepostgres:
	docker rm postgres12

restartpostgres:
	docker restart postgres12

createdb:
	docker exec -it postgres12 psql -U root -c "CREATE DATABASE simple_bank;"


deletedb:
	docker exec -it postgres12 psql -U root -c  "DROP DATABASE IF EXISTS simple_bank;"

migrateup:
	migrate -path db/migration \
  -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" \
  up

migratedown:
	migrate -path db/migration \
  -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" \
 down

 test:
	go test -v -cover ./...	

.PHONY: postgres createdb deletedb migrateup migratedown stoppostgres removepostgres restartpostgres
DB_URL=postgresql://root:secret@localhost:5432/simple-bank?sslmode=disable

# DB
postgres:
	sudo docker run --name postgres17 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17-alpine
startdb:
	sudo docker start postgres17
createdb:
	sudo docker exec -it postgres17 createdb --username=root --owner=root simple-bank
dropdb:
	sudo docker exec -it postgres17 dropdb simple-bank
logindb:
	sudo docker exec -it postgres17 psql -U root -d simple-bank

# Migrations
createmigration:
	migrate create -ext sql -dir db/migrations -seq $(NAME)
migrateup:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up
migrateup1:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up 1
migratedown:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down
migratedown1:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down 1

# DB Docs
db_docs:
	dbdocs build doc/db.dbml
db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

# Proto
proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

evans:
	evans --host localhost --port 9090 -r repl

sqlc:
	sqlc generate
mockdb:
	mockgen -package mockdb -destination db/mock/store.go github.com/WillChen936/simple-bank/db/sqlc Store
test:
	go test -v -cover ./...
server:
	go run main.go

.PHONY: proto
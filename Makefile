# DB
postgres:
	sudo docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17-alpine
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
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple-bank?sslmode=disable" -verbose up
migrateup1:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple-bank?sslmode=disable" -verbose up 1
migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple-bank?sslmode=disable" -verbose down
migratedown1:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple-bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate
mockdb:
	mockgen -package mockdb -destination db/mock/store.go github.com/WillChen936/simple-bank/db/sqlc Store
test:
	go test -v -cover ./...
server:
	go run main.go
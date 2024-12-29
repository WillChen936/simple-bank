# DB
postgres:
	sudo docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17-alpine
createdb:
	sudo docker exec -it postgres17 createdb --username=root --owner=root simple-bank
dropdb:
	sudo docker exec -it postgres17 dropdb simple-bank
logindb:
	sudo docker exec -it postgres17 psql -U root -d simple-bank


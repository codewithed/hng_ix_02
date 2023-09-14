postgres:
	docker run --name postgres15 -p 5432:5432  -e POSTGRES_PASSWORD=dbsecret -e POSTGRES_USER=root -d postgres

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root hng_ix_02

dropdb:
	docker exec -it postgres15 dropdb --username=root hng_ix_02

migrateup:
	migrate -path db/migrations -database "postgresql://root:dbsecret@localhost:5432/hng_ix_02?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migrations -database "postgresql://root:dbsecret@localhost:5432/hng_ix_02?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migrations -database "postgresql://root:dbsecret@localhost:5432/hng_ix_02?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migrations -database "postgresql://root:dbsecret@localhost:5432/hng_ix_02?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate
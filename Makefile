postgres:
	docker run --name postgres15 -p 5432:5432  -e POSTGRES_PASSWORD=dbsecret -e POSTGRES_USER=root -d postgres

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root hng_ix_02

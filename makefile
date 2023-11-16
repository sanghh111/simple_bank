up:
	migrate --path db/migration --database "postgresql://root:passroot@localhost:5432/simple_bank?sslmode=disable" --verbose up
down:
	migrate --path db/migration --database "postgresql://root:passroot@localhost:5432/simple_bank?sslmode=disable" --verbose down
create_db:
	docker exec -it postgres-16 createdb --username=root --owner=root simple_bank
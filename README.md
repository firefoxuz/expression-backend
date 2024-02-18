docker compose up

docker run -v ./db/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://expression_user:expression_password@localhost:5432/expression_db?sslmode=disable
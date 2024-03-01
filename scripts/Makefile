create_migrations:
	migrate create -ext sql -dir ../migrations $(NAME)

apply_migrations:
	migrate -path ../migrations -database "postgres://postgres:check@localhost:5432/webhooker?sslmode=disable" up

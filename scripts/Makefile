create_migrations:
	migrate create -ext sql -dir ../migrations $(NAME)

apply_migrations:
	migrate -path ../migrations -database "postgres://postgres:postgres@localhost:5432/leetcode?sslmode=disable" up

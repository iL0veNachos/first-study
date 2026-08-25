include .env
export

service-run:
	go run main.go


migrate-up:
	migrate -path migrations -database ${SERVER_CONN} up

migrate-down:
	migrate -path migrations -database ${SERVER_CONN} down
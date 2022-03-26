.PHONY: migrate
migrate: 
	goose -dir deployment/migrations postgres "user=postgres password=postgres dbname=storage-service sslmode=disable" up

generate-proto:

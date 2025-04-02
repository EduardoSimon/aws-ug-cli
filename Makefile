.PHONY: up down status clean seed

# Start the DynamoDB local container
up:
	docker-compose up -d

# Stop the DynamoDB local container
down:
	docker-compose down

# Show the status of the DynamoDB local container
status:
	docker-compose ps

# Stop the container and remove the volume
clean:
	docker-compose down -v

# Show logs from the DynamoDB local container
logs:
	docker-compose logs -f

# Generate and seed catalog data into DynamoDB
seed:
	go run . seed
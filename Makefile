container_name = go-books
database_name = books-db
database_user = root
database_password = 123

run:
	@echo "Running the Go application..."
	@go run cmd/main.go
	@echo "Application is running."

db-init:
	@echo "Initializing the database container $(container_name)..."
	@docker run --name $(container_name) -p 5432:5432 -e POSTGRES_USER=$(database_user) -e POSTGRES_PASSWORD=$(database_password) -d postgres:16.1
	@echo "Database container $(container_name) initialized successfully."

db-enter:
	@if [ "`docker ps -q -f name=$(container_name)`" ]; then \
		echo "Entering the database container $(container_name)..."; \
		docker exec -it $(container_name) psql; \
	else \
		echo "Database container $(container_name) is not running. Please start the container first."; \
	fi

db-create:
	@if [ "`docker ps -q -f name=$(container_name)`" ]; then \
		echo "Creating database $(database_name)..."; \
		docker exec -it $(container_name) createdb --username=root --owner=$(database_user) $(database_name); \
		echo "Database $(database_name) created successfully."; \
	else \
		echo "Database container $(container_name) is not running. Please start the container first."; \
	fi

db-drop:
	@if [ "`docker ps -q -f name=$(container_name)`" ]; then \
		echo "Dropping database go-chat..."; \
		docker exec -it $(container_name) dropdb go-chat; \
		echo "Database go-chat dropped successfully."; \
	else \
		echo "Database container $(container_name) is not running. Please start the container first."; \
	fi

db-delete:
	@echo "Stopping and removing the database container $(container_name)..."
	@docker stop $(container_name) || true
	@docker rm $(container_name) || true
	@echo "Database container $(container_name) deleted successfully."

db-start:
	@if [ "`docker ps -aq -f name=$(container_name)`" ]; then \
		echo "Starting the database container $(container_name)..."; \
		docker start $(container_name); \
		echo "Database container $(container_name) started successfully."; \
	else \
		echo "Database container $(container_name) does not exist. Please create the container first."; \
	fi

db-stop:
	@if [ "`docker ps -q -f name=$(container_name)`" ]; then \
		echo "Stopping the database container $(container_name)..."; \
		docker stop $(container_name); \
		echo "Database container $(container_name) stopped successfully."; \
	else \
		echo "Database container $(container_name) is already stopped or does not exist."; \
	fi
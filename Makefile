include .env

empty:
	echo "empty"

# enter the container
backend-ssh:
	docker exec -it $(BACKEND_CONTAINER_NAME) sh

db-sh:
	docker exec -it $(POSTGRES_HOST) sh
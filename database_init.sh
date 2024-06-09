#/bin/bash

CONTAINER_NAME="mongo"
CONTAINER_ID=$(docker ps -q --filter "name=$CONTAINER_NAME")

if [ -z "$CONTAINER_ID" ]; then
  echo "Container $CONTAINER_NAME not found."
  exit 1
fi

# Execute the mongo shell command within the container to create a database
docker exec -it $CONTAINER_ID mongosh --eval "use docker_sandbox"

if [ $? -eq 0 ]; then
  echo "Database docker_sandbox created successfully."
else
  echo "Failed to create database docker_sandbox."
fi

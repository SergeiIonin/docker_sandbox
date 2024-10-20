#!/bin/bash

CONTAINER_NAME="mongo"
DB_NAME="docker_sandbox"
COLLECTION="composes"

CONTAINER_ID=$(docker ps -q --filter "name=$CONTAINER_NAME")

if [ -z "$CONTAINER_ID" ]; then
  echo "Container $CONTAINER_NAME not found."
  exit 1
fi

COLLECTION_EXISTS=$(docker exec $CONTAINER_ID mongosh --quiet --eval "db = db.getSiblingDB('$DB_NAME'); db.getCollectionNames().includes('$COLLECTION')")
if [ "$COLLECTION_EXISTS" = "true" ]; then
  echo "Collection '$COLLECTION' already exists in database '$DB_NAME'. Skipping creation."
else
  docker exec -it $CONTAINER_ID mongosh --eval "db = db.getSiblingDB('$DB_NAME'); db.createCollection('$COLLECTION');"
  
  if [ $? -eq 0 ]; then
    echo "Collection '$COLLECTION' created successfully in database '$DB_NAME'."
  else
    echo "Failed to create collection '$COLLECTION' in database '$DB_NAME'."
  fi
fi

#!/bin/bash

echo "Building and starting the application using Docker Compose..."
docker-compose up --build

echo "Application started. Frontend: http://localhost:3000, Backend: http://localhost:8080"

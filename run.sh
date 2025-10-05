#!/bin/bash

echo "Building and starting the application..."
gnome-terminal -- bash -c "cd backend/deployments && ./run_docker.sh; exec bash"

echo "Please wait a minute until the backend is ready..."
while ! nc -z localhost 8080; do
  sleep 1
done

gnome-terminal -- bash -c "cd frontend && npm start; exec bash"

echo "Application started. Frontend: http://localhost:3000, Backend: http://localhost:8080"
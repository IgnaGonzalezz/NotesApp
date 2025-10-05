#!/bin/bash

echo "Building and starting the application..."
gnome-terminal -- bash -c "cd backend/deployments && ./run_docker.sh; exec bash"

echo "Please wait 20 secs till the backend is ready..."
sleep 20

gnome-terminal -- bash -c "cd frontend && npm start; exec bash"

echo "Application started. Frontend: http://localhost:3000, Backend: http://localhost:8080"
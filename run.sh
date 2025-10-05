#!/bin/bash

echo "Building and starting the application using Docker Compose..."
gnome-terminal -- bash -c "cd backend/deployments && ./run_docker.sh; exec bash"

echo "Esperando 30 segundos a que el backend esté listo..."
sleep 30

gnome-terminal -- bash -c "cd frontend && npm start; exec bash"

echo "Application started. Frontend: http://localhost:3000, Backend: http://localhost:8080"
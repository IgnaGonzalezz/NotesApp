#!/bin/bash

echo "Building and starting the application..."
(cd backend/deployments && ./run_docker.sh) &

echo "Please wait until the backend is ready..."
sleep 60

(cd frontend && npm start) &

echo "Application started. Frontend: http://localhost:3000, Backend: http://localhost:8080"
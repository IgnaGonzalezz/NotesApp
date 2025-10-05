#!/bin/bash

# Parar y eliminar contenedores antiguos si existen
docker stop notes_app notes_db &> /dev/null
docker rm notes_app notes_db &> /dev/null

# Crear una red para que los contenedores se comuniquen
docker network create notes-network &> /dev/null

# Variables de entorno
DB_USER="notesuser"
DB_PASSWORD="notespass"
DB_NAME="notesdb"
DB_HOST="notes_db" # Importante: el nombre del host es el nombre del contenedor de la DB
DB_PORT="5432"

# 1. Iniciar la base de datos
echo "Iniciando base de datos..."
docker run -d \
    --name ${DB_HOST} \
    --network notes-network \
    -e POSTGRES_USER=${DB_USER} \
    -e POSTGRES_PASSWORD=${DB_PASSWORD} \
    -e POSTGRES_DB=${DB_NAME} \
    -v notes_data:/var/lib/postgresql/data \
    postgres:15

# 2. Construir la imagen de la aplicación
echo "Construyendo la imagen de la aplicación..."
docker build -t notes-app-image ../

# Esperar un poco para que la DB esté lista
echo "Esperando a que la base de datos se inicie..."
sleep 10

# 3. Iniciar la aplicación
echo "Iniciando la aplicación..."
docker run -d \
    --name notes_app \
    --network notes-network \
    -p 8080:8080 \
    -e DB_HOST=${DB_HOST} \
    -e DB_USER=${DB_USER} \
    -e DB_PASSWORD=${DB_PASSWORD} \
    -e DB_NAME=${DB_NAME} \
    -e DB_PORT=${DB_PORT} \
    notes-app-image

echo "\nEntorno listo. La aplicación debería estar disponible en http://localhost:8080"

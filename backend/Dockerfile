# Usa una imagen base de Go
FROM golang:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos de módulos de Go y descarga las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia el resto del código fuente de la aplicación
COPY . .

# Compila la aplicación
# El binario se creará en /app/notesapp
RUN go build -o notesapp ./cmd/app

# Expone el puerto en el que correrá la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación cuando se inicie el contenedor
CMD [ "./notesapp" ]

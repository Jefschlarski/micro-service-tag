# Imagem base para compilar o código Go
FROM golang:latest AS build

# Diretório de trabalho
WORKDIR /app

# Copia arquivos do projeto para o contêiner
COPY . .

# Compila o código Go
RUN go build -o micro-service-tag

# Imagem base para executar o binário Go
FROM debian:latest

# Configura diretório de trabalho
WORKDIR /app

# Copia o arquivo go.env
COPY --from=build /app/go.env .

# Copia o binário compilado do estágio anterior
COPY --from=build /app/micro-service-tag .

# Comando padrão para executar o binário Go
CMD ["./micro-service-tag"]
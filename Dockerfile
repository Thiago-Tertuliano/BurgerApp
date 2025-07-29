# Multi-stage build para o BurgerApp
FROM node:18-alpine AS frontend-build

# Configurar diretório de trabalho
WORKDIR /app/frontend

# Copiar arquivos de dependências do frontend
COPY projeto-hamburgueria/package*.json ./

# Instalar dependências
RUN npm ci --only=production

# Copiar código fonte do frontend
COPY projeto-hamburgueria/ .

# Build do frontend
RUN npm run build

# Stage do backend
FROM golang:1.23-alpine AS backend-build

# Instalar dependências do sistema
RUN apk add --no-cache git

# Configurar diretório de trabalho
WORKDIR /app/backend

# Copiar arquivos de dependências do backend
COPY backend-hamburgueria/go.mod backend-hamburgueria/go.sum ./

# Baixar dependências
RUN go mod download

# Copiar código fonte do backend
COPY backend-hamburgueria/ .

# Build do backend
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage final
FROM alpine:latest

# Instalar dependências necessárias
RUN apk --no-cache add ca-certificates tzdata

# Criar usuário não-root
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Configurar diretório de trabalho
WORKDIR /app

# Copiar binário do backend
COPY --from=backend-build /app/backend/main ./backend/main

# Copiar arquivos estáticos do frontend
COPY --from=frontend-build /app/frontend/dist ./frontend/dist

# Copiar arquivos de configuração
COPY backend-hamburgueria/config.env.example ./backend/.env

# Definir permissões
RUN chown -R appuser:appgroup /app

# Mudar para usuário não-root
USER appuser

# Expor portas
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Comando para executar a aplicação
CMD ["./backend/main"]
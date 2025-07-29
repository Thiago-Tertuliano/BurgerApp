# 🚀 Guia de Deploy - BurgerApp

## 📋 Visão Geral

Este documento descreve como fazer deploy do BurgerApp em diferentes ambientes, desde desenvolvimento local até produção.

## 🏗️ Arquitetura de Deploy

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Frontend      │    │    Backend      │    │   PostgreSQL    │
│   (Vue.js)      │◄──►│     (Go)        │◄──►│   (Database)    │
│   Porta 5173    │    │   Porta 8080    │    │   Porta 5432    │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

## 🐳 Deploy com Docker

### Pré-requisitos
- Docker 20.10+
- Docker Compose 2.0+
- 4GB RAM disponível

### Deploy Local com Docker Compose

```bash
# 1. Clone o repositório
git clone <url-do-repositorio>
cd BurgerApp

# 2. Configure as variáveis de ambiente
cp backend-hamburgueria/config.env.example backend-hamburgueria/.env
# Edite o arquivo .env conforme necessário

# 3. Execute o deploy
docker-compose up -d

# 4. Verifique os serviços
docker-compose ps

# 5. Acesse a aplicação
# Frontend: http://localhost:5173
# Backend API: http://localhost:8080
# Health Check: http://localhost:8080/health
```

### Comandos Úteis do Docker

```bash
# Ver logs dos serviços
docker-compose logs -f

# Parar todos os serviços
docker-compose down

# Rebuild das imagens
docker-compose build --no-cache

# Executar apenas o banco de dados
docker-compose up postgres

# Backup do banco de dados
docker-compose exec postgres pg_dump -U postgres hamburgueria > backup.sql

# Restaurar backup
docker-compose exec -T postgres psql -U postgres hamburgueria < backup.sql
```

## ☁️ Deploy em Produção

### Opção 1: Deploy Manual

#### 1. Preparar o Servidor
```bash
# Instalar Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# Instalar Docker Compose
sudo curl -L "https://github.com/docker/compose/releases/download/v2.20.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
```

#### 2. Configurar o Projeto
```bash
# Clonar o repositório
git clone <url-do-repositorio>
cd BurgerApp

# Configurar variáveis de produção
cp backend-hamburgueria/config.env.example backend-hamburgueria/.env.prod
# Editar .env.prod com configurações de produção
```

#### 3. Deploy
```bash
# Usar configuração de produção
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d

# Verificar status
docker-compose ps
```

### Opção 2: Deploy com CI/CD (GitHub Actions)

#### 1. Configurar Secrets no GitHub
- `DOCKER_REGISTRY`: URL do registro Docker
- `DOCKER_USERNAME`: Usuário do Docker
- `DOCKER_PASSWORD`: Senha do Docker
- `PRODUCTION_HOST`: IP/domínio do servidor de produção
- `PRODUCTION_USER`: Usuário SSH do servidor
- `PRODUCTION_KEY`: Chave SSH privada

#### 2. Deploy Automático
```bash
# O deploy acontece automaticamente quando:
# - Push para branch 'main' (produção)
# - Push para branch 'develop' (staging)

# Verificar status no GitHub Actions
# https://github.com/seu-usuario/burgerapp/actions
```

## 🔧 Configurações de Ambiente

### Desenvolvimento
```bash
# .env.development
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=hamburgueria_dev
PORT=8080
JWT_SECRET=dev_secret_key
```

### Staging
```bash
# .env.staging
DB_HOST=staging-db.example.com
DB_PORT=5432
DB_USER=burgerapp_staging
DB_PASSWORD=secure_password
DB_NAME=burgerapp_staging
PORT=8080
JWT_SECRET=staging_secret_key
```

### Produção
```bash
# .env.production
DB_HOST=production-db.example.com
DB_PORT=5432
DB_USER=burgerapp_prod
DB_PASSWORD=very_secure_password
DB_NAME=burgerapp_production
PORT=8080
JWT_SECRET=production_secret_key
```

## 🔒 Segurança

### Configurações de Segurança

#### 1. Firewall
```bash
# Configurar UFW
sudo ufw allow 22/tcp    # SSH
sudo ufw allow 80/tcp     # HTTP
sudo ufw allow 443/tcp    # HTTPS
sudo ufw allow 8080/tcp   # Backend API
sudo ufw enable
```

#### 2. SSL/TLS
```bash
# Instalar Certbot
sudo apt install certbot python3-certbot-nginx

# Obter certificado SSL
sudo certbot --nginx -d seu-dominio.com

# Renovação automática
sudo crontab -e
# Adicionar: 0 12 * * * /usr/bin/certbot renew --quiet
```

#### 3. Banco de Dados
```bash
# Configurar PostgreSQL para aceitar apenas conexões locais
# /etc/postgresql/15/main/postgresql.conf
listen_addresses = 'localhost'

# /etc/postgresql/15/main/pg_hba.conf
local   all             postgres                                peer
local   all             all                                     md5
host    all             all             127.0.0.1/32            md5
```

## 📊 Monitoramento

### Health Checks
```bash
# Verificar status da aplicação
curl http://localhost:8080/health

# Verificar logs
docker-compose logs -f backend

# Monitorar recursos
docker stats
```

### Logs
```bash
# Logs do backend
docker-compose logs backend

# Logs do frontend
docker-compose logs frontend

# Logs do banco de dados
docker-compose logs postgres

# Logs em tempo real
docker-compose logs -f
```

## 🔄 Backup e Recuperação

### Backup Automático
```bash
#!/bin/bash
# backup.sh
DATE=$(date +%Y%m%d_%H%M%S)
BACKUP_DIR="/backups"

# Backup do banco de dados
docker-compose exec -T postgres pg_dump -U postgres hamburgueria > $BACKUP_DIR/backup_$DATE.sql

# Backup dos arquivos de configuração
tar -czf $BACKUP_DIR/config_$DATE.tar.gz backend-hamburgueria/.env*

# Manter apenas os últimos 7 backups
find $BACKUP_DIR -name "backup_*.sql" -mtime +7 -delete
find $BACKUP_DIR -name "config_*.tar.gz" -mtime +7 -delete
```

### Restauração
```bash
# Restaurar banco de dados
docker-compose exec -T postgres psql -U postgres hamburgueria < backup_20241201_143022.sql

# Restaurar configurações
tar -xzf config_20241201_143022.tar.gz
```

## 🚨 Troubleshooting

### Problemas Comuns

#### 1. Container não inicia
```bash
# Verificar logs
docker-compose logs <service-name>

# Verificar recursos
docker system df
docker system prune -f
```

#### 2. Erro de conexão com banco
```bash
# Verificar se o PostgreSQL está rodando
docker-compose ps postgres

# Testar conexão
docker-compose exec postgres psql -U postgres -d hamburgueria
```

#### 3. Frontend não carrega
```bash
# Verificar se o backend está respondendo
curl http://localhost:8080/health

# Verificar logs do frontend
docker-compose logs frontend
```

### Comandos de Debug
```bash
# Entrar no container
docker-compose exec backend sh
docker-compose exec frontend sh

# Verificar variáveis de ambiente
docker-compose exec backend env

# Testar conectividade
docker-compose exec backend ping postgres
```

## 📈 Escalabilidade

### Horizontal Scaling
```bash
# Escalar backend
docker-compose up -d --scale backend=3

# Load balancer com Nginx
# Configurar nginx.conf para distribuir carga
```

### Vertical Scaling
```bash
# Aumentar recursos no docker-compose.yml
services:
  backend:
    deploy:
      resources:
        limits:
          memory: 1G
          cpus: '0.5'
```

## 📞 Suporte

### Contatos
- **Desenvolvedor**: Thiago Matos Tertuliano
- **Email**: [seu-email@exemplo.com]
- **GitHub**: [https://github.com/seu-usuario/burgerapp]

### Documentação Adicional
- [README.md](./README.md) - Documentação geral do projeto
- [API Documentation](./API.md) - Documentação da API
- [Database Schema](./DATABASE.md) - Esquema do banco de dados

---

**🍔 BurgerApp** - Deploy profissional e escalável para sua hamburgueria!
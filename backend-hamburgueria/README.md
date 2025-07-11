# 🍔 BurgerApp - Backend API

## 📋 Descrição

API REST desenvolvida em **Go** com **Gin** para o sistema de pedidos da hamburgueria. O backend fornece endpoints para gerenciar produtos, categorias, ingredientes e pedidos, com persistência em banco de dados PostgreSQL.

## 🚀 Tecnologias Utilizadas

### Core
- **Go 1.23** - Linguagem de programação
- **Gin** - Framework web para APIs REST
- **PostgreSQL** - Banco de dados relacional
- **lib/pq** - Driver PostgreSQL para Go

### Dependências Principais
- `gin-gonic/gin` v1.10.1 - Framework web
- `gin-contrib/cors` v1.7.6 - Middleware CORS
- `lib/pq` v1.10.9 - Driver PostgreSQL
- `joho/godotenv` v1.5.1 - Variáveis de ambiente
- `golang-jwt/jwt/v5` v5.0.0 - Autenticação JWT
- `golang.org/x/crypto` v0.39.0 - Criptografia

## 📁 Estrutura do Projeto

```
backend-hamburgueria/
├── config/              # Configurações da aplicação
│   └── config.go        # Middleware CORS
├── database/            # Conexão e configuração do banco
│   └── database.go      # Conexão PostgreSQL e criação de tabelas
├── handlers/            # Manipuladores das requisições HTTP
│   └── handlers.go      # Handlers para produtos, pedidos, etc.
├── models/              # Modelos de dados
│   └── models.go        # Structs para JSON e banco
├── routes/              # Configuração das rotas
│   └── routes.go        # Definição dos endpoints
├── scripts/             # Scripts SQL e utilitários
├── main.go              # Ponto de entrada da aplicação
├── go.mod               # Dependências do Go
├── go.sum               # Checksums das dependências
├── config.env.example   # Exemplo de configuração
└── README.md            # Este arquivo
```

## 🎯 Funcionalidades

### 1. **Gestão de Produtos**
- Listar produtos disponíveis
- Buscar produtos por categoria
- Informações completas com imagens e preços

### 2. **Gestão de Categorias**
- Listar categorias de produtos
- Organização hierárquica do menu

### 3. **Gestão de Ingredientes**
- Listar ingredientes para montagem
- Categorização por tipo (pão, carne, queijo, etc.)
- Preços adicionais para personalização

### 4. **Sistema de Pedidos**
- Criar novos pedidos
- Listar pedidos com filtros
- Atualizar status dos pedidos
- Detalhes completos dos pedidos

### 5. **Transações Seguras**
- Uso de transações SQL para consistência
- Validação de dados de entrada
- Tratamento de erros robusto

## 🔧 Configuração e Instalação

### Pré-requisitos
- **Go 1.23+** instalado
- **PostgreSQL** rodando
- **Git** para clonar o repositório

### Instalação

1. **Clone o repositório**
```bash
git clone <url-do-repositorio>
cd backend-hamburgueria
```

2. **Instale as dependências**
```bash
go mod download
```

3. **Configure o banco de dados**
```bash
# Crie o banco de dados PostgreSQL
createdb hamburgueria

# Ou use psql
psql -U postgres -c "CREATE DATABASE hamburgueria;"
```

4. **Configure as variáveis de ambiente**
```bash
# Copie o arquivo de exemplo
cp config.env.example .env

# Edite as configurações conforme necessário
nano .env
```

5. **Execute o projeto**
```bash
go run main.go
```

6. **Teste a API**
```bash
curl http://localhost:8080/health
```

## 📖 Como Usar

### Desenvolvimento
```bash
go run main.go          # Executar em modo desenvolvimento
go build                # Compilar para produção
./backend-hamburgueria  # Executar binário compilado
```

### Scripts Disponíveis
- `go run main.go` - Executar aplicação
- `go build` - Compilar para produção
- `go test ./...` - Executar testes
- `go mod tidy` - Limpar dependências

## 🏗️ Arquitetura

### Padrão MVC Simplificado

#### 1. **Models** (`models/models.go`)
- Structs para representar dados
- Tags JSON para serialização
- Separação entre modelos de dados e requisições

#### 2. **Handlers** (`handlers/handlers.go`)
- Lógica de negócio
- Validação de dados
- Interação com banco de dados
- Respostas HTTP

#### 3. **Routes** (`routes/routes.go`)
- Definição dos endpoints
- Agrupamento de rotas
- Middleware de autenticação

#### 4. **Database** (`database/database.go`)
- Conexão com PostgreSQL
- Criação automática de tabelas
- Configuração de pool de conexões

### Fluxo de Dados

```
Cliente → Routes → Handlers → Database → PostgreSQL
   ↑                                    ↓
   ← JSON Response ← ← ← ← ← ← ← ← ← ← ← ←
```

## 🗄️ Banco de Dados

### Tabelas Principais

#### 1. **categories**
```sql
- id (SERIAL PRIMARY KEY)
- name (VARCHAR(100))
- description (TEXT)
- created_at (TIMESTAMP)
```

#### 2. **products**
```sql
- id (SERIAL PRIMARY KEY)
- name (VARCHAR(200))
- description (TEXT)
- price (DECIMAL(10,2))
- category_id (INTEGER FK)
- image_url (VARCHAR(500))
- is_available (BOOLEAN)
- created_at (TIMESTAMP)
```

#### 3. **ingredients**
```sql
- id (SERIAL PRIMARY KEY)
- name (VARCHAR(100))
- price (DECIMAL(10,2))
- category (VARCHAR(50))
- is_available (BOOLEAN)
- created_at (TIMESTAMP)
```

#### 4. **orders**
```sql
- id (SERIAL PRIMARY KEY)
- customer_name (VARCHAR(200))
- table_number (INTEGER)
- total_amount (DECIMAL(10,2))
- status (VARCHAR(50))
- notes (TEXT)
- created_at (TIMESTAMP)
- updated_at (TIMESTAMP)
```

#### 5. **order_items**
```sql
- id (SERIAL PRIMARY KEY)
- order_id (INTEGER FK)
- product_id (INTEGER FK)
- ingredients (TEXT)
- quantity (INTEGER)
- unit_price (DECIMAL(10,2))
- total_price (DECIMAL(10,2))
- notes (TEXT)
- created_at (TIMESTAMP)
```

## 🔌 Endpoints da API

### Produtos e Categorias
```http
GET /api/products      # Listar produtos
GET /api/categories    # Listar categorias
GET /api/ingredients   # Listar ingredientes
```

### Pedidos
```http
GET    /api/orders              # Listar pedidos
GET    /api/orders?status=preparing  # Filtrar por status
POST   /api/orders              # Criar pedido
GET    /api/orders/:id          # Detalhes do pedido
PUT    /api/orders/:id/status   # Atualizar status
```

### Health Check
```http
GET /health  # Verificar status do servidor
```

## 📊 Modelos de Dados

### Product
```json
{
  "id": 1,
  "name": "Classic Burger",
  "description": "Hambúrguer tradicional",
  "price": 24.90,
  "category_id": 1,
  "category": {
    "id": 1,
    "name": "Burgers",
    "description": "Hambúrgueres"
  },
  "image_url": "/burger.jpg",
  "is_available": true,
  "created_at": "2024-01-01T10:00:00Z"
}
```

### Order
```json
{
  "id": 1,
  "customer_name": "João Silva",
  "table_number": 5,
  "total_amount": 45.80,
  "status": "preparing",
  "notes": "Sem cebola",
  "created_at": "2024-01-01T10:00:00Z",
  "updated_at": "2024-01-01T10:05:00Z",
  "items": [...]
}
```

## 🔒 Segurança

### CORS
- Configurado para permitir frontend
- Origens permitidas: localhost:5173, localhost:5174
- Métodos: GET, POST, PUT, DELETE, OPTIONS

### Validação
- Validação de dados de entrada
- Sanitização de parâmetros
- Tratamento de erros robusto

### Transações
- Uso de transações SQL para consistência
- Rollback automático em caso de erro
- Isolamento de operações críticas

## 🛠️ Desenvolvimento

### Estrutura de Handlers
```go
func HandlerName(c *gin.Context, db *sql.DB) {
    // 1. Validar dados de entrada
    // 2. Executar lógica de negócio
    // 3. Interagir com banco de dados
    // 4. Retornar resposta JSON
}
```

### Padrões Utilizados
- **Dependency Injection** - Passagem de `db` para handlers
- **Error Handling** - Tratamento consistente de erros
- **JSON Tags** - Serialização automática
- **SQL Prepared Statements** - Prevenção de SQL injection

### Boas Práticas
- Código limpo e bem documentado
- Separação de responsabilidades
- Reutilização de código
- Testes unitários (quando aplicável)

## 🚀 Deploy

### Build de Produção
```bash
# Compilar para Linux
GOOS=linux GOARCH=amd64 go build -o backend-hamburgueria

# Compilar para Windows
GOOS=windows GOARCH=amd64 go build -o backend-hamburgueria.exe

# Compilar para macOS
GOOS=darwin GOARCH=amd64 go build -o backend-hamburgueria
```

### Docker (Opcional)
```dockerfile
FROM golang:1.23-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]
```

## 🔧 Configurações

### Variáveis de Ambiente
```bash
# Servidor
PORT=8080

# Banco de dados
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=hamburgueria

# Segurança
JWT_SECRET=seu_jwt_secret_aqui
```

### Configuração do Gin
```go
// Modo de produção
gin.SetMode(gin.ReleaseMode)

// Middleware CORS
r.Use(config.CORS())

// Middleware de logging
r.Use(gin.Logger())
r.Use(gin.Recovery())
```

## 📚 Recursos de Aprendizado

### Go
- [Documentação Oficial](https://golang.org/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go by Example](https://gobyexample.com/)

### Gin Framework
- [Documentação Oficial](https://gin-gonic.com/)
- [Exemplos de Uso](https://github.com/gin-gonic/examples)

### PostgreSQL
- [Documentação Oficial](https://www.postgresql.org/docs/)
- [lib/pq Driver](https://github.com/lib/pq)

## 🤝 Contribuição

### Como Contribuir
1. Fork o projeto
2. Crie uma branch para sua feature
3. Commit suas mudanças
4. Push para a branch
5. Abra um Pull Request

### Padrões de Código
- Use `gofmt` para formatação
- Siga as convenções do Go
- Escreva testes quando possível
- Documente funções públicas

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

## 👥 Autores

- **Desenvolvedor**: [Seu Nome]
- **Data**: 2024
- **Versão**: 1.0.0

---

## 🎯 Próximos Passos

### Melhorias Planejadas
- [ ] Autenticação JWT completa
- [ ] Middleware de logging
- [ ] Cache com Redis
- [ ] Rate limiting
- [ ] Documentação Swagger
- [ ] Testes automatizados
- [ ] CI/CD pipeline

### Recursos Técnicos
- [ ] Migrations com golang-migrate
- [ ] Validação com go-playground/validator
- [ ] Logging estruturado
- [ ] Métricas com Prometheus
- [ ] Health checks avançados

---

**🍔 BurgerApp Backend** - API robusta e escalável para o sistema de pedidos! 
# 🍔 BurgerApp - Sistema Completo de Pedidos

## 📋 Descrição

Sistema completo de pedidos para hamburgueria desenvolvido com **frontend em Vue.js 3** e **backend em Go**. O projeto permite visualizar o cardápio, montar lanches personalizados e gerenciar pedidos em tempo real através de uma interface moderna e responsiva, com persistência completa em banco de dados PostgreSQL.

## 🏗️ Arquitetura do Sistema

```
┌─────────────────┐    HTTP/JSON    ┌─────────────────┐
│   Frontend      │ ◄──────────────► │    Backend      │
│   Vue.js 3      │                 │   Go + Gin      │
│   (Porta 5173)  │                 │  (Porta 8080)   │
└─────────────────┘                 └─────────────────┘
                                              │
                                              ▼
                                    ┌─────────────────┐
                                    │   PostgreSQL    │
                                    │   (Porta 5432)  │
                                    └─────────────────┘
```

## 🚀 Tecnologias Utilizadas

### Frontend (Vue.js 3)
- **Vue.js 3** - Framework JavaScript progressivo
- **Composition API** - Sistema de reatividade
- **Vite** - Build tool e dev server moderno
- **CSS3** - Estilização com variáveis CSS e animações

### Backend (Go)
- **Go 1.23** - Linguagem de programação
- **Gin** - Framework web para APIs REST
- **PostgreSQL** - Banco de dados relacional
- **lib/pq** - Driver PostgreSQL para Go

### Dependências Principais
- **Frontend**: Vue 3.5.17, Vite 7.0.0
- **Backend**: Gin 1.10.1, CORS 1.7.6, JWT 5.0.0

## 📁 Estrutura do Projeto

```
Projeto/
├── projeto-hamburgueria/          # Frontend Vue.js
│   ├── src/
│   │   ├── components/            # Componentes Vue
│   │   │   ├── Header.vue         # Cabeçalho da aplicação
│   │   │   ├── Menu.vue           # Cardápio de produtos
│   │   │   ├── Kitchen.vue        # Interface da cozinha
│   │   │   └── CustomBurger.vue   # Montador de lanches
│   │   ├── assets/                # Recursos estáticos
│   │   ├── App.vue                # Componente raiz
│   │   ├── main.js                # Ponto de entrada
│   │   └── api.js                 # Configuração da API
│   ├── public/                    # Arquivos estáticos
│   ├── package.json               # Dependências frontend
│   └── README.md                  # Documentação frontend
│
├── backend-hamburgueria/          # Backend Go
│   ├── config/                    # Configurações
│   │   └── config.go              # Middleware CORS
│   ├── database/                  # Conexão com banco
│   │   └── database.go            # PostgreSQL e tabelas
│   ├── handlers/                  # Manipuladores HTTP
│   │   └── handlers.go            # Handlers da API
│   ├── models/                    # Modelos de dados
│   │   └── models.go              # Structs Go
│   ├── routes/                    # Configuração de rotas
│   │   └── routes.go              # Endpoints da API
│   ├── scripts/                   # Scripts SQL
│   ├── main.go                    # Ponto de entrada
│   ├── go.mod                     # Dependências Go
│   └── README.md                  # Documentação backend
│
└── README.md                      # Este arquivo
```

## 🎯 Funcionalidades

### 1. **Cardápio Interativo**
- Visualização de produtos por categoria
- Cards com imagens 

### 2. **Montador de Lanches Personalizados**
- Seleção de ingredientes por categoria
- Interface intuitiva

### 3. **Sistema de Pedidos**
- Envio automático para o backend
- Sincronização em tempo real

### 4. **Interface da Cozinha**
- Painel lateral responsivo
- Lista de pedidos em tempo real
- Controle de status (preparando → pronto → entregue)
- Barra de progresso visual

### 5. **Persistência Completa**
- Banco de dados PostgreSQL
- Dados não se perdem ao recarregar

## 🔧 Configuração e Instalação

### Pré-requisitos
- **Node.js 16+** (para frontend)
- **Go 1.23+** (para backend)
- **PostgreSQL 12+** (banco de dados)

### Instalação Completa

#### 1. **Clone o repositório**
```bash
git clone <url-do-repositorio>
cd Projeto
```

#### 2. **Configure o banco de dados**
```bash
# Crie o banco de dados PostgreSQL
createdb hamburgueria

# Ou use psql
psql -U postgres -c "CREATE DATABASE hamburgueria;"
```

#### 3. **Configure o backend**
```bash
cd backend-hamburgueria

# Instale as dependências Go
go mod download

# Configure as variáveis de ambiente
cp config.env.example .env
# Edite o arquivo .env conforme necessário

# Execute o backend
go run main.go
```

#### 4. **Configure o frontend**
```bash
cd ../projeto-hamburgueria

# Instale as dependências Node.js
npm install

# Execute o frontend
npm run dev
```

#### 5. **Acesse a aplicação**
- **Frontend**: http://localhost:5173
- **Backend API**: http://localhost:8080
- **Health Check**: http://localhost:8080/health

## 📖 Como Usar

### Desenvolvimento

#### Frontend
```bash
cd projeto-hamburgueria
npm run dev          # Servidor de desenvolvimento
npm run build        # Build de produção
npm run preview      # Preview do build
```

#### Backend
```bash
cd backend-hamburgueria
go run main.go       # Executar em desenvolvimento
go build             # Compilar para produção
./backend-hamburgueria  # Executar binário
```

### Scripts Disponíveis

#### Frontend
- `npm run dev` - Servidor de desenvolvimento 
- `npm run build` - Build otimizado para produção
- `npm run preview` - Preview do build de produção

#### Backend
- `go run main.go` - Executar aplicação
- `go build` - Compilar para produção
- `go test ./...` - Executar testes
- `go mod tidy` - Limpar dependências

## 🗄️ Banco de Dados

### Tabelas Principais

#### 1. **categories** - Categorias de Produtos
```sql
- id (SERIAL PRIMARY KEY)
- name (VARCHAR(100)) - Nome da categoria
- description (TEXT) - Descrição da categoria
- created_at (TIMESTAMP) - Data de criação
```

#### 2. **products** - Produtos do Menu
```sql
- id (SERIAL PRIMARY KEY)
- name (VARCHAR(200)) - Nome do produto
- description (TEXT) - Descrição do produto
- price (DECIMAL(10,2)) - Preço
- category_id (INTEGER FK) - FK para categories
- image_url (VARCHAR(500)) - URL da imagem
- is_available (BOOLEAN) - Se está disponível
- created_at (TIMESTAMP) - Data de criação
```

#### 3. **ingredients** - Ingredientes para Montagem
```sql
- id (SERIAL PRIMARY KEY)
- name (VARCHAR(100)) - Nome do ingrediente
- price (DECIMAL(10,2)) - Preço adicional
- category (VARCHAR(50)) - Tipo: pão, carne, queijo, molhos
- is_available (BOOLEAN) - Se está disponível
- created_at (TIMESTAMP) - Data de criação
```

#### 4. **orders** - Pedidos
```sql
- id (SERIAL PRIMARY KEY)
- customer_name (VARCHAR(200)) - Nome do cliente
- table_number (INTEGER) - Número da mesa
- total_amount (DECIMAL(10,2)) - Valor total
- status (VARCHAR(50)) - Status: pending, preparing, ready, delivered
- notes (TEXT) - Observações do pedido
- created_at (TIMESTAMP) - Data de criação
- updated_at (TIMESTAMP) - Data de atualização
```

#### 5. **order_items** - Itens dos Pedidos
```sql
- id (SERIAL PRIMARY KEY)
- order_id (INTEGER FK) - FK para orders
- product_id (INTEGER FK) - FK para products
- ingredients (TEXT) - JSON com ingredientes customizados
- quantity (INTEGER) - Quantidade
- unit_price (DECIMAL(10,2)) - Preço unitário
- total_price (DECIMAL(10,2)) - Preço total do item
- notes (TEXT) - Observações do item
- created_at (TIMESTAMP) - Data de criação
```

## 🔌 API Endpoints

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

## 🔄 Fluxo de Dados

### 1. **Carregamento Inicial**
```
Frontend → Backend API → PostgreSQL → Dados → Frontend
```

### 2. **Atualização de Status**
```
Cozinha → Backend API → UPDATE SQL → PostgreSQL → Confirmação → Cozinha
```

### 3. **Sincronização Automática**
```
Cozinha → Backend API (a cada 10s) → PostgreSQL → Atualização → Cozinha
```

## 🎨 Design System

### Paleta de Cores
- **Primárias**: `#1a1a1a`, `#ffffff`
- **Destaque**: `#ff6b35` (laranja), `#f7931e` (dourado)
- **Sucesso**: `#10b981` (verde)
- **Informação**: `#3b82f6` (azul)
- **Texto**: `#2d2d2d` (escuro), `#6b7280` (claro)

### Tipografia
- **Fonte**: Inter (Google Fonts)
- **Pesos**: 400, 500, 600, 700
- **Tamanhos**: 14px, 16px, 18px, 24px, 30px

## 📱 Responsividade

### Breakpoints
- **Desktop**: > 768px
- **Tablet**: 768px - 1024px
- **Mobile**: < 768px

## 🔒 Segurança

### Frontend
- Validação de dados de entrada
- Sanitização de parâmetros
- CORS configurado

### Backend
- Validação de dados de entrada
- Prepared Statements (SQL injection)
- Transações SQL para consistência
- Middleware CORS configurado

## 🛠️ Desenvolvimento

### Padrões Utilizados

#### Frontend (Vue.js 3)
- **Composition API** para lógica reativa
- **Props** para comunicação pai-filho
- **Emits** para comunicação filho-pai
- **Computed** para valores derivados
- **Lifecycle hooks** para efeitos colaterais

#### Backend (Go)
- **Dependency Injection** - Passagem de `db` para handlers
- **Error Handling** - Tratamento consistente de erros
- **JSON Tags** - Serialização automática
- **SQL Prepared Statements** - Prevenção de SQL injection

## 🚀 Deploy

### Frontend
```bash
cd projeto-hamburgueria
npm run build
# Arquivos gerados em dist/
```

### Backend
```bash
cd backend-hamburgueria
go build -o backend-hamburgueria
# Binário executável gerado
```

### Docker 
```dockerfile
# Frontend
FROM node:16-alpine
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# Backend
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

#### Frontend (.env)
```bash
VITE_API_URL=http://localhost:8080/api
```

#### Backend (.env)
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

## 📚 Recursos de Aprendizado

### Frontend
- [Vue.js 3](https://vuejs.org/) - Documentação oficial
- [Composition API](https://vuejs.org/guide/extras/composition-api-faq.html) - Guia da Composition API
- [Vite](https://vitejs.dev/) - Build tool moderno

### Backend
- [Go](https://golang.org/) - Linguagem de programação
- [Gin Framework](https://gin-gonic.com/) - Framework web
- [PostgreSQL](https://www.postgresql.org/docs/) - Banco de dados


## 🤝 Contribuição

### Como Contribuir
1. Fork o projeto
2. Crie uma branch para sua feature
3. Commit suas mudanças
4. Push para a branch
5. Abra um Pull Request

### Padrões de Código
- Use ESLint e Prettier (frontend)
- Use `gofmt` (backend)
- Siga os guias de estilo das tecnologias
- Escreva testes quando possível
- Documente mudanças importantes


## 👥 Autores

- **Desenvolvedor**: Thiago Matos Tertuliano
- **Data**: Julho 2025
- **Versão**: 1.0.0

---

## 📖 Documentação Detalhada

- **[Frontend README](./projeto-hamburgueria/README.md)** - Documentação completa do frontend
- **[Backend README](./backend-hamburgueria/README.md)** - Documentação completa do backend
- **[Estrutura do Banco](./backend-hamburgueria/ESTRUTURA_BANCO.md)** - Documentação do banco de dados

---

**🍔 BurgerApp** - Sistema completo e moderno para transformar a experiência de pedidos de hambúrgueres! 

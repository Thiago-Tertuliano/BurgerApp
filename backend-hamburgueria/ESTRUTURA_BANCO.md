# Estrutura do Banco de Dados - Hamburgueria

## 📊 Tabelas Principais

### 1. **categories** - Categorias de Produtos
```sql
- id (SERIAL PRIMARY KEY)
- name (VARCHAR(100)) - Nome da categoria
- description (TEXT) - Descrição da categoria
- created_at (TIMESTAMP)
```

**Dados baseados no Menu.vue:**
- Burgers (Hambúrgueres tradicionais e especiais)
- Bebidas (Refrigerantes, sucos e outras bebidas)
- Acompanhamentos (Batatas, saladas e outros acompanhamentos)
- Sobremesas (Sobremesas deliciosas)

### 2. **products** - Produtos do Menu
```sql
- id (SERIAL PRIMARY KEY)
- name (VARCHAR(200)) - Nome do produto
- description (TEXT) - Descrição do produto
- price (DECIMAL(10,2)) - Preço
- category_id (INTEGER) - FK para categories
- image_url (VARCHAR(500)) - URL da imagem
- is_available (BOOLEAN) - Se está disponível
- created_at (TIMESTAMP)
```

**Dados baseados no Menu.vue:**
- Classic Burger (R$ 24.90)
- Bacon Deluxe (R$ 29.90)
- Refrigerante (R$ 6.00)
- Batata Frita (R$ 12.90)

### 3. **ingredients** - Ingredientes para Montagem Personalizada
```sql
- id (SERIAL PRIMARY KEY)
- name (VARCHAR(100)) - Nome do ingrediente
- price (DECIMAL(10,2)) - Preço adicional
- category (VARCHAR(50)) - Tipo: pão, carne, queijo, vegetais, molhos
- is_available (BOOLEAN) - Se está disponível
- created_at (TIMESTAMP)
```

**Dados baseados no CustomBurger.vue:**

**Pães:**
- Australiano (R$ 5.00)
- Brioche (R$ 4.00)
- Tradicional (R$ 3.00)

**Carnes:**
- Angus 120g (R$ 12.00)
- Frango Grelhado (R$ 10.00)
- Veggie (R$ 9.00)

**Queijos:**
- Cheddar (R$ 3.00)
- Mussarela (R$ 2.50)
- Sem queijo (R$ 0.00)

**Vegetais:**
- Alface (R$ 1.00)
- Tomate (R$ 1.00)
- Cebola Crispy (R$ 2.00)
- Picles (R$ 1.50)

**Molhos:**
- Barbecue (R$ 2.00)
- Maionese da Casa (R$ 1.50)
- Sem molho (R$ 0.00)

### 4. **orders** - Pedidos
```sql
- id (SERIAL PRIMARY KEY)
- customer_name (VARCHAR(200)) - Nome do cliente
- table_number (INTEGER) - Número da mesa
- total_amount (DECIMAL(10,2)) - Valor total
- status (VARCHAR(50)) - Status: pending, preparing, ready, delivered
- notes (TEXT) - Observações do pedido
- created_at (TIMESTAMP)
- updated_at (TIMESTAMP)
```

**Status dos pedidos (baseados no Kitchen.vue):**
- `pending` - Pendente
- `preparing` - Em preparação
- `ready` - Pronto para entrega
- `delivered` - Entregue

### 5. **order_items** - Itens dos Pedidos
```sql
- id (SERIAL PRIMARY KEY)
- order_id (INTEGER) - FK para orders
- product_id (INTEGER) - FK para products
- ingredients (TEXT) - JSON com ingredientes customizados
- quantity (INTEGER) - Quantidade
- unit_price (DECIMAL(10,2)) - Preço unitário
- total_price (DECIMAL(10,2)) - Preço total do item
- notes (TEXT) - Observações do item
- created_at (TIMESTAMP)
```

## 🔄 Relacionamentos

```
categories (1) ←→ (N) products
orders (1) ←→ (N) order_items
products (1) ←→ (N) order_items
```

## 📝 Dados Mockados Substituídos

### Frontend → Backend

**Menu.vue:**
- `categories` array → tabela `categories`
- `products` array → tabela `products`

**CustomBurger.vue:**
- `options` object → tabela `ingredients`
- Cálculo de preço → lógica no backend

**App.vue:**
- `orders` array → tabela `orders`
- Status management → endpoints de atualização

**Kitchen.vue:**
- Status display → dados do banco
- Actions → endpoints PUT /orders/:id/status

## 🚀 Endpoints Necessários

### Produtos e Categorias
- `GET /api/products` - Listar produtos
- `GET /api/categories` - Listar categorias

### Ingredientes
- `GET /api/ingredients` - Listar ingredientes para montagem

### Pedidos
- `GET /api/orders` - Listar pedidos
- `GET /api/orders?status=preparing` - Filtrar por status
- `POST /api/orders` - Criar pedido
- `PUT /api/orders/:id/status` - Atualizar status
- `GET /api/orders/:id` - Detalhes do pedido

## 💡 Benefícios da Migração

1. **Persistência de dados** - Pedidos não se perdem ao recarregar
2. **Histórico completo** - Todos os pedidos ficam registrados
3. **Múltiplos usuários** - Vários atendentes podem usar simultaneamente
4. **Relatórios** - Possibilidade de gerar relatórios de vendas
5. **Escalabilidade** - Fácil adicionar novas funcionalidades
6. **Backup** - Dados seguros e recuperáveis

## 🔧 Scripts de Migração

- `scripts/seed.sql` - Dados de exemplo baseados no frontend
- `database/database.go` - Criação automática das tabelas
- `handlers/handlers.go` - Lógica de negócio
- `routes/routes.go` - Endpoints da API 
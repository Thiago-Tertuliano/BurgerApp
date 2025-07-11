// Pacote de conexão e configuração do banco de dados
package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// Driver PostgreSQL para Go - permite conectar com PostgreSQL
	_ "github.com/lib/pq"
)

// Connect estabelece conexão com o PostgreSQL
// Retorna uma conexão *sql.DB e um erro se houver problemas
func Connect() (*sql.DB, error) {
	// ===== CONFIGURAÇÃO DE CONEXÃO =====
	// Obter configurações do banco de dados das variáveis de ambiente
	// Se não estiverem definidas, usar valores padrão

	// Host do banco de dados
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	// Porta do banco de dados
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	// Usuário do banco de dados
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "postgres"
	}

	// Senha do banco de dados
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "postgres"
	}

	// Nome do banco de dados
	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "hamburgueria"
	}

	// ===== STRING DE CONEXÃO =====
	// Construir string de conexão PostgreSQL
	// Formato: host=localhost port=5432 user=postgres password=postgres dbname=hamburgueria sslmode=disable
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// ===== ESTABELECER CONEXÃO =====
	// Abrir conexão com o banco de dados
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	// ===== TESTAR CONEXÃO =====
	// Verificar se a conexão está funcionando
	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Conectado ao PostgreSQL com sucesso!")

	// ===== CRIAR TABELAS =====
	// Criar tabelas se não existirem
	// Isso garante que a aplicação funcione mesmo em um banco vazio
	if err := createTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

// createTables cria as tabelas necessárias no banco de dados
// Usa CREATE TABLE IF NOT EXISTS para não sobrescrever tabelas existentes
func createTables(db *sql.DB) error {
	// ===== TABELA CATEGORIAS =====
	// Armazena categorias de produtos (Burgers, Bebidas, etc.)
	categoriesTable := `
	CREATE TABLE IF NOT EXISTS categories (
		id SERIAL PRIMARY KEY,                    -- ID único auto-incremento
		name VARCHAR(100) NOT NULL,               -- Nome da categoria
		description TEXT,                         -- Descrição da categoria
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- Data de criação
	);`

	// ===== TABELA PRODUTOS =====
	// Armazena produtos do menu (hambúrgueres, bebidas, etc.)
	productsTable := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,                    -- ID único auto-incremento
		name VARCHAR(200) NOT NULL,               -- Nome do produto
		description TEXT,                         -- Descrição do produto
		price DECIMAL(10,2) NOT NULL,            -- Preço com 2 casas decimais
		category_id INTEGER REFERENCES categories(id),  -- FK para categoria
		image_url VARCHAR(500),                   -- URL da imagem do produto
		is_available BOOLEAN DEFAULT true,        -- Se o produto está disponível
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- Data de criação
	);`

	// ===== TABELA INGREDIENTES =====
	// Armazena ingredientes para montagem de lanches personalizados
	ingredientsTable := `
	CREATE TABLE IF NOT EXISTS ingredients (
		id SERIAL PRIMARY KEY,                    -- ID único auto-incremento
		name VARCHAR(100) NOT NULL,               -- Nome do ingrediente
		price DECIMAL(10,2) NOT NULL,            -- Preço adicional do ingrediente
		category VARCHAR(50) NOT NULL,           -- Tipo: pão, carne, queijo, molho
		is_available BOOLEAN DEFAULT true,        -- Se o ingrediente está disponível
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- Data de criação
	);`

	// ===== TABELA PEDIDOS =====
	// Armazena os pedidos dos clientes
	ordersTable := `
	CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,                    -- ID único auto-incremento
		customer_name VARCHAR(200),               -- Nome do cliente
		table_number INTEGER,                     -- Número da mesa
		total_amount DECIMAL(10,2) NOT NULL,     -- Valor total do pedido
		status VARCHAR(50) DEFAULT 'pending',    -- Status: pending, preparing, ready, delivered
		notes TEXT,                              -- Observações do pedido
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Data de criação
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP   -- Data de atualização
	);`

	// ===== TABELA ITENS DO PEDIDO =====
	// Armazena os itens de cada pedido
	orderItemsTable := `
	CREATE TABLE IF NOT EXISTS order_items (
		id SERIAL PRIMARY KEY,                    -- ID único auto-incremento
		order_id INTEGER REFERENCES orders(id) ON DELETE CASCADE,  -- FK para pedido
		product_id INTEGER REFERENCES products(id),  -- FK para produto
		ingredients TEXT,                         -- JSON com ingredientes customizados
		quantity INTEGER NOT NULL,                -- Quantidade do item
		unit_price DECIMAL(10,2) NOT NULL,       -- Preço unitário
		total_price DECIMAL(10,2) NOT NULL,      -- Preço total do item
		notes TEXT,                              -- Observações do item
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- Data de criação
	);`

	// ===== EXECUTAR CRIAÇÃO DAS TABELAS =====
	// Array com todas as queries de criação
	tables := []string{categoriesTable, productsTable, ingredientsTable, ordersTable, orderItemsTable}

	// Executar cada query de criação
	for _, table := range tables {
		if _, err := db.Exec(table); err != nil {
			return fmt.Errorf("erro ao criar tabela: %v", err)
		}
	}

	log.Println("Tabelas criadas/verificadas com sucesso!")
	return nil
}

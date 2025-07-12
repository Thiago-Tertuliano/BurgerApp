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

	return db, nil
}


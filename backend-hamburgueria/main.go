// Pacote principal da aplicação
package main

import (
	"log"
	"os"

	// Framework web Gin para criar a API REST
	"github.com/gin-gonic/gin"
	// Biblioteca para carregar variáveis de ambiente do arquivo .env
	"github.com/joho/godotenv"
	// Pacotes internos do projeto
	"backend-hamburgueria/config"
	"backend-hamburgueria/database"
	"backend-hamburgueria/routes"
)

// Função principal - ponto de entrada da aplicação
func main() {
	// ===== CONFIGURAÇÃO INICIAL =====
	// Carregar variáveis de ambiente do arquivo .env
	// Se o arquivo não existir, usa as variáveis do sistema
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}

	// ===== CONEXÃO COM BANCO DE DADOS =====
	// Estabelecer conexão com o PostgreSQL
	// Se houver erro, a aplicação para (log.Fatal)
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Erro ao conectar com banco de dados:", err)
	}
	// Garantir que a conexão seja fechada ao final da aplicação
	defer db.Close()

	// ===== CONFIGURAÇÃO DO SERVIDOR WEB =====
	// Definir modo de produção para o Gin (desabilita debug)
	gin.SetMode(gin.ReleaseMode)
	// Criar instância do Gin com middlewares padrão
	r := gin.Default()

	// ===== MIDDLEWARES =====
	// Configurar CORS para permitir requisições do frontend
	// Permite que o frontend (localhost:5173) acesse a API
	r.Use(config.CORS())

	// ===== CONFIGURAÇÃO DAS ROTAS =====
	// Configurar todas as rotas da API
	// Passa a instância do Gin e a conexão com o banco
	routes.SetupRoutes(r, db)

	// ===== INICIALIZAÇÃO DO SERVIDOR =====
	// Obter porta do ambiente ou usar 8080 como padrão
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Log informativo sobre o servidor
	log.Printf("Servidor rodando na porta %s", port)
	// Iniciar o servidor HTTP na porta especificada
	// Se houver erro, a aplicação para (log.Fatal)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}

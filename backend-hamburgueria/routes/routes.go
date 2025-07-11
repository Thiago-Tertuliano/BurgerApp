// Pacote de configuração das rotas da API
package routes

import (
	"database/sql"

	// Handlers (manipuladores) das requisições
	"backend-hamburgueria/handlers"

	// Framework web Gin
	"github.com/gin-gonic/gin"
)

// SetupRoutes configura todas as rotas da API
// Recebe a instância do Gin e a conexão com o banco de dados
func SetupRoutes(r *gin.Engine, db *sql.DB) {
	// ===== GRUPO DE ROTAS DA API =====
	// Todas as rotas da API começam com /api
	api := r.Group("/api")
	{
		// ===== ROTAS DE PRODUTOS =====
		// GET /api/products - Listar todos os produtos disponíveis
		api.GET("/products", func(c *gin.Context) {
			handlers.GetProducts(c, db)
		})

		// ===== ROTAS DE CATEGORIAS =====
		// GET /api/categories - Listar todas as categorias
		api.GET("/categories", func(c *gin.Context) {
			handlers.GetCategories(c, db)
		})

		// ===== ROTAS DE INGREDIENTES =====
		// GET /api/ingredients - Listar todos os ingredientes para montagem
		api.GET("/ingredients", func(c *gin.Context) {
			handlers.GetIngredients(c, db)
		})

		// ===== ROTAS DE PEDIDOS =====
		// POST /api/orders - Criar um novo pedido
		api.POST("/orders", func(c *gin.Context) {
			handlers.CreateOrder(c, db)
		})

		// GET /api/orders - Listar todos os pedidos
		// Suporta filtro: GET /api/orders?status=preparing
		api.GET("/orders", func(c *gin.Context) {
			handlers.GetOrders(c, db)
		})

		// GET /api/orders/:id - Obter detalhes de um pedido específico
		api.GET("/orders/:id", func(c *gin.Context) {
			handlers.GetOrderDetails(c, db)
		})

		// PUT /api/orders/:id/status - Atualizar status de um pedido
		// Usado pela cozinha para marcar pedidos como pronto/entregue
		api.PUT("/orders/:id/status", func(c *gin.Context) {
			handlers.UpdateOrderStatus(c, db)
		})
	}

	// ===== ROTA DE HEALTH CHECK =====
	// GET /health - Verificar se o servidor está funcionando
	// Útil para monitoramento e testes
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Backend da Hamburgueria funcionando!",
		})
	})
}

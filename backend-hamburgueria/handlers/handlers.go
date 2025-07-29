// Pacote de handlers (manipuladores) das requisições HTTP
package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	// Modelos de dados da aplicação
	"backend-hamburgueria/models"

	// Framework web Gin
	"github.com/gin-gonic/gin"
)

// DBInterface define a interface para operações de banco de dados
// Isso permite que os testes usem mocks
type DBInterface interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
	Begin() (*sql.Tx, error)
}

// ===== HANDLERS DE PRODUTOS E CATEGORIAS =====

// GetProducts retorna todos os produtos disponíveis
// Endpoint: GET /api/products
func GetProducts(c *gin.Context, db DBInterface) {
	// Query SQL com JOIN para buscar produtos e suas categorias
	query := `
		SELECT p.id, p.name, p.description, p.price, p.category_id, p.image_url, p.is_available, p.created_at,
			   c.id, c.name, c.description, c.created_at
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id
		WHERE p.is_available = true
		ORDER BY p.created_at DESC
	`

	// Executar a query no banco de dados
	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar produtos"})
		return
	}
	defer rows.Close() // Garantir que as linhas sejam fechadas

	// ===== PROCESSAR RESULTADOS =====
	var products []models.Product
	for rows.Next() {
		var p models.Product
		var cat models.Category
		// Ler cada linha do resultado
		err := rows.Scan(
			&p.ID, &p.Name, &p.Description, &p.Price, &p.CategoryID, &p.ImageURL, &p.IsAvailable, &p.CreatedAt,
			&cat.ID, &cat.Name, &cat.Description, &cat.CreatedAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler produto"})
			return
		}
		// Associar categoria ao produto
		p.Category = cat
		products = append(products, p)
	}

	// Retornar produtos como JSON
	c.JSON(http.StatusOK, products)
}

// GetCategories retorna todas as categorias
// Endpoint: GET /api/categories
func GetCategories(c *gin.Context, db DBInterface) {
	// Query SQL para buscar categorias ordenadas por nome
	query := `SELECT id, name, description, created_at FROM categories ORDER BY name`

	// Executar a query
	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar categorias"})
		return
	}
	defer rows.Close()

	// ===== PROCESSAR RESULTADOS =====
	var categories []models.Category
	for rows.Next() {
		var cat models.Category
		// Ler cada linha do resultado
		err := rows.Scan(&cat.ID, &cat.Name, &cat.Description, &cat.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler categoria"})
			return
		}
		categories = append(categories, cat)
	}

	// Retornar categorias como JSON
	c.JSON(http.StatusOK, categories)
}

// GetIngredients retorna todos os ingredientes
// Endpoint: GET /api/ingredients
func GetIngredients(c *gin.Context, db DBInterface) {
	// Query SQL para buscar ingredientes disponíveis ordenados por categoria e nome
	query := `SELECT id, name, price, category, is_available, created_at FROM ingredients WHERE is_available = true ORDER BY category, name`

	// Executar a query
	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar ingredientes"})
		return
	}
	defer rows.Close()

	// ===== PROCESSAR RESULTADOS =====
	var ingredients []models.Ingredient
	for rows.Next() {
		var ing models.Ingredient
		// Ler cada linha do resultado
		err := rows.Scan(&ing.ID, &ing.Name, &ing.Price, &ing.Category, &ing.IsAvailable, &ing.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler ingrediente"})
			return
		}
		ingredients = append(ingredients, ing)
	}

	// Retornar ingredientes como JSON
	c.JSON(http.StatusOK, ingredients)
}

// ===== HANDLERS DE PEDIDOS =====

// CreateOrder cria um novo pedido
// Endpoint: POST /api/orders
func CreateOrder(c *gin.Context, db DBInterface) {
	// ===== VALIDAR DADOS DE ENTRADA =====
	var req models.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// ===== INICIAR TRANSAÇÃO =====
	// Usar transação para garantir consistência dos dados
	tx, err := db.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao iniciar transação"})
		return
	}
	defer tx.Rollback() // Rollback em caso de erro

	// ===== CALCULAR TOTAL DO PEDIDO =====
	var totalAmount float64
	for _, item := range req.Items {
		// Para produtos customizados (ID 1), usar preço calculado no frontend
		if item.ProductID == 1 {
			// Extrair preço dos ingredientes se disponível
			if item.Ingredients != "" {
				// Tentar extrair preço do JSON de ingredientes
				// Por enquanto, usar preço padrão para customizados
				totalAmount += 15.00 // Preço padrão para lanches customizados
			} else {
				totalAmount += 15.00
			}
		} else {
			// Buscar preço do produto no banco
			var productPrice float64
			err := tx.QueryRow("SELECT price FROM products WHERE id = $1", item.ProductID).Scan(&productPrice)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Produto não encontrado"})
				return
			}

			// Calcular preço total do item
			itemTotal := productPrice * float64(item.Quantity)
			totalAmount += itemTotal
		}
	}

	// ===== INSERIR PEDIDO =====
	var orderID int
	err = tx.QueryRow(`
		INSERT INTO orders (customer_name, table_number, total_amount, notes)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`, req.CustomerName, req.TableNumber, totalAmount, req.Notes).Scan(&orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar pedido"})
		return
	}

	// ===== INSERIR ITENS DO PEDIDO =====
	for _, item := range req.Items {
		var productPrice float64
		var totalPrice float64

		if item.ProductID == 1 {
			// Para produtos customizados, usar preço padrão
			productPrice = 15.00
			totalPrice = 15.00
		} else {
			// Buscar preço do produto
			err := tx.QueryRow("SELECT price FROM products WHERE id = $1", item.ProductID).Scan(&productPrice)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Produto não encontrado"})
				return
			}
			totalPrice = productPrice * float64(item.Quantity)
		}

		// Inserir item do pedido
		_, err = tx.Exec(`
			INSERT INTO order_items (order_id, product_id, ingredients, quantity, unit_price, total_price, notes)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
		`, orderID, item.ProductID, item.Ingredients, item.Quantity, productPrice, totalPrice, item.Notes)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao adicionar item ao pedido"})
			return
		}
	}

	// ===== COMMIT DA TRANSAÇÃO =====
	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao finalizar pedido"})
		return
	}

	// Retornar resposta de sucesso
	c.JSON(http.StatusCreated, gin.H{
		"message":      "Pedido criado com sucesso",
		"order_id":     orderID,
		"total_amount": totalAmount,
	})
}

// GetOrders retorna todos os pedidos
// Endpoint: GET /api/orders
// Suporta filtro por status: GET /api/orders?status=preparing
func GetOrders(c *gin.Context, db DBInterface) {
	// Obter parâmetro de status da query string
	status := c.Query("status")

	// ===== CONSTRUIR QUERY DINÂMICA =====
	var query string
	var args []interface{}

	if status != "" {
		// Query com filtro por status
		query = `SELECT id, customer_name, table_number, total_amount, status, notes, created_at, updated_at FROM orders WHERE status = $1 ORDER BY created_at DESC`
		args = append(args, status)
	} else {
		// Query sem filtro - todos os pedidos
		query = `SELECT id, customer_name, table_number, total_amount, status, notes, created_at, updated_at FROM orders ORDER BY created_at DESC`
	}

	// ===== EXECUTAR QUERY =====
	rows, err := db.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar pedidos"})
		return
	}
	defer rows.Close()

	// ===== PROCESSAR RESULTADOS =====
	var orders []models.Order
	for rows.Next() {
		var order models.Order
		// Ler cada linha do resultado
		err := rows.Scan(&order.ID, &order.CustomerName, &order.TableNumber, &order.TotalAmount, &order.Status, &order.Notes, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler pedido"})
			return
		}
		orders = append(orders, order)
	}

	// Retornar pedidos como JSON
	c.JSON(http.StatusOK, orders)
}

// UpdateOrderStatus atualiza o status de um pedido
// Endpoint: PUT /api/orders/:id/status
func UpdateOrderStatus(c *gin.Context, db DBInterface) {
	// ===== VALIDAR ID DO PEDIDO =====
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do pedido inválido"})
		return
	}

	// ===== VALIDAR DADOS DE ENTRADA =====
	var req models.UpdateOrderStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status inválido"})
		return
	}

	// ===== VALIDAR STATUS =====
	// Lista de status válidos
	validStatuses := []string{"pending", "preparing", "ready", "delivered"}
	isValid := false
	for _, status := range validStatuses {
		if req.Status == status {
			isValid = true
			break
		}
	}
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status inválido"})
		return
	}

	// ===== ATUALIZAR STATUS =====
	// Executar UPDATE no banco de dados
	_, err = db.Exec("UPDATE orders SET status = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2", req.Status, orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar status"})
		return
	}

	// Retornar resposta de sucesso
	c.JSON(http.StatusOK, gin.H{"message": "Status atualizado com sucesso"})
}

// GetOrderDetails retorna os detalhes de um pedido específico
// Endpoint: GET /api/orders/:id
func GetOrderDetails(c *gin.Context, db DBInterface) {
	// ===== VALIDAR ID DO PEDIDO =====
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do pedido inválido"})
		return
	}

	// ===== BUSCAR PEDIDO =====
	var order models.Order
	err = db.QueryRow(`
		SELECT id, customer_name, table_number, total_amount, status, notes, created_at, updated_at
		FROM orders WHERE id = $1
	`, orderID).Scan(&order.ID, &order.CustomerName, &order.TableNumber, &order.TotalAmount, &order.Status, &order.Notes, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pedido não encontrado"})
		return
	}

	// ===== BUSCAR ITENS DO PEDIDO =====
	// Query com JOIN para buscar itens e produtos
	rows, err := db.Query(`
		SELECT oi.id, oi.order_id, oi.product_id, oi.ingredients, oi.quantity, oi.unit_price, oi.total_price, oi.notes, oi.created_at,
			   p.id, p.name, p.description, p.price, p.category_id, p.image_url, p.is_available, p.created_at
		FROM order_items oi
		LEFT JOIN products p ON oi.product_id = p.id
		WHERE oi.order_id = $1
	`, orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar itens do pedido"})
		return
	}
	defer rows.Close()

	// ===== PROCESSAR ITENS =====
	for rows.Next() {
		var item models.OrderItem
		var product models.Product
		// Ler cada linha do resultado
		err := rows.Scan(
			&item.ID, &item.OrderID, &item.ProductID, &item.Ingredients, &item.Quantity, &item.UnitPrice, &item.TotalPrice, &item.Notes, &item.CreatedAt,
			&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL, &product.IsAvailable, &product.CreatedAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler item do pedido"})
			return
		}
		// Associar produto ao item
		item.Product = product
		order.Items = append(order.Items, item)
	}

	// Retornar pedido completo como JSON
	c.JSON(http.StatusOK, order)
}

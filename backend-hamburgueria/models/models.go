// Pacote de modelos de dados da aplicação
package models

import (
	"time"
)

// ===== MODELOS PRINCIPAIS =====

// Category representa uma categoria de produtos
// Exemplo: Burgers, Bebidas, Acompanhamentos, Sobremesas
type Category struct {
	ID          int       `json:"id"`          // ID único da categoria
	Name        string    `json:"name"`        // Nome da categoria
	Description string    `json:"description"` // Descrição da categoria
	CreatedAt   time.Time `json:"created_at"`  // Data de criação
}

// Product representa um produto do menu
// Exemplo: Classic Burger, Bacon Deluxe, Refrigerante
type Product struct {
	ID          int       `json:"id"`                 // ID único do produto
	Name        string    `json:"name"`               // Nome do produto
	Description string    `json:"description"`        // Descrição do produto
	Price       float64   `json:"price"`              // Preço do produto
	CategoryID  int       `json:"category_id"`        // ID da categoria (FK)
	Category    Category  `json:"category,omitempty"` // Categoria completa (opcional)
	ImageURL    string    `json:"image_url"`          // URL da imagem do produto
	IsAvailable bool      `json:"is_available"`       // Se o produto está disponível
	CreatedAt   time.Time `json:"created_at"`         // Data de criação
}

// Ingredient representa um ingrediente para montagem de lanches
// Exemplo: Pão Australiano, Carne Angus, Queijo Cheddar
type Ingredient struct {
	ID          int       `json:"id"`           // ID único do ingrediente
	Name        string    `json:"name"`         // Nome do ingrediente
	Price       float64   `json:"price"`        // Preço adicional do ingrediente
	Category    string    `json:"category"`     // Tipo: pão, carne, queijo, vegetais, molhos
	IsAvailable bool      `json:"is_available"` // Se o ingrediente está disponível
	CreatedAt   time.Time `json:"created_at"`   // Data de criação
}

// Order representa um pedido
// Contém informações do cliente e status do pedido
type Order struct {
	ID           int         `json:"id"`              // ID único do pedido
	CustomerName string      `json:"customer_name"`   // Nome do cliente
	TableNumber  int         `json:"table_number"`    // Número da mesa
	TotalAmount  float64     `json:"total_amount"`    // Valor total do pedido
	Status       string      `json:"status"`          // Status: pending, preparing, ready, delivered
	Notes        string      `json:"notes"`           // Observações do pedido
	CreatedAt    time.Time   `json:"created_at"`      // Data de criação
	UpdatedAt    time.Time   `json:"updated_at"`      // Data de última atualização
	Items        []OrderItem `json:"items,omitempty"` // Itens do pedido (opcional)
}

// OrderItem representa um item de um pedido
// Cada pedido pode ter múltiplos itens
type OrderItem struct {
	ID          int       `json:"id"`                // ID único do item
	OrderID     int       `json:"order_id"`          // ID do pedido (FK)
	ProductID   int       `json:"product_id"`        // ID do produto (FK)
	Product     Product   `json:"product,omitempty"` // Produto completo (opcional)
	Ingredients string    `json:"ingredients"`       // JSON string com ingredientes customizados
	Quantity    int       `json:"quantity"`          // Quantidade do item
	UnitPrice   float64   `json:"unit_price"`        // Preço unitário
	TotalPrice  float64   `json:"total_price"`       // Preço total do item
	Notes       string    `json:"notes"`             // Observações do item
	CreatedAt   time.Time `json:"created_at"`        // Data de criação
}

// ===== MODELOS DE REQUISIÇÃO =====

// CreateOrderRequest representa a requisição para criar um pedido
// Usado quando o frontend envia dados para criar um novo pedido
type CreateOrderRequest struct {
	CustomerName string             `json:"customer_name"` // Nome do cliente
	TableNumber  int                `json:"table_number"`  // Número da mesa
	Items        []OrderItemRequest `json:"items"`         // Lista de itens do pedido
	Notes        string             `json:"notes"`         // Observações do pedido
}

// OrderItemRequest representa um item de pedido na requisição
// Usado dentro de CreateOrderRequest para especificar os itens
type OrderItemRequest struct {
	ProductID   int    `json:"product_id"`  // ID do produto
	Ingredients string `json:"ingredients"` // JSON string com ingredientes customizados
	Quantity    int    `json:"quantity"`    // Quantidade
	Notes       string `json:"notes"`       // Observações do item
}

// UpdateOrderStatusRequest representa a requisição para atualizar status do pedido
// Usado quando a cozinha atualiza o status de um pedido
type UpdateOrderStatusRequest struct {
	Status string `json:"status"` // Novo status: pending, preparing, ready, delivered
}
type StatusResponse struct {
		Message string `json: "message"`
}

type ErrorResponse struct {
		Error string `json: "error"`
}

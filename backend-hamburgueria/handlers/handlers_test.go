package handlers

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// MockDB simula o comportamento do banco de dados para testes
// Implementa a interface DBInterface
type MockDB struct {
	// Mock para Query
	QueryFunc func(query string, args ...interface{}) (*sql.Rows, error)
	// Mock para QueryRow
	QueryRowFunc func(query string, args ...interface{}) *sql.Row
	// Mock para Exec
	ExecFunc func(query string, args ...interface{}) (sql.Result, error)
	// Mock para Begin
	BeginFunc func() (*sql.Tx, error)
}

// Query implementa a interface DBInterface
func (m *MockDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if m.QueryFunc != nil {
		return m.QueryFunc(query, args...)
	}
	return nil, nil
}

// QueryRow implementa a interface DBInterface
func (m *MockDB) QueryRow(query string, args ...interface{}) *sql.Row {
	if m.QueryRowFunc != nil {
		return m.QueryRowFunc(query, args...)
	}
	return nil
}

// Exec implementa a interface DBInterface
func (m *MockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	if m.ExecFunc != nil {
		return m.ExecFunc(query, args...)
	}
	return nil, nil
}

// Begin implementa a interface DBInterface
func (m *MockDB) Begin() (*sql.Tx, error) {
	if m.BeginFunc != nil {
		return m.BeginFunc()
	}
	return nil, nil
}

// setupTest configura o router e mock DB para testes
func setupTest() (*gin.Engine, *MockDB) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	mockDB := &MockDB{}
	return router, mockDB
}

// Teste para GetProducts
func TestGetProducts(t *testing.T) {
	router, mockDB := setupTest()

	// Configurar rota
	router.GET("/products", func(c *gin.Context) {
		GetProducts(c, mockDB)
	})

	// Mock da resposta do banco - retornar erro para simular falha
	mockDB.QueryFunc = func(query string, args ...interface{}) (*sql.Rows, error) {
		return nil, sql.ErrNoRows
	}

	// Criar requisição
	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()

	// Executar requisição
	router.ServeHTTP(w, req)

	// Verificar resposta - esperamos erro interno devido ao mock
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// Teste para GetCategories
func TestGetCategories(t *testing.T) {
	router, mockDB := setupTest()

	// Configurar rota
	router.GET("/categories", func(c *gin.Context) {
		GetCategories(c, mockDB)
	})

	// Mock da resposta do banco - retornar erro para simular falha
	mockDB.QueryFunc = func(query string, args ...interface{}) (*sql.Rows, error) {
		return nil, sql.ErrNoRows
	}

	// Criar requisição
	req, _ := http.NewRequest("GET", "/categories", nil)
	w := httptest.NewRecorder()

	// Executar requisição
	router.ServeHTTP(w, req)

	// Verificar resposta - esperamos erro interno devido ao mock
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// Teste para GetIngredients
func TestGetIngredients(t *testing.T) {
	router, mockDB := setupTest()

	// Configurar rota
	router.GET("/ingredients", func(c *gin.Context) {
		GetIngredients(c, mockDB)
	})

	// Mock da resposta do banco - retornar erro para simular falha
	mockDB.QueryFunc = func(query string, args ...interface{}) (*sql.Rows, error) {
		return nil, sql.ErrNoRows
	}

	// Criar requisição
	req, _ := http.NewRequest("GET", "/ingredients", nil)
	w := httptest.NewRecorder()

	// Executar requisição
	router.ServeHTTP(w, req)

	// Verificar resposta - esperamos erro interno devido ao mock
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// Teste para GetOrders
func TestGetOrders(t *testing.T) {
	router, mockDB := setupTest()

	// Configurar rota
	router.GET("/orders", func(c *gin.Context) {
		GetOrders(c, mockDB)
	})

	// Mock da resposta do banco - retornar erro para simular falha
	mockDB.QueryFunc = func(query string, args ...interface{}) (*sql.Rows, error) {
		return nil, sql.ErrNoRows
	}

	// Criar requisição
	req, _ := http.NewRequest("GET", "/orders", nil)
	w := httptest.NewRecorder()

	// Executar requisição
	router.ServeHTTP(w, req)

	// Verificar resposta - esperamos erro interno devido ao mock
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// Teste para CreateOrder
func TestCreateOrder(t *testing.T) {
	router, mockDB := setupTest()

	// Configurar rota
	router.POST("/orders", func(c *gin.Context) {
		CreateOrder(c, mockDB)
	})

	// Mock das respostas do banco - retornar erro para simular falha
	mockDB.BeginFunc = func() (*sql.Tx, error) {
		return nil, sql.ErrConnDone
	}

	// Criar requisição
	req, _ := http.NewRequest("POST", "/orders", nil)
	w := httptest.NewRecorder()

	// Executar requisição
	router.ServeHTTP(w, req)

	// Verificar resposta - esperamos Bad Request porque não enviamos dados JSON válidos
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// Teste para UpdateOrderStatus
func TestUpdateOrderStatus(t *testing.T) {
	router, mockDB := setupTest()

	// Configurar rota
	router.PUT("/orders/:id/status", func(c *gin.Context) {
		UpdateOrderStatus(c, mockDB)
	})

	// Mock da resposta do banco - retornar erro para simular falha
	mockDB.ExecFunc = func(query string, args ...interface{}) (sql.Result, error) {
		return nil, sql.ErrConnDone
	}

	// Criar requisição
	req, _ := http.NewRequest("PUT", "/orders/1/status", nil)
	w := httptest.NewRecorder()

	// Executar requisição
	router.ServeHTTP(w, req)

	// Verificar resposta - esperamos Bad Request porque não enviamos dados JSON válidos
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

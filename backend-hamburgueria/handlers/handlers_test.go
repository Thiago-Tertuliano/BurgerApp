package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// Mock para o banco de dados
type MockDB struct {
	mock.Mock
}

func (m *MockDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Where(query interface{}, args ...interface{}) *gorm.DB {
	mockArgs := m.Called(query, args)
	return mockArgs.Get(0).(*gorm.DB)
}

func (m *MockDB) First(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Save(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

// Setup do teste
func setupTest() (*gin.Engine, *MockDB) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	mockDB := &MockDB{}
	return router, mockDB
}

// Teste para GetProducts
func TestGetProducts(t *testing.T) {
	router, mockDB := setupTest()
	
	// Mock do comportamento esperado
	mockDB.On("Find", mock.AnythingOfType("*[]models.Product")).Return(&gorm.DB{})
	
	// Configurar rota
	router.GET("/api/products", func(c *gin.Context) {
		GetProducts(c, mockDB)
	})
	
	// Criar request
	req, _ := http.NewRequest("GET", "/api/products", nil)
	w := httptest.NewRecorder()
	
	// Executar request
	router.ServeHTTP(w, req)
	
	// Verificações
	assert.Equal(t, http.StatusOK, w.Code)
	mockDB.AssertExpectations(t)
}

// Teste para GetCategories
func TestGetCategories(t *testing.T) {
	router, mockDB := setupTest()
	
	// Mock do comportamento esperado
	mockDB.On("Find", mock.AnythingOfType("*[]models.Category")).Return(&gorm.DB{})
	
	// Configurar rota
	router.GET("/api/categories", func(c *gin.Context) {
		GetCategories(c, mockDB)
	})
	
	// Criar request
	req, _ := http.NewRequest("GET", "/api/categories", nil)
	w := httptest.NewRecorder()
	
	// Executar request
	router.ServeHTTP(w, req)
	
	// Verificações
	assert.Equal(t, http.StatusOK, w.Code)
	mockDB.AssertExpectations(t)
}

// Teste para GetIngredients
func TestGetIngredients(t *testing.T) {
	router, mockDB := setupTest()
	
	// Mock do comportamento esperado
	mockDB.On("Find", mock.AnythingOfType("*[]models.Ingredient")).Return(&gorm.DB{})
	
	// Configurar rota
	router.GET("/api/ingredients", func(c *gin.Context) {
		GetIngredients(c, mockDB)
	})
	
	// Criar request
	req, _ := http.NewRequest("GET", "/api/ingredients", nil)
	w := httptest.NewRecorder()
	
	// Executar request
	router.ServeHTTP(w, req)
	
	// Verificações
	assert.Equal(t, http.StatusOK, w.Code)
	mockDB.AssertExpectations(t)
}

// Teste para GetOrders
func TestGetOrders(t *testing.T) {
	router, mockDB := setupTest()
	
	// Mock do comportamento esperado
	mockDB.On("Find", mock.AnythingOfType("*[]models.Order")).Return(&gorm.DB{})
	
	// Configurar rota
	router.GET("/api/orders", func(c *gin.Context) {
		GetOrders(c, mockDB)
	})
	
	// Criar request
	req, _ := http.NewRequest("GET", "/api/orders", nil)
	w := httptest.NewRecorder()
	
	// Executar request
	router.ServeHTTP(w, req)
	
	// Verificações
	assert.Equal(t, http.StatusOK, w.Code)
	mockDB.AssertExpectations(t)
}

// Teste para CreateOrder
func TestCreateOrder(t *testing.T) {
	router, mockDB := setupTest()
	
	// Dados de teste
	orderData := map[string]interface{}{
		"customer_name": "João Silva",
		"table_number":  5,
		"total_amount":  25.50,
		"notes":        "Sem cebola",
		"items": []map[string]interface{}{
			{
				"product_id":  1,
				"quantity":    2,
				"unit_price":  12.75,
				"total_price": 25.50,
				"ingredients": "pão, carne, queijo",
			},
		},
	}
	
	// Mock do comportamento esperado
	mockDB.On("Create", mock.AnythingOfType("*models.Order")).Return(&gorm.DB{})
	
	// Configurar rota
	router.POST("/api/orders", func(c *gin.Context) {
		CreateOrder(c, mockDB)
	})
	
	// Criar request
	jsonData, _ := json.Marshal(orderData)
	req, _ := http.NewRequest("POST", "/api/orders", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	
	// Executar request
	router.ServeHTTP(w, req)
	
	// Verificações
	assert.Equal(t, http.StatusCreated, w.Code)
	mockDB.AssertExpectations(t)
}

// Teste para UpdateOrderStatus
func TestUpdateOrderStatus(t *testing.T) {
	router, mockDB := setupTest()
	
	// Dados de teste
	statusData := map[string]string{
		"status": "preparing",
	}
	
	// Mock do comportamento esperado
	mockDB.On("Where", "id = ?", "1").Return(mockDB)
	mockDB.On("First", mock.AnythingOfType("*models.Order")).Return(&gorm.DB{})
	mockDB.On("Save", mock.AnythingOfType("*models.Order")).Return(&gorm.DB{})
	
	// Configurar rota
	router.PUT("/api/orders/:id/status", func(c *gin.Context) {
		UpdateOrderStatus(c, mockDB)
	})
	
	// Criar request
	jsonData, _ := json.Marshal(statusData)
	req, _ := http.NewRequest("PUT", "/api/orders/1/status", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	
	// Executar request
	router.ServeHTTP(w, req)
	
	// Verificações
	assert.Equal(t, http.StatusOK, w.Code)
	mockDB.AssertExpectations(t)
}

// Teste para HealthCheck
func TestHealthCheck(t *testing.T) {
	router, _ := setupTest()
	
	// Configurar rota
	router.GET("/health", HealthCheck)
	
	// Criar request
	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	
	// Executar request
	router.ServeHTTP(w, req)
	
	// Verificações
	assert.Equal(t, http.StatusOK, w.Code)
	
	// Verificar resposta JSON
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "OK", response["status"])
}
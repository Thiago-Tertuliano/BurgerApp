// Pacote de configurações da aplicação
package config

import (
	"time"

	// Middleware CORS para permitir requisições cross-origin
	"github.com/gin-contrib/cors"
	// Framework web Gin
	"github.com/gin-gonic/gin"
)

// CORS configura o middleware CORS para permitir requisições do frontend
// CORS (Cross-Origin Resource Sharing) é necessário para que o frontend
// possa fazer requisições para a API quando estão em portas diferentes
func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		// Origens permitidas - URLs do frontend que podem acessar a API
		AllowOrigins: []string{
			"http://localhost:5173", // Vite dev server
			"http://localhost:5174", // Vite dev server (porta alternativa)
			"http://localhost:3000", // Outros servidores de desenvolvimento
		},
		// Métodos HTTP permitidos
		AllowMethods: []string{
			"GET",     // Buscar dados
			"POST",    // Criar novos recursos
			"PUT",     // Atualizar recursos existentes
			"DELETE",  // Remover recursos
			"OPTIONS", // Preflight requests do CORS
		},
		// Headers HTTP permitidos
		AllowHeaders: []string{
			"Origin",        // Origem da requisição
			"Content-Type",  // Tipo do conteúdo
			"Accept",        // Tipos aceitos
			"Authorization", // Token de autenticação
		},
		// Headers expostos para o frontend
		ExposeHeaders: []string{
			"Content-Length", // Tamanho do conteúdo
		},
		// Permitir credenciais (cookies, headers de autenticação)
		AllowCredentials: true,
		// Tempo de cache para preflight requests (12 horas)
		MaxAge: 12 * time.Hour,
	})
}

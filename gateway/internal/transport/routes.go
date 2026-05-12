package transport

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/Dukvaha27/flash-score/gateway/internal/transport/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterProxies(g *gin.Engine) {
	urlStr := os.Getenv("GATEWAY_GIN_URL")
	if urlStr == "" {
		log.Fatal("GATEWAY_GIN_URL is not set")
	}

	userServiceURL, err := url.Parse(urlStr)

	if err != nil {
		log.Fatal("invalid user service URL:", err)
	}

	userProxy := httputil.NewSingleHostReverseProxy(userServiceURL)

	public := g.Group("")

	public.Any("/register", func(c *gin.Context) {
		userProxy.ServeHTTP(c.Writer, c.Request)
	})

	public.Any("/login", func(c *gin.Context) {
		userProxy.ServeHTTP(c.Writer, c.Request)
	})

	protected := g.Group("")
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET is not set")
		return
	}
	protected.Use(middlewares.AuthMiddleware(secret))

	protected.Any("/users/me", func(c *gin.Context) {
		userID, ok := c.Get("user_id")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.Request.Header.Set("X-User-ID", fmt.Sprintf("%v", userID))
		userProxy.ServeHTTP(c.Writer, c.Request)
	})

}

package transport

import (
	"fmt"
	"log"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/Dukvaha27/flash-score/gateway/internal/transport/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterProxies(g *gin.Engine) {
	userServiceURL, err := url.Parse(os.Getenv("GATEWAY_GIN_URL"))
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
	protected.Use(middlewares.AuthMiddleware())

	protected.Any("/users/me", func(c *gin.Context) {
		userID, ok := c.Get("user_id")
		if !ok {
			fmt.Println("ошибка получения user_id")
			return
		}
		c.Request.Header.Set("X-User-ID", fmt.Sprintf("%v", userID))
		userProxy.ServeHTTP(c.Writer, c.Request)
	})

}

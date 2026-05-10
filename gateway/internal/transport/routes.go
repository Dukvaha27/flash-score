package transport

import (
	"log"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

func RegisterProxies(g *gin.Engine) {
	userServiceURL, err := url.Parse(os.Getenv("GATEWAY_GIN_URL"))
	if err != nil {
		log.Fatal("invalid user service URL:", err)
	}

	userProxy := httputil.NewSingleHostReverseProxy(userServiceURL)

	g.Any("/user/register", func(c *gin.Context) {
		userProxy.ServeHTTP(c.Writer, c.Request)
	})

}

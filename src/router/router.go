package router

import (
	"net/http"
	"api-server-demo/src/handler/user"
	"api-server-demo/src/router/middleware"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// pprof router
	pprof.Register(g)

	// register
	g.POST("/register", user.Create)
	// api for authentication functionalities
	g.POST("/login", user.Login)

	return g
}
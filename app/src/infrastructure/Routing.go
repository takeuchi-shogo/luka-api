package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/takeuchi-shogo/luka-api/src/interface/controllers/product"
)

type Routing struct {
	// DB   *DB
	Gin  *gin.Engine
	Port string
}

func NewRouting(c *Config) *Routing {
	r := &Routing{
		// DB:   db,
		Gin:  gin.Default(),
		Port: c.Routing.Port,
	}

	r.setRouting()
	return r
}

func (r *Routing) setRouting() {

	usersController := product.NewUsersController()
	v1 := r.Gin.Group("/v1/product")
	{

		// Users
		v1.GET("/users", func(ctx *gin.Context) { usersController.GetList(ctx) })

		v1.GET("/users/:id", func(ctx *gin.Context) { usersController.Get(ctx) })

		// Test用
		v1.GET("/test", func(ctx *gin.Context) { ctx.JSON(200, "testtest") })
	}
}

func (r *Routing) Run(port string) {
	r.Gin.Run(port)
}

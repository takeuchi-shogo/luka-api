package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/takeuchi-shogo/luka-api/src/interface/controllers/product"
)

type Routing struct {
	DB   *DB
	Gin  *gin.Engine
	Port string
}

func NewRouting(c *Config, db *DB) *Routing {
	r := &Routing{
		DB:   db,
		Gin:  gin.Default(),
		Port: c.Routing.Port,
	}

	r.setRouting()
	return r
}

func (r *Routing) setRouting() {

	commentsController := product.NewCommentsController(r.DB)
	threadsController := product.NewThreadsController(r.DB)
	tokensController := product.NewTokensController(r.DB)
	usersController := product.NewUsersController()
	v1 := r.Gin.Group("/v1/product")
	{
		// Comment To Threads
		v1.POST("/comments", func(ctx *gin.Context) { commentsController.Post(ctx) })

		// Threads
		v1.POST("/threads", func(ctx *gin.Context) { threadsController.Post(ctx) })

		// Tokens
		v1.POST("/tokens", func(ctx *gin.Context) { tokensController.Post(ctx) })

		// Users
		v1.GET("/users", func(ctx *gin.Context) { usersController.GetList(ctx) })
		v1.POST("/users", func(ctx *gin.Context) { usersController.Post(ctx) })

		v1.GET("/users/:id", func(ctx *gin.Context) { usersController.Get(ctx) })

		// Test用
		v1.GET("/test", func(ctx *gin.Context) { ctx.JSON(200, "testtest") })
	}
}

func (r *Routing) Run(port string) {
	r.Gin.Run(port)
}

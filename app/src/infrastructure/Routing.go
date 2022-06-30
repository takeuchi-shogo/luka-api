package infrastructure

import "github.com/gin-gonic/gin"

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
	v1 := r.Gin.Group("/v1/product")
	{
		v1.GET("/test", func(ctx *gin.Context) { ctx.JSON(200, "kokokoko") })
	}
}

func (r *Routing) Run(port string) {
	r.Gin.Run(port)
}

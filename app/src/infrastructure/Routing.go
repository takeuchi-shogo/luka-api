package infrastructure

import (
	"github.com/gin-contrib/cors"
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

	r.cors(c)
	r.setRouting()
	return r
}

func (r *Routing) cors(c *Config) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = c.CORS.AllowOrigins
	r.Gin.Use(cors.New(corsConfig))
}

func (r *Routing) setRouting() {

	commentsController := product.NewCommentsController(r.DB)
	followersController := product.NewFollowersController(r.DB)
	followingsController := product.NewFollowingsController(r.DB)
	meController := product.NewMeController(r.DB)
	threadsController := product.NewThreadsController(r.DB)
	tokensController := product.NewTokensController(r.DB)
	usersController := product.NewUsersController(r.DB)

	v1 := r.Gin.Group("/v1/product")
	{
		// Comment To Threads
		v1.GET("/comments", func(ctx *gin.Context) { commentsController.GetList(ctx) })
		v1.POST("/comments", func(ctx *gin.Context) { commentsController.Post(ctx) })

		// Followers
		v1.GET("/followers", func(ctx *gin.Context) { followersController.GetList(ctx) })

		// Followings
		v1.GET("/followings", func(ctx *gin.Context) { followingsController.GetList(ctx) })

		// Me
		v1.GET("/me", func(ctx *gin.Context) { meController.Get(ctx) })
		v1.POST("/me", func(ctx *gin.Context) { meController.Post(ctx) })

		v1.PATCH("/me", func(ctx *gin.Context) { meController.Patch(ctx) })

		// Threads
		v1.GET("/threads", func(ctx *gin.Context) { threadsController.GetList(ctx) })
		v1.POST("/threads", func(ctx *gin.Context) { threadsController.Post(ctx) })

		v1.GET("/threads/:id", func(ctx *gin.Context) { threadsController.Get(ctx) })
		v1.PATCH("/threads/:id", func(ctx *gin.Context) { threadsController.Patch(ctx) })
		v1.DELETE("/threads/:id", func(ctx *gin.Context) { threadsController.Delete(ctx) })

		// Tokens
		v1.POST("/tokens", func(ctx *gin.Context) { tokensController.Post(ctx) })
		v1.POST("/tokens/refresh", func(ctx *gin.Context) { tokensController.Refresh(ctx) })

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

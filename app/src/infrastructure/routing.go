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

	articlesController := product.NewArticlesController(r.DB)
	commentsController := product.NewCommentsController(r.DB)
	favoriteArticlesController := product.NewFavoriteArticlesController(r.DB)
	followsController := product.NewFollowsController(r.DB)
	meController := product.NewMeController(r.DB)
	tokensController := product.NewTokensController(r.DB)
	usersController := product.NewUsersController(r.DB)

	v1 := r.Gin.Group("/v1/product")
	{
		// Articles
		v1.GET("/articles", func(ctx *gin.Context) { articlesController.GetList(ctx) })
		v1.POST("/articles", func(ctx *gin.Context) { articlesController.Post(ctx) })

		v1.GET("/articles/:id", func(ctx *gin.Context) { articlesController.Get(ctx) })
		v1.PATCH("/articles/:id", func(ctx *gin.Context) { articlesController.Patch(ctx) })
		v1.DELETE("/articles/:id", func(ctx *gin.Context) { articlesController.Delete(ctx) })

		// Comment To Threads
		v1.GET("/comments", func(ctx *gin.Context) { commentsController.GetList(ctx) })
		v1.POST("/comments", func(ctx *gin.Context) { commentsController.Post(ctx) })

		// Favorites
		// v1.GET("/favorites/article")
		v1.PUT("/favoriteArticles", func(ctx *gin.Context) { favoriteArticlesController.Post(ctx) })
		v1.DELETE("/favoriteArticles/:id", func(ctx *gin.Context) { favoriteArticlesController.Delete(ctx) })

		// Followers
		v1.GET("/follows", func(ctx *gin.Context) { followsController.GetList(ctx) })

		// Me
		v1.GET("/me", func(ctx *gin.Context) { meController.Get(ctx) })
		v1.POST("/me", func(ctx *gin.Context) { meController.Post(ctx) })

		v1.PATCH("/me", func(ctx *gin.Context) { meController.Patch(ctx) })

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

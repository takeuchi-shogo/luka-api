package product

import (
	"strconv"

	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/interface/controllers"
	"github.com/takeuchi-shogo/luka-api/src/interface/database"
	"github.com/takeuchi-shogo/luka-api/src/interface/gateways"
	"github.com/takeuchi-shogo/luka-api/src/usecase/product"
)

type FavoriteArticlesController struct {
	Token      product.UserTokenInteractor
	Interactor product.FavoriteArticleInteractor
}

func NewFavoriteArticlesController(db gateways.DB) *FavoriteArticlesController {
	return &FavoriteArticlesController{
		Token: product.UserTokenInteractor{
			DB:        &gateways.DBRepository{DB: db},
			User:      &database.UserRepository{},
			UserToken: &database.UserTokenRepository{},
		},
		Interactor: product.FavoriteArticleInteractor{
			DB: &gateways.DBRepository{DB: db},
		},
	}
}

func (c *FavoriteArticlesController) Post(ctx controllers.Context) {
	token, res := c.Token.Verification(ctx.PostForm("accessToken"))
	if res.Error != nil {
		ctx.JSON(res.StatusCode, controllers.NewErrorResponse(res.Error, res.Message))
		return
	}

	articleID, _ := strconv.Atoi(ctx.PostForm("articleId"))

	favorite, res := c.Interactor.Create(domain.FavoriteArticles{
		UserID:    token.UserID,
		ArticleID: articleID,
	})

	if res.Error != nil {
		return
	}

	ctx.JSON(res.StatusCode, controllers.NewH("success", favorite))
}

func (c *FavoriteArticlesController) Delete(ctx controllers.Context) {
	_, res := c.Token.Verification(ctx.PostForm("accessToken"))
	if res.Error != nil {
		ctx.JSON(res.StatusCode, controllers.NewErrorResponse(res.Error, res.Message))
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	res = c.Interactor.Delete(id)
	if res.Error != nil {
		ctx.JSON(res.StatusCode, controllers.NewErrorResponse(res.Error, res.Message))
		return
	}

	ctx.JSON(res.StatusCode, controllers.NewH("success", nil))
}

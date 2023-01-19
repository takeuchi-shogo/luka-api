package product

import (
	"strconv"

	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/interface/controllers"
	"github.com/takeuchi-shogo/luka-api/src/interface/database"
	"github.com/takeuchi-shogo/luka-api/src/interface/gateways"
	apierrors "github.com/takeuchi-shogo/luka-api/src/pkg/api-errors"
	"github.com/takeuchi-shogo/luka-api/src/usecase/product"
)

type ArticlesController struct {
	Token      product.UserTokenInteractor
	Interactor product.ArticleInteractor
}

func NewArticlesController(db gateways.DB) *ArticlesController {
	return &ArticlesController{
		Token: product.UserTokenInteractor{
			DB:        &gateways.DBRepository{DB: db},
			User:      &database.UserRepository{},
			UserToken: &database.UserTokenRepository{},
		},
		Interactor: product.ArticleInteractor{
			Article:         &database.ArticleRepository{},
			Comment:         &database.CommentRepository{},
			DB:              &gateways.DBRepository{DB: db},
			FavoriteComment: &database.FavoriteCommentRepository{},
			FavoriteArticle: &database.FavoriteArticleRepository{},
			User:            &database.UserRepository{},
		},
	}
}

func (c *ArticlesController) Get(ctx controllers.Context) {
	_, res := c.Token.Verification(ctx.Query("accessToken"))
	if res.Error != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}

	articleID, _ := strconv.Atoi(ctx.Param("id"))

	article, res := c.Interactor.Get(articleID)
	if res.Error != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}

	ctx.JSON(res.StatusCode, controllers.NewH("success", article))

}

func (c *ArticlesController) GetList(ctx controllers.Context) {

	_, res := c.Token.Verification(ctx.Query("accessToken"))
	if res.Error != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}

	articles, res := c.Interactor.GetList()
	if res.Error != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}

	ctx.JSON(res.StatusCode, controllers.NewH("success", articles))

}

func (c *ArticlesController) Post(ctx controllers.Context) {
	token, res := c.Token.Verification(ctx.PostForm("accessToken"))
	if res.Error != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}

	title := ctx.PostForm("title")
	description := ctx.PostForm("description")

	article, res := c.Interactor.Post(domain.Articles{
		UserID:      token.UserID,
		Title:       title,
		Description: description,
	})
	if res.Error != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.StatusCode, controllers.NewH("success", article))
}

func (c *ArticlesController) Patch(ctx controllers.Context) {
	token, res := c.Token.Verification(ctx.PostForm("accessToken"))
	if res.Error != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}

	articleID, articleIDErr := strconv.Atoi(ctx.Param("id"))
	if articleIDErr != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(articleIDErr.Error(), nil))
		return
	}
	title := ctx.PostForm("title")
	description := ctx.PostForm("description")

	updateArticle, res := c.Interactor.Save(domain.ArticlesForPatch{
		ID:          articleID,
		UserID:      token.UserID,
		Title:       title,
		Description: description,
	})
	if res.Error != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.StatusCode, controllers.NewH("success", updateArticle))
}

func (c *ArticlesController) Delete(ctx controllers.Context) {
	token, res := c.Token.Verification(ctx.PostForm("accessToken"))
	if res.Error != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}

	articleID, _ := strconv.Atoi(ctx.Param("id"))

	if res = c.Interactor.Delete(domain.Articles{
		ID:     articleID,
		UserID: token.UserID,
	}); res.Error != nil {
		ctx.JSON(res.StatusCode, controllers.NewErrorResponse(apierrors.BadRequest.New(res.Error, res.Message)))
		return
	}
	ctx.JSON(res.StatusCode, controllers.NewH("success", nil))
}

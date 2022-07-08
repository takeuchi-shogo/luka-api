package product

import (
	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/interface/controllers"
	"github.com/takeuchi-shogo/luka-api/src/interface/database"
	"github.com/takeuchi-shogo/luka-api/src/usecase/product"
)

type TokensController struct {
	Interactor product.UserTokenInteractor
}

func NewTokensController(db database.DB) *TokensController {
	return &TokensController{
		product.UserTokenInteractor{
			DB:        &database.DBRepository{DB: db},
			User:      &database.UserRepository{},
			UserToken: &database.UserTokenRepository{},
		},
	}
}

func (c *TokensController) Post(ctx controllers.Context) {
	userInfo := ctx.PostForm("id")
	password := ctx.PostForm("password")

	token, res := c.Interactor.Create(domain.Users{
		ScreenName: userInfo,
		Password:   password,
	})
	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}
	ctx.JSON(res.StatusCode, controllers.NewH("success", token))
}

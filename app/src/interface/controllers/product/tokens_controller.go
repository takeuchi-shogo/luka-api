package product

import (
	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/interface/controllers"
	"github.com/takeuchi-shogo/luka-api/src/interface/database"
	"github.com/takeuchi-shogo/luka-api/src/interface/gateways"
	"github.com/takeuchi-shogo/luka-api/src/usecase/product"
)

type TokensController struct {
	Interactor product.UserTokenInteractor
}

func NewTokensController(db gateways.DB) *TokensController {
	return &TokensController{
		product.UserTokenInteractor{
			DB:        &gateways.DBRepository{DB: db},
			User:      &database.UserRepository{},
			UserToken: &database.UserTokenRepository{},
		},
	}
}

func (c *TokensController) Post(ctx controllers.Context) {
	userInfo := ctx.PostForm("screenName")
	password := ctx.PostForm("password")

	token, res := c.Interactor.Create(domain.Users{
		ScreenName: userInfo,
		Password:   password,
	})
	// fmt.Println(res)
	if res.Error != nil {
		ctx.JSON(res.StatusCode, controllers.NewErrorResponse(res.Error, res.Message))
		return
	}
	ctx.JSON(res.StatusCode, controllers.NewH("success", token))
}

func (c *TokensController) Refresh(ctx controllers.Context) {
	accessToken := ctx.PostForm("accessToken")
	refreshToken := ctx.PostForm("refreshToken")

	token, res := c.Interactor.Refresh(domain.UserTokens{
		Token:        accessToken,
		RefreshToken: refreshToken,
	})
	if res.Error != nil {
		ctx.JSON(res.StatusCode, controllers.NewErrorResponse(res.Error, res.Message))
		return
	}
	ctx.JSON(res.StatusCode, controllers.NewH("success", token))
}

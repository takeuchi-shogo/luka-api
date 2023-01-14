package product

import (
	"strconv"

	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/interface/controllers"
	"github.com/takeuchi-shogo/luka-api/src/interface/database"
	"github.com/takeuchi-shogo/luka-api/src/interface/gateways"
	"github.com/takeuchi-shogo/luka-api/src/usecase/product"
)

type UsersController struct {
	Token      product.UserTokenInteractor
	Interactor product.UserInteractor
}

func NewUsersController(db gateways.DB) *UsersController {
	return &UsersController{
		Token: product.UserTokenInteractor{
			DB:        &gateways.DBRepository{DB: db},
			User:      &database.UserRepository{},
			UserToken: &database.UserTokenRepository{},
		},
		Interactor: product.UserInteractor{
			DB:   &gateways.DBRepository{DB: db},
			User: &database.UserRepository{},
		},
	}
}

func (c *UsersController) Get(ctx controllers.Context) {
	_, res := c.Token.Verification(ctx.Query("accessToken"))
	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}

	userID, _ := strconv.Atoi(ctx.Param("id"))
	user, res := c.Interactor.Get(domain.Users{
		ID: userID,
	})
	if res.ErrorMessage != nil {
		// ここでLogとしてDevErrorを流す？
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}

	ctx.JSON(200, controllers.NewH("success", user))
}

func (c *UsersController) GetList(ctx controllers.Context) {
	token, res := c.Token.Verification(ctx.Query("accessToken"))
	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}

	users, res := c.Interactor.GetList(token.UserID)
	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}

	ctx.JSON(res.StatusCode, controllers.NewH("success", users))
}

package product

import (
	"strconv"

	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/interface/controllers"
	"github.com/takeuchi-shogo/luka-api/src/usecase/product"
)

type UsersController struct {
	Interactor product.UserInteractor
}

func NewUsersController() *UsersController {
	return &UsersController{
		Interactor: product.UserInteractor{},
	}
}

func (c *UsersController) Get(ctx controllers.Context) {
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

}

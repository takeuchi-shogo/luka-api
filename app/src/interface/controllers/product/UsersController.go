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

func (controller *UsersController) Post(ctx controllers.Context) {

	displayName := ctx.PostForm("displayName")
	screenName := ctx.PostForm("screenName")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")
	age, _ := strconv.Atoi(ctx.PostForm("age"))
	gender := ctx.PostForm("gender")
	prefecture := ctx.PostForm("prefecture")

	newUser, res := controller.Interactor.Create(domain.Users{
		DisplayName: displayName,
		ScreenName:  screenName,
		Password:    password,
		Email:       email,
		Age:         age,
		Gender:      gender,
		Prefecture:  prefecture,
	})
	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}

	ctx.JSON(res.StatusCode, controllers.NewH("success", newUser))
}

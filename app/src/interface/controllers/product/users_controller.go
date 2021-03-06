package product

import (
	"strconv"

	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/interface/controllers"
	"github.com/takeuchi-shogo/luka-api/src/interface/database"
	"github.com/takeuchi-shogo/luka-api/src/usecase/product"
)

type UsersController struct {
	Token      product.UserTokenInteractor
	Interactor product.UserInteractor
}

func NewUsersController(db database.DB) *UsersController {
	return &UsersController{
		Token: product.UserTokenInteractor{
			DB:        &database.DBRepository{DB: db},
			User:      &database.UserRepository{},
			UserToken: &database.UserTokenRepository{},
		},
		Interactor: product.UserInteractor{
			DB:   &database.DBRepository{DB: db},
			User: &database.UserRepository{},
		},
	}
}

func (c *UsersController) Get(ctx controllers.Context) {
	_, res := c.Token.Authorization(ctx.Query("accessToken"))
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
	token, res := c.Token.Authorization(ctx.Query("accessToken"))
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

func (c *UsersController) Post(ctx controllers.Context) {

	displayName := ctx.PostForm("displayName")
	screenName := ctx.PostForm("screenName")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")
	age, _ := strconv.Atoi(ctx.PostForm("age"))
	gender := ctx.PostForm("gender")
	prefecture := ctx.PostForm("prefecture")

	newUser, res := c.Interactor.Create(domain.Users{
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

func (c *UsersController) Patch(ctx controllers.Context) {
	token, res := c.Token.Authorization(ctx.PostForm("accessToken"))
	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}

	updateUser := domain.UserForPatch{}

	updateUser.ID = token.UserID
	updateUser.DisplayName = ctx.PostForm("displayName")
	updateUser.ScreenName = ctx.PostForm("screenName")
	updateUser.Password = ctx.PostForm("password")
	updateUser.Email = ctx.PostForm("email")
	updateUser.Age, _ = strconv.Atoi(ctx.PostForm("age"))
	updateUser.Gender = ctx.PostForm("gender")
	updateUser.Prefecture = ctx.PostForm("prefecture")

	user, res := c.Interactor.Save(updateUser)
	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}
	ctx.JSON(res.StatusCode, controllers.NewH("success", user))
}

func (c *UsersController) Delete(ctx controllers.Context) {
	_, res := c.Token.Authorization(ctx.PostForm("accessToken"))
	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}

	screenName := ctx.Param("screenName")
	if res := c.Interactor.Delete(domain.Users{
		ScreenName: screenName,
	}); res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}
	ctx.JSON(res.StatusCode, controllers.NewH("success", nil))
}

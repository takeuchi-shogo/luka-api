package product

import (
	"fmt"
	"strconv"

	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/interface/controllers"
	"github.com/takeuchi-shogo/luka-api/src/interface/database"
	"github.com/takeuchi-shogo/luka-api/src/interface/gateways"
	"github.com/takeuchi-shogo/luka-api/src/usecase/product"
)

type MeController struct {
	Token      product.UserTokenInteractor
	Interactor product.MeInteractor
}

func NewMeController(db gateways.DB) *MeController {
	return &MeController{
		Token: product.UserTokenInteractor{
			DB:        &gateways.DBRepository{DB: db},
			User:      &database.UserRepository{},
			UserToken: &database.UserTokenRepository{},
		},
		Interactor: product.MeInteractor{
			DB:     &gateways.DBRepository{DB: db},
			Follow: &database.FollowRepository{},
			Thread: &database.ArticleRepository{},
			User:   &database.UserRepository{},
		},
	}
}

func (c *MeController) Get(ctx controllers.Context) {
	token, res := c.Token.Verification(ctx.Query("accessToken"))

	if res.Error != nil {
		ctx.JSON(res.StatusCode, controllers.NewErrorResponse(res.Error, res.Message))
		return
	}

	me, res := c.Interactor.Get(domain.Users{
		ID: token.UserID,
	})

	if res.Error != nil {
		ctx.JSON(res.StatusCode, controllers.NewErrorResponse(res.Error, res.Message))
		return
	}
	ctx.JSON(res.StatusCode, controllers.NewH("success", me))
}

func (c *MeController) Post(ctx controllers.Context) {

	displayName := ctx.PostForm("displayName")
	screenName := ctx.PostForm("screenName")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")
	age, _ := strconv.Atoi(ctx.PostForm("age"))
	gender, _ := strconv.Atoi(ctx.PostForm("gender"))
	prefecture, _ := strconv.Atoi(ctx.PostForm("prefecture"))

	newUser, res := c.Interactor.Create(domain.Users{
		DisplayName: displayName,
		ScreenName:  screenName,
		Password:    password,
		Email:       email,
		Age:         age,
		Gender:      gender,
		Prefecture:  prefecture,
	})
	if res.Error != nil {
		ctx.JSON(res.StatusCode, controllers.NewErrorResponse(res.Error, res.Message))
		return
	}

	ctx.JSON(res.StatusCode, controllers.NewH("success", newUser))
}

func (c *MeController) Patch(ctx controllers.Context) {
	token, res := c.Token.Verification(ctx.PostForm("accessToken"))
	if res.Error != nil {
		ctx.JSON(res.StatusCode, controllers.NewErrorResponse(res.Error, res.Message))
		return
	}

	updateUser := domain.UserForPatch{}

	updateUser.ID = token.UserID
	updateUser.DisplayName = ctx.PostForm("displayName")
	updateUser.ScreenName = ctx.PostForm("screenName")
	updateUser.Password = ctx.PostForm("password")
	updateUser.Email = ctx.PostForm("email")
	updateUser.Age, _ = strconv.Atoi(ctx.PostForm("age"))
	updateUser.Gender, _ = strconv.Atoi(ctx.PostForm("gender"))
	updateUser.Prefecture, _ = strconv.Atoi(ctx.PostForm("prefecture"))

	user, res := c.Interactor.Save(updateUser)
	if res.Error != nil {
		fmt.Println(res)
		ctx.JSON(res.StatusCode, controllers.NewErrorResponse(res.Error, res.Message))
		return
	}
	ctx.JSON(res.StatusCode, controllers.NewH("success", user))
}

// func (c *MeController) Delete(ctx controllers.Context) {
// 	_, res := c.Token.Authorization(ctx.PostForm("accessToken"))
// 	if res.Error != nil {
// 		ctx.JSON(res.StatusCode, controllers.NewErrorResponse(res.Error, res.Message))
// 		return
// 	}

// 	screenName := ctx.Param("screenName")
// 	if res := c.Interactor.Delete(domain.Users{
// 		ScreenName: screenName,
// 	}); res.Error != nil {
// 		ctx.JSON(res.StatusCode, controllers.NewErrorResponse(res.Error, res.Message))
// 		return
// 	}
// 	ctx.JSON(res.StatusCode, controllers.NewH("success", nil))
// }

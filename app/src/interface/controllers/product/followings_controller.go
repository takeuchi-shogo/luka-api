package product

import (
	"strconv"

	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/interface/controllers"
	"github.com/takeuchi-shogo/luka-api/src/interface/database"
	"github.com/takeuchi-shogo/luka-api/src/interface/gateways"
	"github.com/takeuchi-shogo/luka-api/src/usecase/product"
)

type FollowingsController struct {
	Token      product.UserTokenInteractor
	Interactor product.FollowingInteractor
}

func NewFollowingsController(db gateways.DB) *FollowingsController {
	return &FollowingsController{
		Token: product.UserTokenInteractor{
			DB:        &gateways.DBRepository{DB: db},
			User:      &database.UserRepository{},
			UserToken: &database.UserTokenRepository{},
		},
		Interactor: product.FollowingInteractor{},
	}
}

func (c *FollowingsController) GetList(ctx controllers.Context) {

	_, res := c.Token.Verification(ctx.Query("accessToken"))

	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}

	userID, _ := strconv.Atoi(ctx.Query("userId"))

	followings, res := c.Interactor.GetList(userID)

	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}

	ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), followings))
}

func (c *FollowingsController) Post(ctx controllers.Context) {
	token, res := c.Token.Verification(ctx.PostForm("accessToken"))

	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}

	toUserID, _ := strconv.Atoi(ctx.PostForm("userId"))

	newFollowing, res := c.Interactor.Create(domain.Followings{
		UserID:   token.UserID,
		ToUserID: toUserID,
	})

	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
	}

	ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), newFollowing))
}

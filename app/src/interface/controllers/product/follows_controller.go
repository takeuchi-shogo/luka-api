package product

import (
	"strconv"

	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/interface/controllers"
	"github.com/takeuchi-shogo/luka-api/src/interface/database"
	"github.com/takeuchi-shogo/luka-api/src/interface/gateways"
	"github.com/takeuchi-shogo/luka-api/src/usecase/product"
)

type FollowsController struct {
	Token      product.UserTokenInteractor
	Interactor product.FollowInteractor
}

func NewFollowsController(db gateways.DB) *FollowsController {
	return &FollowsController{
		Token: product.UserTokenInteractor{
			DB:        &gateways.DBRepository{DB: db},
			User:      &database.UserRepository{},
			UserToken: &database.UserTokenRepository{},
		},
		Interactor: product.FollowInteractor{
			DB:     &gateways.DBRepository{DB: db},
			Follow: &database.FollowRepository{},
		},
	}
}

func (c *FollowsController) GetList(ctx controllers.Context) {

	userID, _ := strconv.Atoi(ctx.Query("userId"))
	followers, res := c.Interactor.GetList(domain.Follows{
		UserID: userID,
	})
	if res.Error != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.StatusCode, controllers.NewH("success", followers))
}

func (c *FollowsController) Post(ctx controllers.Context) {

	token, res := c.Token.Verification(ctx.PostForm("accessToken"))
	if res.Error != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}

	toUserID, _ := strconv.Atoi(ctx.PostForm("toUserId"))

	newFollower, res := c.Interactor.Create(domain.Follows{
		UserID:   token.UserID,
		ToUserID: toUserID,
	})
	if res.Error != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.StatusCode, controllers.NewH("success", newFollower))
}

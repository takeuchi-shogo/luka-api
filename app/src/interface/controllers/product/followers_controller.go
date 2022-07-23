package product

import (
	"strconv"

	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/interface/controllers"
	"github.com/takeuchi-shogo/luka-api/src/interface/database"
	"github.com/takeuchi-shogo/luka-api/src/usecase/product"
)

type FollowersController struct {
	Token      product.UserTokenInteractor
	Interactor product.FollowerInteractor
}

func NewFollowersController(db database.DB) *FollowersController {
	return &FollowersController{
		Token: product.UserTokenInteractor{
			DB:        &database.DBRepository{DB: db},
			User:      &database.UserRepository{},
			UserToken: &database.UserTokenRepository{},
		},
		Interactor: product.FollowerInteractor{
			DB:       &database.DBRepository{DB: db},
			Follower: &database.FollowerRepository{},
		},
	}
}

func (c *FollowersController) GetList(ctx controllers.Context) {

	userID, _ := strconv.Atoi(ctx.Query("userId"))
	followers, res := c.Interactor.GetList(domain.Followers{
		UserID: userID,
	})
	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}
	ctx.JSON(res.StatusCode, controllers.NewH("success", followers))
}

func (c *FollowersController) Post(ctx controllers.Context) {

	token, res := c.Token.Authorization(ctx.PostForm("accessToken"))
	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}

	toUserID, _ := strconv.Atoi(ctx.PostForm("toUserId"))

	newFollower, res := c.Interactor.Create(domain.Followers{
		UserID:   token.UserID,
		ToUserID: toUserID,
	})
	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}
	ctx.JSON(res.StatusCode, controllers.NewH("success", newFollower))
}

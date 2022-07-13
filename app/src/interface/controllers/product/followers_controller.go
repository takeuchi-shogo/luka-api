package product

import (
	"github.com/takeuchi-shogo/luka-api/src/interface/controllers"
	"github.com/takeuchi-shogo/luka-api/src/interface/database"
	"github.com/takeuchi-shogo/luka-api/src/usecase/product"
)

type FollowersController struct {
	Interactor product.FollowerInteractor
}

func NewFollowersController(db database.DB) *FollowersController {
	return &FollowersController{
		Interactor: product.FollowerInteractor{},
	}
}

func (c *FollowersController) GetList(ctx controllers.Context) {

}

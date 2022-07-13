package product

import (
	"github.com/takeuchi-shogo/luka-api/src/interface/controllers"
	"github.com/takeuchi-shogo/luka-api/src/interface/database"
	"github.com/takeuchi-shogo/luka-api/src/usecase/product"
)

type FollowingsController struct {
	Interactor product.FollowingInteractor
}

func NewFollowingsController(db database.DB) *FollowingsController {
	return &FollowingsController{
		Interactor: product.FollowingInteractor{},
	}
}

func (c FollowingsController) GetList(ctx controllers.Context) {

}

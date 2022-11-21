package product

import (
	"github.com/takeuchi-shogo/luka-api/src/interface/gateways"
	"github.com/takeuchi-shogo/luka-api/src/usecase/product"
)

type FavoriteThreadsController struct {
	Interactor product.FavoriteThreadInteractor
}

func NewFavoriteThreadsController(db gateways.DB) *FavoriteThreadsController {
	return &FavoriteThreadsController{
		Interactor: product.FavoriteThreadInteractor{
			DB: &gateways.DBRepository{DB: db},
		},
	}
}

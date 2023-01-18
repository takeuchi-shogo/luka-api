package product

import (
	"github.com/takeuchi-shogo/luka-api/src/interface/gateways"
	"github.com/takeuchi-shogo/luka-api/src/usecase/product"
)

type FavoriteArticlesController struct {
	Interactor product.FavoriteArticleInteractor
}

func NewFavoriteArticlesController(db gateways.DB) *FavoriteArticlesController {
	return &FavoriteArticlesController{
		Interactor: product.FavoriteArticleInteractor{
			DB: &gateways.DBRepository{DB: db},
		},
	}
}

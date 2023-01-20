package product

import (
	"github.com/takeuchi-shogo/luka-api/src/interface/database"
	"github.com/takeuchi-shogo/luka-api/src/interface/gateways"
	"github.com/takeuchi-shogo/luka-api/src/usecase/product"
)

type FavoriteCommentsController struct {
	Interactor product.FavoriteCommentInteractor
}

func NewFavoriteCommentsController(db gateways.DB) *FavoriteCommentsController {
	return &FavoriteCommentsController{
		Interactor: product.FavoriteCommentInteractor{
			DB:              &gateways.DBRepository{DB: db},
			FavoriteComment: &database.FavoriteCommentRepository{},
		},
	}
}

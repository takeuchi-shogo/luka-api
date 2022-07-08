package product

import (
	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/usecase"
)

type CommentInteractor struct{}

func (interactor *CommentInteractor) Create(comment domain.Comments) (newComment domain.Comments, resultStatus *usecase.ResultStatus) {
	return newComment, usecase.NewResultStatus(200, "")
}

package product

import (
	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/usecase"
)

type CommentInteractor struct{}

func (i *CommentInteractor) GetList(comment domain.Comments) (comments []domain.Comments, resultStatus *usecase.ResultStatus) {
	return comments, usecase.NewResultStatus(200, "")
}

func (i *CommentInteractor) Create(comment domain.Comments) (newComment domain.Comments, resultStatus *usecase.ResultStatus) {
	return newComment, usecase.NewResultStatus(200, "")
}

package product

import (
	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/usecase"
)

type CommentInteractor struct {
	DB      usecase.DBRepository
	Comment usecase.CommentRepository
}

func (i *CommentInteractor) GetList(comment domain.Comments) (comments []domain.Comments, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	comments, err := i.Comment.FindByArticleID(db, comment.ArticleID)
	if err != nil {
		return []domain.Comments{}, usecase.NewResultStatus(400, domain.ErrCommentNotFound)
	}

	return comments, usecase.NewResultStatus(200, "")
}

func (i *CommentInteractor) Create(comment domain.Comments) (newComment domain.Comments, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	newComment, err := i.Comment.Create(db, comment)
	if err != nil {
		return domain.Comments{}, usecase.NewResultStatus(400, domain.ErrCreateComment)
	}

	return newComment, usecase.NewResultStatus(200, "")
}

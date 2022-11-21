package product

import (
	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/usecase"
)

type FavoriteCommentInteractor struct {
	DB              usecase.DBRepository
	FavoriteComment usecase.FavoriteCommentRepository
}

type FavoriteCommentList struct {
	Lists []domain.FavoriteComments
}

func (i *FavoriteCommentInteractor) GetList(commentID int) (favoriteCommentList FavoriteCommentList, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	favorites, err := i.FavoriteComment.FindByCommentID(db, commentID)

	if err != nil {
		return FavoriteCommentList{Lists: []domain.FavoriteComments{}}, usecase.NewResultStatus(400, domain.ErrFavoriteCommentNotFound)
	}

	return FavoriteCommentList{Lists: favorites}, usecase.NewResultStatus(200, "")
}

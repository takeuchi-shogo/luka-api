package product

import (
	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/usecase"
)

type FavoriteThreadInteractor struct {
	DB             usecase.DBRepository
	FavoriteThread usecase.FavoriteThreadRepository
}

type FavoriteThreadList struct {
	Lists []domain.FavoriteThreads
}

func (i *FavoriteThreadInteractor) GetList(threadID int) (favoriteThreadList FavoriteThreadList, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	favorites, err := i.FavoriteThread.FindByThreadID(db, threadID)

	if err != nil {
		return FavoriteThreadList{Lists: []domain.FavoriteThreads{}}, usecase.NewResultStatus(400, domain.ErrFavoriteThreadNotFound)
	}

	return FavoriteThreadList{Lists: favorites}, usecase.NewResultStatus(200, "")
}

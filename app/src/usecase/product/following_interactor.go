package product

import (
	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/usecase"
)

type FollowingInteractor struct {
	DB        usecase.DBRepository
	Following usecase.FollowingRepository
}

func (i *FollowingInteractor) GetList(userID int) (followings []domain.Followings, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	followings, err := i.Following.FindByUserID(db, userID)

	if err != nil {
		return []domain.Followings{}, usecase.NewResultStatus(404, domain.ErrFollowingNotFound)
	}

	return followings, usecase.NewResultStatus(200, "")
}

func (i *FollowingInteractor) Create(following domain.Followings) (newFollowing domain.Followings, resultStatus *usecase.ResultStatus) {
	return newFollowing, usecase.NewResultStatus(200, "")
}

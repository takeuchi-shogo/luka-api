package product

import (
	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/usecase"
)

type FollowingInteractor struct{}

func (i *FollowingInteractor) GetList(userID int) (followings []domain.Followings, resultStatus *usecase.ResultStatus) {
	return followings, usecase.NewResultStatus(200, "")
}

func (i *FollowingInteractor) Create(following domain.Followings) (newFollowing domain.Followings, resultStatus *usecase.ResultStatus) {
	return newFollowing, usecase.NewResultStatus(200, "")
}

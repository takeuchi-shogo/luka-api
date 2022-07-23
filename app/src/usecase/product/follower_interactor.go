package product

import (
	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/usecase"
)

type FollowerInteractor struct {
	DB       usecase.DBRepository
	Follower usecase.FollowerRepository
	User     usecase.UserRepository
}

func (i *FollowerInteractor) GetList(follower domain.Followers) (followers []domain.Followers, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	followers, err := i.Follower.FindByUserID(db, follower.UserID)
	if err != nil {
		return []domain.Followers{}, usecase.NewResultStatus(400, domain.ErrFollowerNotFound)
	}
	return followers, usecase.NewResultStatus(200, "")
}

func (i *FollowerInteractor) Create(follower domain.Followers) (newFollower domain.Followers, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	if _, err := i.User.FindByID(db, follower.ToUserID); err != nil {
		return domain.Followers{}, usecase.NewResultStatus(400, domain.ErrUserNotFound)
	}

	newFollower, err := i.Follower.Create(db, follower)
	if err != nil {
		return domain.Followers{}, usecase.NewResultStatus(400, domain.ErrCreateFollower)
	}

	return newFollower, usecase.NewResultStatus(200, "")
}

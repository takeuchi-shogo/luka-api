package product

import (
	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/usecase"
)

type FollowInteractor struct {
	DB     usecase.DBRepository
	Follow usecase.FollowRepository
	User   usecase.UserRepository
}

func (i *FollowInteractor) GetList(follow domain.Follows) (followers []domain.Follows, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	followers, err := i.Follow.FindByUserID(db, follow.UserID)
	if err != nil {
		return []domain.Follows{}, usecase.NewResultStatus(400, err, domain.ErrFollowerNotFound)
	}
	return followers, usecase.NewResultStatus(200, nil, "")
}

func (i *FollowInteractor) Create(follower domain.Follows) (newFollower domain.Follows, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	if _, err := i.User.FindByID(db, follower.ToUserID); err != nil {
		return domain.Follows{}, usecase.NewResultStatus(400, err, domain.ErrUserNotFound)
	}

	newFollower, err := i.Follow.Create(db, follower)
	if err != nil {
		return domain.Follows{}, usecase.NewResultStatus(400, err, domain.ErrCreateFollower)
	}

	return newFollower, usecase.NewResultStatus(200, nil, "")
}

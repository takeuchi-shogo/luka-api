package product

import (
	"errors"

	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/usecase"
)

type UserInteractor struct {
	DB   usecase.DBRepository
	User usecase.UserRepository
}

func (i *UserInteractor) Get(user domain.Users) (foundUser domain.Users, resultStatus *usecase.ResultStatus) {
	err := errors.New("テスト")
	if err != nil {
		return domain.Users{}, usecase.NewResultStatus(400, domain.ErrGetUserAccount)
	}
	return foundUser, usecase.NewResultStatus(200, "")
}

func (i *UserInteractor) GetList(userID int) (users []domain.Users, resultStatus *usecase.ResultStatus) {

	// db := i.DB.Connect()

	return users, usecase.NewResultStatus(200, "")
}

func (i *UserInteractor) Create(user domain.Users) (newUser domain.Users, resultSatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	if _, err := i.User.FindByScreenName(db, user.ScreenName); err == nil {
		return domain.Users{}, usecase.NewResultStatus(400, domain.ExistUserScreenName)
	}

	newUser, err := i.User.Create(db, user)
	if err != nil {
		return domain.Users{}, usecase.NewResultStatus(400, domain.ErrCreateUserAccount)
	}
	return newUser, usecase.NewResultStatus(200, "")
}

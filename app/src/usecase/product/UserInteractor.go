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

func (interactor *UserInteractor) Get(user domain.Users) (foundUser domain.Users, resultStatus *usecase.ResultStatus) {
	err := errors.New("テスト")
	if err != nil {
		return domain.Users{}, usecase.NewResultStatus(400, domain.GetUserAccountError)
	}
	return foundUser, usecase.NewResultStatus(200, "")
}

func (interactor *UserInteractor) Create(user domain.Users) (newUser domain.Users, resultSatus *usecase.ResultStatus) {

	db := interactor.DB.Connect()

	if _, err := interactor.User.FindByScreenName(db, user.ScreenName); err == nil {
		return domain.Users{}, usecase.NewResultStatus(400, domain.ExistUserScreenName)
	}

	newUser, err := interactor.User.Create(db, user)
	if err != nil {
		return domain.Users{}, usecase.NewResultStatus(400, domain.CreateUserAccountError)
	}
	return newUser, usecase.NewResultStatus(200, "")
}

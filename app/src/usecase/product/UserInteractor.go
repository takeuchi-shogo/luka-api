package product

import (
	"errors"

	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/usecase"
)

type UserInteractor struct{}

func (interactor *UserInteractor) Get(user domain.Users) (foundUser domain.Users, resultStatus *usecase.ResultStatus) {
	err := errors.New("テスト")
	if err != nil {
		return domain.Users{}, usecase.NewResultStatus(400, domain.GetUserAccountError)
	}
	return foundUser, usecase.NewResultStatus(200, "")
}

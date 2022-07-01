package product

import (
	"errors"

	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/usecase"
)

type UserInteractor struct{}

func (i *UserInteractor) Get(user domain.Users) (foundUser domain.Users, resultStatus *usecase.ResultStatus) {
	err := errors.New("テスト")
	if err != nil {
		return domain.Users{}, usecase.NewResultStatus(400, 2010)
	}
	return foundUser, usecase.NewResultStatus(200, 0)
}

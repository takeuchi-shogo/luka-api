package product

import (
	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/usecase"
)

type UserTokenInteractor struct {
	DB        usecase.DBRepository
	User      usecase.UserRepository
	UserToken usecase.UserTokenRepository
}

func (interactor UserTokenInteractor) Create(user domain.Users) (newToken domain.UserTokens, resultStatus *usecase.ResultStatus) {

	db := interactor.DB.Connect()

	foundUser, err := interactor.User.FindByScreenName(db, user.ScreenName)
	if err != nil {
		return domain.UserTokens{}, usecase.NewResultStatus(404, domain.SignInError)
	}

	if user.Password != foundUser.Password {
		return domain.UserTokens{}, usecase.NewResultStatus(404, domain.SignInError)
	}

	newUserToken := domain.UserTokens{}
	newUserToken.UserID = foundUser.ID

	token, err := interactor.UserToken.Create(db, newUserToken)
	if err != nil {
		return domain.UserTokens{}, usecase.NewResultStatus(404, domain.CreateUserTokenError)
	}

	return token, usecase.NewResultStatus(200, "")
}

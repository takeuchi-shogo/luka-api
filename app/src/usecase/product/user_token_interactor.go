package product

import (
	"time"

	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/usecase"
)

type UserTokenInteractor struct {
	DB        usecase.DBRepository
	User      usecase.UserRepository
	UserToken usecase.UserTokenRepository
}

func (interactor *UserTokenInteractor) Authorization(accessToken string) (token domain.UserTokens, resultStatus *usecase.ResultStatus) {

	db := interactor.DB.Connect()

	token, err := interactor.UserToken.FindByToken(db, accessToken)
	if err != nil {
		return domain.UserTokens{}, usecase.NewResultStatus(404, domain.ErrAuthorization)
	}

	if token.TokenExpiredAt < time.Now().Unix() {
		if token.RefreshTokenExpiredAt < time.Now().Unix() {
			return domain.UserTokens{}, usecase.NewResultStatus(406, domain.ErrRefreshTokenExpire)
		}
		return domain.UserTokens{}, usecase.NewResultStatus(404, domain.ErrTokenExpire)
	}

	return token, usecase.NewResultStatus(200, "")
}

func (interactor *UserTokenInteractor) Create(user domain.Users) (newToken domain.UserTokens, resultStatus *usecase.ResultStatus) {

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

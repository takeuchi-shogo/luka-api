package product

import (
	"errors"
	"net/http"
	"time"

	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/usecase"
)

type UserTokenInteractor struct {
	DB        usecase.DBRepository
	User      usecase.UserRepository
	UserToken usecase.UserTokenRepository
}

func (i *UserTokenInteractor) Verification(accessToken string) (token domain.UserTokens, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	token, err := i.UserToken.FindByToken(db, accessToken)
	if err != nil {
		return domain.UserTokens{}, usecase.NewResultStatus(406, err, domain.ErrAuthorization)
	}

	if token.TokenExpiredAt < time.Now().Unix() {
		if time.Now().Unix() < token.RefreshTokenExpiredAt {
			return domain.UserTokens{}, usecase.NewResultStatus(406, err, domain.ErrRefreshTokenExpire)
		}
		return domain.UserTokens{}, usecase.NewResultStatus(406, err, domain.ErrTokenExpire)
	}

	return token, usecase.NewResultStatus(http.StatusOK, nil, "")
}

func (i *UserTokenInteractor) Create(user domain.Users) (newToken domain.UserTokens, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	foundUser, err := i.User.FindByScreenName(db, user.ScreenName)
	if err != nil {
		return domain.UserTokens{}, usecase.NewResultStatus(401, err, domain.ErrSignIn)
	}

	if foundUser.GetPassword(user.Password) != foundUser.Password {
		return domain.UserTokens{}, usecase.NewResultStatus(401, errors.New("id and password required"), domain.ErrSignIn)
	}

	newUserToken := domain.UserTokens{}
	newUserToken.UserID = foundUser.ID

	token, err := i.UserToken.Create(db, newUserToken)
	if err != nil {
		return domain.UserTokens{}, usecase.NewResultStatus(401, err, domain.ErrCreateUserToken)
	}

	return token, usecase.NewResultStatus(200, nil, "")
}

func (i *UserTokenInteractor) Refresh(userToken domain.UserTokens) (newUserToken domain.UserTokens, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	foundUserToken, err := i.UserToken.FindByToken(db, userToken.Token)
	if err != nil {
		return domain.UserTokens{}, usecase.NewResultStatus(400, err, "見つかりません")
	}
	if foundUserToken.RefreshToken != userToken.RefreshToken {
		return domain.UserTokens{}, usecase.NewResultStatus(400, err, "トークンが一致しません")
	}
	if foundUserToken.RefreshTokenExpiredAt < time.Now().Unix() {
		return domain.UserTokens{}, usecase.NewResultStatus(400, err, "有効期限切れです")
	}

	newUserToken = domain.UserTokens{}
	newUserToken.UserID = foundUserToken.UserID

	res, err := i.UserToken.Create(db, newUserToken)
	if err != nil {
		return domain.UserTokens{}, usecase.NewResultStatus(400, err, "作成失敗")
	}

	return res, usecase.NewResultStatus(200, nil, "")
}

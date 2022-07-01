package domain

import "errors"

type Error struct {
	DeveloperError error
	UserError      error
}

type DeveloperError struct {
	// 各種開発者向けエラーリスト
	GetUserAccountErr string
}

type UserError struct {
	// 各種ユーザー向けエラーリスト
	GetUserAccountErr string
}

func NewError(code int) *Error {

	errorMessage := new(Error)

	developerError := NewDeveloperError()
	userError := NewUserError()

	switch code {
	case 2010:
		errorMessage.DeveloperError = errors.New(developerError.GetUserAccountErr)
		errorMessage.UserError = errors.New(userError.GetUserAccountErr)
	default:

	}

	return errorMessage
}

func NewDeveloperError() *DeveloperError {
	return &DeveloperError{
		GetUserAccountErr: "user is not found",
	}
}

func NewUserError() *UserError {
	return &UserError{
		GetUserAccountErr: "ユーザー情報の取得に失敗しました。",
	}
}

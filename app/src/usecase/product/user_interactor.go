package product

import (
	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/usecase"
)

type UserInteractor struct {
	DB   usecase.DBRepository
	User usecase.UserRepository
}

func (i *UserInteractor) Get(user domain.Users) (foundUser domain.Users, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	user, err := i.User.FindByID(db, user.ID)
	if err != nil {
		return domain.Users{}, usecase.NewResultStatus(400, err, domain.ErrGetUserAccount)
	}
	return user, usecase.NewResultStatus(200, nil, "")
}

func (i *UserInteractor) GetList(userID int) (users []domain.Users, resultStatus *usecase.ResultStatus) {

	// db := i.DB.Connect()

	return users, usecase.NewResultStatus(200, nil, "")
}

func (i *UserInteractor) Create(user domain.Users) (newUser domain.Users, resultSatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	if _, err := i.User.FindByScreenName(db, user.ScreenName); err == nil {
		return domain.Users{}, usecase.NewResultStatus(400, err, domain.ExistUserScreenName)
	}

	newUser, err := i.User.Create(db, user)
	if err != nil {
		return domain.Users{}, usecase.NewResultStatus(400, err, domain.ErrCreateUserAccount)
	}
	return newUser, usecase.NewResultStatus(200, nil, "")
}

func (i *UserInteractor) Save(user domain.UserForPatch) (updateUser domain.Users, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	foundUser, err := i.User.FindByID(db, user.ID)
	if err != nil {
		return domain.Users{}, usecase.NewResultStatus(400, err, domain.ErrUserNotFound)
	}

	foundUser.DisplayName = user.DisplayName
	foundUser.ScreenName = user.ScreenName
	foundUser.Email = user.Email
	foundUser.Age = user.Age
	foundUser.Gender = user.Gender
	foundUser.Prefecture = user.Prefecture

	foundUser.Password = foundUser.GetPassword(user.Password)

	updateUser, err = i.User.Save(db, foundUser)
	if err != nil {
		return domain.Users{}, usecase.NewResultStatus(404, err, domain.ErrUpdateUserAccount)
	}

	return updateUser, usecase.NewResultStatus(200, nil, "")
}

func (i *UserInteractor) Delete(user domain.Users) *usecase.ResultStatus {

	db := i.DB.Connect()

	foundUser, err := i.User.FindByScreenName(db, user.ScreenName)
	if err != nil {
		return usecase.NewResultStatus(400, err, domain.ErrUserNotFound)
	}

	if err = i.User.Delete(db, foundUser); err != nil {
		return usecase.NewResultStatus(404, err, domain.ErrDeleteUserAccount)
	}
	return usecase.NewResultStatus(200, nil, "")
}

package product

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/usecase"
)

type MeInteractor struct {
	DB     usecase.DBRepository
	Follow usecase.FollowRepository
	Thread usecase.ThreadRepository
	User   usecase.UserRepository
}

func (i *MeInteractor) Get(user domain.Users) (me domain.UsersForGet, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	foundMe, err := i.User.FindByID(db, user.ID)

	if err != nil {
		return domain.UsersForGet{}, usecase.NewResultStatus(400, domain.ErrUserNotFound)
	}

	builtMe, err := i.build(db, foundMe)
	if err != nil {
		return domain.UsersForGet{}, usecase.NewResultStatus(400, err.Error())
	}
	fmt.Println(builtMe)

	return builtMe, usecase.NewResultStatus(200, "")
}

func (i *MeInteractor) Create(user domain.Users) (newUser domain.Users, resultSatus *usecase.ResultStatus) {

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

func (i MeInteractor) Save(user domain.UserForPatch) (updatedMe domain.Users, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	foundMe, err := i.User.FindByID(db, user.ID)

	if err != nil {
		return domain.Users{}, usecase.NewResultStatus(400, domain.ErrUserNotFound)
	}

	foundMe.ScreenName = user.ScreenName
	foundMe.DisplayName = user.DisplayName
	foundMe.Password = foundMe.GetPassword(user.Password)
	foundMe.Email = user.Email
	foundMe.Age = user.Age
	foundMe.Gender = user.Gender
	foundMe.Prefecture = user.Prefecture

	updatedMe, err = i.User.Save(db, foundMe)

	if err != nil {
		return domain.Users{}, usecase.NewResultStatus(400, domain.ErrUserNotFound)
	}

	return updatedMe, usecase.NewResultStatus(200, "")
}

func (i *MeInteractor) build(db *gorm.DB, me domain.Users) (builtMe domain.UsersForGet, err error) {
	threadCnt, _ := i.Thread.CountByUserID(db, me.ID)

	followingCnt, _ := i.Follow.CountByUserID(db, me.ID)

	followerCnt, _ := i.Follow.CountByToUserID(db, me.ID)

	builtMe = me.BuildForGet()

	builtMe.ThreadCnt = threadCnt
	builtMe.FollowingsCnt = followingCnt
	builtMe.FollowersCnt = followerCnt

	return builtMe, nil
}

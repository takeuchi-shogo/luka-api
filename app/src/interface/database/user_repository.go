package database

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type UserRepository struct{}

func (r *UserRepository) FindByID(db *gorm.DB, id int) (user domain.Users, err error) {
	user = domain.Users{}
	db.First(id, &user)
	if user.ID <= 0 {
		return domain.Users{}, errors.New("user is not found")
	}
	return user, nil
}

func (r *UserRepository) FindByScreenName(db *gorm.DB, screenName string) (user domain.Users, err error) {
	user = domain.Users{}
	db.Where("screen_name = ?", screenName).First(&screenName)
	if user.ID <= 0 {
		return domain.Users{}, errors.New("user is not found")
	}
	return user, nil
}

func (r *UserRepository) Create(db *gorm.DB, user domain.Users) (newUser domain.Users, err error) {

	newUser = domain.Users{}

	newUser.DisplayName = user.DisplayName
	newUser.ScreenName = user.ScreenName
	newUser.Password = user.GetPassword(user.Password)
	newUser.Email = user.Email
	newUser.Age = user.Age
	newUser.Gender = user.Gender
	newUser.Prefecture = user.Prefecture

	currentTime := time.Now().Unix()
	newUser.CreatedAt = currentTime
	newUser.UpdatedAt = currentTime
	newUser.DeletedAt = nil

	db.NewRecord(&newUser)
	err = db.Create(&newUser).Error

	return newUser, err
}

func (r *UserRepository) Save(db *gorm.DB, user domain.Users) (updateUser domain.Users, err error) {

	err = db.Save(&user).Error

	return user, nil
}

func (r *UserRepository) Delete(db *gorm.DB, user domain.Users) error {
	return db.Delete(&user).Error
}

package main

import (
	"fmt"
	"time"

	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/infrastructure"
	"github.com/takeuchi-shogo/luka-api/src/interface/database"
	"github.com/takeuchi-shogo/luka-api/src/interface/gateways"
	"github.com/takeuchi-shogo/luka-api/src/usecase/product"
)

func main() {

	s := float64(time.Now().UnixNano())

	config := infrastructure.NewConfig()
	db := infrastructure.NewDB(config)

	interactor := &product.UserInteractor{
		DB:   &gateways.DBRepository{DB: db},
		User: &database.UserRepository{},
	}

	users := createUsers()

	for _, user := range users {
		_, res := interactor.Create(user)

		if res.ErrorMessage != nil {
			fmt.Println(res.ErrorMessage.Error())
			break
		}
	}

	f := float64(time.Now().UnixNano())
	fmt.Printf("%.3fsec\n", (s-f)/1000000000)
}

func createUsers() []domain.Users {

	newUsers := []domain.Users{}

	// for i := 0; i < 2; i++ {
	newUser := domain.Users{}

	newUser.DisplayName = "男性テスター"
	newUser.ScreenName = "tester"
	newUser.Password = "okokok"
	newUser.Email = "tester@example.com"
	newUser.Age = 20
	newUser.Gender = 1
	newUser.Prefecture = 3

	newUsers = append(newUsers, newUser)
	// }

	// for i := 0; i < 2; i++ {
	newUser = domain.Users{}

	newUser.DisplayName = "女性テスター"
	newUser.ScreenName = "woman_tester"
	newUser.Password = "okokok"
	newUser.Email = "woman_tester@example.com"
	newUser.Age = 20
	newUser.Gender = 2
	newUser.Prefecture = 4

	newUsers = append(newUsers, newUser)
	// }

	return newUsers
}

package product

import (
	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/usecase"
)

type ThreadInteractor struct {
	DB     usecase.DBRepository
	Thread usecase.ThreadRepository
	User   usecase.UserRepository
}

func (interactor *ThreadInteractor) Post(thread domain.Threads) (newThead domain.Threads, resultStatus *usecase.ResultStatus) {

	db := interactor.DB.Connect()

	user, err := interactor.User.FindByID(db, thread.UserID)
	if err != nil {
		return domain.Threads{}, usecase.NewResultStatus(404, domain.ErrUserNotFound)
	}

	// タイトル、内容共に禁止用語などあればここでチェックする
	//
	//

	newThead, err = interactor.Thread.Create(db, domain.Threads{
		UserID:      user.ID,
		Title:       thread.Title,
		Description: thread.Description,
	})
	if err != nil {
		return domain.Threads{}, usecase.NewResultStatus(400, domain.ErrCreateThread)
	}

	return newThead, usecase.NewResultStatus(200, "")
}

package product

import (
	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/usecase"
)

type ThreadInteractor struct {
	Comment usecase.CommentRepository
	DB      usecase.DBRepository
	Thread  usecase.ThreadRepository
	User    usecase.UserRepository
}

func (i *ThreadInteractor) Get(threadID int) (thread domain.Threads, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	thread, err := i.Thread.FindByID(db, threadID)
	if err != nil {
		return domain.Threads{}, usecase.NewResultStatus(400, domain.ErrThreadNotFound)
	}

	user, err := i.User.FindByID(db, thread.UserID)
	if err != nil {
		return domain.Threads{}, usecase.NewResultStatus(400, domain.ErrUserNotFound)
	}

	comments, err := i.Comment.FindByThreadID(db, thread.ID)
	if err != nil {
		return domain.Threads{}, usecase.NewResultStatus(400, domain.ErrCommentNotFound)
	}

	buildThread := thread.BuildForGet()

	buildThread.User = user
	buildThread.Comments = comments

	return thread, usecase.NewResultStatus(200, "")
}

func (i *ThreadInteractor) Post(thread domain.Threads) (newThead domain.Threads, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	user, err := i.User.FindByID(db, thread.UserID)
	if err != nil {
		return domain.Threads{}, usecase.NewResultStatus(404, domain.ErrUserNotFound)
	}

	// タイトル、内容共に禁止用語などあればここでチェックする
	//
	//

	newThead, err = i.Thread.Create(db, domain.Threads{
		UserID:      user.ID,
		Title:       thread.Title,
		Description: thread.Description,
	})
	if err != nil {
		return domain.Threads{}, usecase.NewResultStatus(400, domain.ErrCreateThread)
	}

	return newThead, usecase.NewResultStatus(200, "")
}

func (i *ThreadInteractor) Save(thread domain.ThreadsForPatch) (updateThread domain.Threads, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	user, err := i.User.FindByID(db, thread.UserID)
	if err != nil {
		return domain.Threads{}, usecase.NewResultStatus(400, domain.ErrUserNotFound)
	}

	foundThread, err := i.Thread.FindByID(db, thread.ID)
	if err != nil {
		return domain.Threads{}, usecase.NewResultStatus(400, domain.ErrThreadNotFound)
	}

	foundThread.ID = thread.ID
	foundThread.UserID = user.ID
	foundThread.Title = thread.Title
	foundThread.Description = thread.Description

	updateThread, err = i.Thread.Save(db, foundThread)
	if err != nil {
		return domain.Threads{}, usecase.NewResultStatus(400, domain.ErrSaveThread)
	}

	return updateThread, usecase.NewResultStatus(200, "")
}

func (i *ThreadInteractor) Delete(thread domain.Threads) (resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	if _, err := i.Thread.FindByID(db, thread.ID); err != nil {
		return usecase.NewResultStatus(400, domain.ErrDeleteThread)
	}
	return usecase.NewResultStatus(200, "")
}

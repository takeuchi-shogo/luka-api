package product

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/usecase"
)

type ThreadInteractor struct {
	Comment         usecase.CommentRepository
	DB              usecase.DBRepository
	FavoriteComment usecase.FavoriteCommentRepository
	FavoriteThread  usecase.FavoriteThreadRepository
	Thread          usecase.ThreadRepository
	User            usecase.UserRepository
}

type ThreadList struct {
	Lists []domain.ThreadsForGet `json:"lists"`
}

func (i *ThreadInteractor) Get(threadID int) (thread domain.ThreadsForGet, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	res, err := i.Thread.FindByID(db, threadID)
	if err != nil {
		return domain.ThreadsForGet{}, usecase.NewResultStatus(400, domain.ErrThreadNotFound)
	}

	buildThread, errorMessage := i.build(db, res)
	if errorMessage != "" {
		return domain.ThreadsForGet{}, usecase.NewResultStatus(400, errorMessage)
	}

	return buildThread, usecase.NewResultStatus(200, "")
}

func (i *ThreadInteractor) GetList() (threadList ThreadList, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	threads, err := i.Thread.Find(db)

	if err != nil {
		return ThreadList{Lists: []domain.ThreadsForGet{}}, usecase.NewResultStatus(400, domain.ErrThreadNotFound)
	}

	buildThreads := []domain.ThreadsForGet{}

	for _, thread := range threads {
		buildThread, err := i.build(db, thread)
		if err != "" {
			continue
		}
		buildThreads = append(buildThreads, buildThread)
	}

	return ThreadList{Lists: buildThreads}, usecase.NewResultStatus(200, "")
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

	thread, err := i.Thread.FindByID(db, thread.ID)
	if err != nil {
		return usecase.NewResultStatus(400, domain.ErrDeleteThread)
	}
	fmt.Println("thread", thread)

	if err := i.Thread.Delete(db, thread); err != nil {
		return usecase.NewResultStatus(400, domain.ErrDeleteThread)
	}
	return usecase.NewResultStatus(200, "")
}

func (i *ThreadInteractor) build(db *gorm.DB, thread domain.Threads) (buildThread domain.ThreadsForGet, errorMessage string) {

	user, err := i.User.FindByID(db, thread.UserID)
	if err != nil {
		return domain.ThreadsForGet{}, domain.ErrGetUserAccount
	}

	comments, _ := i.Comment.FindByThreadID(db, thread.ID)

	favoriteThreads, _ := i.FavoriteThread.FindByThreadID(db, thread.ID)

	buildThread = thread.BuildForGet()

	buildThread.User = user
	buildThread.Comments = comments

	buildThread.CommentCnt = len(comments)
	buildThread.FavoriteCnt = len(favoriteThreads)

	return buildThread, ""
}

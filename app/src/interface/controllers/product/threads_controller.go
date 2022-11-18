package product

import (
	"strconv"

	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/interface/controllers"
	"github.com/takeuchi-shogo/luka-api/src/interface/database"
	"github.com/takeuchi-shogo/luka-api/src/interface/gateways"
	"github.com/takeuchi-shogo/luka-api/src/usecase/product"
)

type ThreadsController struct {
	Token      product.UserTokenInteractor
	Interactor product.ThreadInteractor
}

func NewThreadsController(db gateways.DB) *ThreadsController {
	return &ThreadsController{
		Token: product.UserTokenInteractor{
			DB:        &gateways.DBRepository{DB: db},
			User:      &database.UserRepository{},
			UserToken: &database.UserTokenRepository{},
		},
		Interactor: product.ThreadInteractor{
			Comment: &database.CommentRepository{},
			DB:      &gateways.DBRepository{DB: db},
			Thread:  &database.ThreadRepository{},
			User:    &database.UserRepository{},
		},
	}
}

func (c *ThreadsController) Get(ctx controllers.Context) {
	_, res := c.Token.Authorization(ctx.Query("accessToken"))
	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}

	threadID, _ := strconv.Atoi(ctx.Param("id"))

	thread, res := c.Interactor.Get(threadID)
	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}

	ctx.JSON(res.StatusCode, controllers.NewH("success", thread))

}

func (c *ThreadsController) GetList(ctx controllers.Context) {

	_, res := c.Token.Authorization(ctx.Query("accessToken"))
	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}

	threads, res := c.Interactor.GetList()
	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}

	ctx.JSON(res.StatusCode, controllers.NewH("success", threads))

}

func (c *ThreadsController) Post(ctx controllers.Context) {
	token, res := c.Token.Authorization(ctx.PostForm("accessToken"))
	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}

	title := ctx.PostForm("title")
	description := ctx.PostForm("description")

	thread, res := c.Interactor.Post(domain.Threads{
		UserID:      token.UserID,
		Title:       title,
		Description: description,
	})
	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}
	ctx.JSON(res.StatusCode, controllers.NewH("success", thread))
}

func (c *ThreadsController) Patch(ctx controllers.Context) {
	token, res := c.Token.Authorization(ctx.PostForm("accessToken"))
	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}

	threadID, _ := strconv.Atoi(ctx.PostForm("threadId"))
	title := ctx.PostForm("title")
	description := ctx.PostForm("description")

	updateThread, res := c.Interactor.Save(domain.ThreadsForPatch{
		ID:          threadID,
		UserID:      token.UserID,
		Title:       title,
		Description: description,
	})
	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}
	ctx.JSON(res.StatusCode, controllers.NewH("success", updateThread))
}

func (c *ThreadsController) Delete(ctx controllers.Context) {
	token, res := c.Token.Authorization(ctx.Query("accessToken"))
	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}

	threadID, _ := strconv.Atoi(ctx.PostForm("threadId"))

	if res = c.Interactor.Delete(domain.Threads{
		ID:     threadID,
		UserID: token.UserID,
	}); res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}
	ctx.JSON(res.StatusCode, controllers.NewH("success", nil))
}

package product

import (
	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/interface/controllers"
	"github.com/takeuchi-shogo/luka-api/src/interface/database"
	"github.com/takeuchi-shogo/luka-api/src/usecase/product"
)

type ThreadsController struct {
	Token      product.UserTokenInteractor
	Interactor product.ThreadInteractor
}

func NewThreadsController(db database.DB) *ThreadsController {
	return &ThreadsController{
		Token: product.UserTokenInteractor{
			DB:        &database.DBRepository{DB: db},
			User:      &database.UserRepository{},
			UserToken: &database.UserTokenRepository{},
		},
		Interactor: product.ThreadInteractor{},
	}
}

func (c *ThreadsController) GetList(ctx controllers.Context) {

}

func (c *ThreadsController) Post(ctx controllers.Context) {
	token, res := c.Token.Authorization(ctx.Query("accessToken"))
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

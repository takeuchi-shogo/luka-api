package product

import (
	"strconv"

	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/interface/controllers"
	"github.com/takeuchi-shogo/luka-api/src/interface/database"
	"github.com/takeuchi-shogo/luka-api/src/usecase/product"
)

type CommentsController struct {
	Interactor product.CommentInteractor
}

func NewCommentsController(db database.DB) *CommentsController {
	return &CommentsController{
		Interactor: product.CommentInteractor{},
	}
}

func (c *CommentsController) Post(ctx controllers.Context) {

	threadID, _ := strconv.Atoi(ctx.PostForm("threadId"))
	userID, _ := strconv.Atoi(ctx.PostForm("userId"))
	content := ctx.PostForm("content")

	newComment, res := c.Interactor.Create(domain.Comments{
		ThreadID: threadID,
		UserID:   userID,
		Content:  content,
	})
	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}
	ctx.JSON(res.StatusCode, controllers.NewH("success", newComment))
}

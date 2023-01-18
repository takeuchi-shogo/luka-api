package product

import (
	"strconv"

	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/interface/controllers"
	"github.com/takeuchi-shogo/luka-api/src/interface/gateways"
	"github.com/takeuchi-shogo/luka-api/src/usecase/product"
)

type CommentsController struct {
	Interactor product.CommentInteractor
}

func NewCommentsController(db gateways.DB) *CommentsController {
	return &CommentsController{
		Interactor: product.CommentInteractor{},
	}
}

func (c *CommentsController) GetList(ctx controllers.Context) {

	articleID, _ := strconv.Atoi(ctx.Query("articleId"))

	comments, res := c.Interactor.GetList(domain.Comments{
		ArticleID: articleID,
	})
	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}
	ctx.JSON(res.StatusCode, controllers.NewH("success", comments))
}

func (c *CommentsController) Post(ctx controllers.Context) {

	articleID, _ := strconv.Atoi(ctx.PostForm("articleId"))
	userID, _ := strconv.Atoi(ctx.PostForm("userId"))
	content := ctx.PostForm("content")

	newComment, res := c.Interactor.Create(domain.Comments{
		ArticleID: articleID,
		UserID:    userID,
		Content:   content,
	})
	if res.ErrorMessage != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.ErrorMessage.Error(), nil))
		return
	}
	ctx.JSON(res.StatusCode, controllers.NewH("success", newComment))
}

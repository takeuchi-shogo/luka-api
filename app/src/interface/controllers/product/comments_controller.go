package product

import (
	"strconv"

	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/interface/controllers"
	"github.com/takeuchi-shogo/luka-api/src/interface/database"
	"github.com/takeuchi-shogo/luka-api/src/interface/gateways"
	"github.com/takeuchi-shogo/luka-api/src/usecase/product"
)

type CommentsController struct {
	Token      product.UserTokenInteractor
	Interactor product.CommentInteractor
}

func NewCommentsController(db gateways.DB) *CommentsController {
	return &CommentsController{
		Token: product.UserTokenInteractor{
			DB:        &gateways.DBRepository{DB: db},
			User:      &database.UserRepository{},
			UserToken: &database.UserTokenRepository{},
		},
		Interactor: product.CommentInteractor{},
	}
}

func (c *CommentsController) GetList(ctx controllers.Context) {

	articleID, _ := strconv.Atoi(ctx.Query("articleId"))

	comments, res := c.Interactor.GetList(domain.Comments{
		ArticleID: articleID,
	})
	if res.Error != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.StatusCode, controllers.NewH("success", comments))
}

func (c *CommentsController) Post(ctx controllers.Context) {

	token, res := c.Token.Verification(ctx.PostForm("accessToken"))

	if res.Error != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}

	articleID, _ := strconv.Atoi(ctx.PostForm("articleId"))
	toUserID, _ := strconv.Atoi(ctx.PostForm("toUserId"))
	content := ctx.PostForm("content")

	newComment, res := c.Interactor.Create(domain.Comments{
		ArticleID: articleID,
		UserID:    token.UserID,
		ToUserID:  toUserID,
		Content:   content,
	})
	if res.Error != nil {
		ctx.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.StatusCode, controllers.NewH("success", newComment))
}

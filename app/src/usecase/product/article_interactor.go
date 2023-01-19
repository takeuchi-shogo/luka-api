package product

import (
	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/usecase"
)

type ArticleInteractor struct {
	Comment         usecase.CommentRepository
	DB              usecase.DBRepository
	FavoriteComment usecase.FavoriteCommentRepository
	FavoriteArticle usecase.FavoriteArticleRepository
	Article         usecase.ArticleRepository
	User            usecase.UserRepository
}

type ArticleList struct {
	Lists []domain.ArticlesForGet `json:"lists"`
}

func (i *ArticleInteractor) Get(articleID int) (Article domain.ArticlesForGet, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	res, err := i.Article.FindByID(db, articleID)
	if err != nil {
		return domain.ArticlesForGet{}, usecase.NewResultStatus(400, err, domain.ErrArticleNotFound)
	}

	buildArticle, err := i.build(db, res)
	if err != nil {
		return domain.ArticlesForGet{}, usecase.NewResultStatus(400, err, domain.ErrGetUserAccount)
	}

	return buildArticle, usecase.NewResultStatus(200, nil, "")
}

func (i *ArticleInteractor) GetList() (articleList ArticleList, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	articles, err := i.Article.Find(db)

	if err != nil {
		return ArticleList{Lists: []domain.ArticlesForGet{}}, usecase.NewResultStatus(400, err, domain.ErrArticleNotFound)
	}

	buildArticles := []domain.ArticlesForGet{}

	for _, Article := range articles {
		buildArticle, err := i.build(db, Article)
		if err != nil {
			continue
		}
		buildArticles = append(buildArticles, buildArticle)
	}

	return ArticleList{Lists: buildArticles}, usecase.NewResultStatus(200, err, "")
}

func (i *ArticleInteractor) Post(article domain.Articles) (newArticle domain.Articles, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	user, err := i.User.FindByID(db, article.UserID)
	if err != nil {
		return domain.Articles{}, usecase.NewResultStatus(404, err, domain.ErrUserNotFound)
	}

	// タイトル、内容共に禁止用語などあればここでチェックする
	//
	//

	newArticle, err = i.Article.Create(db, domain.Articles{
		UserID:      user.ID,
		Title:       article.Title,
		Description: article.Description,
	})
	if err != nil {
		return domain.Articles{}, usecase.NewResultStatus(400, err, domain.ErrCreateArticle)
	}

	return newArticle, usecase.NewResultStatus(200, nil, "")
}

func (i *ArticleInteractor) Save(article domain.ArticlesForPatch) (updateArticle domain.Articles, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	user, err := i.User.FindByID(db, article.UserID)
	if err != nil {
		return domain.Articles{}, usecase.NewResultStatus(400, err, domain.ErrUserNotFound)
	}

	foundArticle, err := i.Article.FindByID(db, article.ID)
	if err != nil {
		return domain.Articles{}, usecase.NewResultStatus(400, err, domain.ErrArticleNotFound)
	}

	foundArticle.ID = article.ID
	foundArticle.UserID = user.ID
	foundArticle.Title = article.Title
	foundArticle.Description = article.Description

	updateArticle, err = i.Article.Save(db, foundArticle)
	if err != nil {
		return domain.Articles{}, usecase.NewResultStatus(400, err, domain.ErrSaveArticle)
	}

	return updateArticle, usecase.NewResultStatus(200, nil, "")
}

func (i *ArticleInteractor) Delete(article domain.Articles) (resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	foundArticle, err := i.Article.FindByID(db, article.ID)
	if err != nil {
		return usecase.NewResultStatus(400, err, domain.ErrDeleteArticle)
	}

	if err := i.Article.Delete(db, foundArticle); err != nil {
		return usecase.NewResultStatus(400, err, domain.ErrDeleteArticle)
	}
	return usecase.NewResultStatus(200, nil, "")
}

func (i *ArticleInteractor) build(db *gorm.DB, article domain.Articles) (buildArticle domain.ArticlesForGet, err error) {

	user, err := i.User.FindByID(db, article.UserID)
	if err != nil {
		return domain.ArticlesForGet{}, err
	}

	comments, _ := i.Comment.FindByArticleID(db, article.ID)

	favoriteArticles, _ := i.FavoriteArticle.FindByArticleID(db, article.ID)

	buildArticle = article.BuildForGet()

	buildArticle.User = user
	buildArticle.Comments = comments

	buildArticle.CommentCnt = len(comments)
	buildArticle.FavoriteCnt = len(favoriteArticles)

	return buildArticle, nil
}

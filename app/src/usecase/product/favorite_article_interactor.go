package product

import (
	"github.com/takeuchi-shogo/luka-api/src/domain"
	"github.com/takeuchi-shogo/luka-api/src/usecase"
)

type FavoriteArticleInteractor struct {
	Article         usecase.ArticleRepository
	DB              usecase.DBRepository
	FavoriteArticle usecase.FavoriteArticleRepository
}

type FavoriteArticleList struct {
	Lists []domain.FavoriteArticles
}

func (i *FavoriteArticleInteractor) GetList(ArticleID int) (favoriteArticleList FavoriteArticleList, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	favorites, err := i.FavoriteArticle.FindByArticleID(db, ArticleID)

	if err != nil {
		return FavoriteArticleList{Lists: []domain.FavoriteArticles{}}, usecase.NewResultStatus(400, domain.ErrFavoriteArticleNotFound)
	}

	return FavoriteArticleList{Lists: favorites}, usecase.NewResultStatus(200, "")
}

func (i *FavoriteArticleInteractor) Create(favorite domain.FavoriteArticles) (newFavorite domain.FavoriteArticles, resultStatus *usecase.ResultStatus) {

	db := i.DB.Connect()

	if _, err := i.Article.FindByID(db, favorite.ArticleID); err != nil {
		return domain.FavoriteArticles{}, usecase.NewResultStatus(400, domain.ErrFavoriteArticleCreate)
	}

	newFavorite, err := i.FavoriteArticle.Create(db, favorite)

	if err != nil {
		return domain.FavoriteArticles{}, usecase.NewResultStatus(400, domain.ErrFavoriteArticleCreate)
	}
	return newFavorite, usecase.NewResultStatus(200, "")
}

func (i *FavoriteArticleInteractor) Delete(id int) *usecase.ResultStatus {

	db := i.DB.Connect()

	foundFavorite, err := i.FavoriteArticle.FindByID(db, id)

	if err != nil {
		return usecase.NewResultStatus(404, domain.ErrFavoriteArticleNotFound)
	}

	if err = i.FavoriteArticle.Delete(db, foundFavorite.ID); err != nil {
		return usecase.NewResultStatus(400, domain.ErrDeleteFavoriteArticle)
	}

	return usecase.NewResultStatus(200, "")
}

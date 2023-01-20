package database

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/luka-api/src/domain"
)

type FavoriteArticleRepository struct{}

func (r *FavoriteArticleRepository) FindByID(db *gorm.DB, id int) (favorite domain.FavoriteArticles, err error) {

	favorite = domain.FavoriteArticles{}

	db.First(&favorite, id)

	if favorite.ID == 0 {
		return domain.FavoriteArticles{}, err
	}

	return favorite, nil
}

func (r *FavoriteArticleRepository) FindByUserID(db *gorm.DB, userID int) (favorites []domain.FavoriteArticles, err error) {

	favorites = []domain.FavoriteArticles{}

	db.Where("user_id = ?", userID).Find(&favorites)

	if len(favorites) == 0 {
		return []domain.FavoriteArticles{}, err
	}

	return favorites, nil
}

func (r FavoriteArticleRepository) FindByArticleID(db *gorm.DB, articleID int) (favorites []domain.FavoriteArticles, err error) {

	favorites = []domain.FavoriteArticles{}

	db.Where("article_id = ?", articleID).Find(&favorites)

	if len(favorites) == 0 {
		return []domain.FavoriteArticles{}, err
	}

	return favorites, nil
}

func (r *FavoriteArticleRepository) Create(db *gorm.DB, favorite domain.FavoriteArticles) (newFavorite domain.FavoriteArticles, err error) {

	newFavorite = domain.FavoriteArticles{}

	newFavorite.UserID = favorite.UserID
	newFavorite.ArticleID = favorite.ArticleID

	currentTime := time.Now().Unix()
	newFavorite.CreatedAt = currentTime
	// newFavorite.UpdatedAt = currentTime

	db.NewRecord((&newFavorite))
	err = db.Create(&newFavorite).Error
	return newFavorite, nil
}

func (r *FavoriteArticleRepository) Delete(db *gorm.DB, id int) error {

	favorite, err := r.FindByID(db, id)

	if err != nil {
		return err
	}

	err = db.Delete(&favorite).Error

	return err
}

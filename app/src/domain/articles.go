package domain

import "errors"

type Articles struct {
	ID          int    `json:"id"`
	UserID      int    `json:"userId"`
	Title       string `json:"title"`
	Description string `json:"description"`

	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	DeletedAt *int64 `json:"deletedAt"`
}

type ArticlesForGet struct {
	ID          int    `json:"id"`
	UserID      int    `json:"userId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"createdAt"`

	User     Users      `json:"user"`
	Comments []Comments `json:"comments"`

	CommentCnt  int `json:"commentCnt"`
	FavoriteCnt int `json:"favoriteCnt"`
}

type ArticlesForPatch struct {
	ID          int
	UserID      int
	Title       string
	Description string
}

func (a *Articles) Validate() error {
	if err := a.validateTitle(); err != nil {
		return err
	}
	if err := a.validateDescription(); err != nil {
		return err
	}

	return nil
}

func (a *Articles) validateTitle() error {
	if a.Title == "" {
		return errors.New("title is required")
	}
	return nil
}

func (a *Articles) validateDescription() error {
	if a.Description == "" {
		return errors.New("description is required")
	}
	return nil
}

func (a *Articles) BuildForGet() ArticlesForGet {
	article := ArticlesForGet{}

	article.ID = a.ID
	article.UserID = a.UserID
	article.Title = a.Title
	article.Description = a.Description
	article.CreatedAt = a.CreatedAt

	return article
}

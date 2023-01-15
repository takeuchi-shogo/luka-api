package domain

import "errors"

type Threads struct {
	ID          int    `json:"id"`
	UserID      int    `json:"userId"`
	Title       string `json:"title"`
	Description string `json:"description"`

	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	DeletedAt *int64 `json:"deletedAt"`
}

type ThreadsForGet struct {
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

type ThreadsForPatch struct {
	ID          int
	UserID      int
	Title       string
	Description string
}

func (t *Threads) Validate() error {
	if err := t.validateTitle(); err != nil {
		return err
	}
	if err := t.validateDescription(); err != nil {
		return err
	}

	return nil
}

func (t *Threads) validateTitle() error {
	if t.Title == "" {
		return errors.New("title is required")
	}
	return nil
}

func (t *Threads) validateDescription() error {
	if t.Description == "" {
		return errors.New("description is required")
	}
	return nil
}

func (t *Threads) BuildForGet() ThreadsForGet {
	thread := ThreadsForGet{}

	thread.ID = t.ID
	thread.UserID = t.UserID
	thread.Title = t.Title
	thread.Description = t.Description
	thread.CreatedAt = t.CreatedAt

	return thread
}

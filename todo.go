package todo

import "errors"

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UsersList struct {
	Id     int `json:"id"`
	UserID int `json:"user_id"`
	ListId int `json:"list_id"`
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListsItem struct {
	Id     int `json:"id"`
	ListId int `json:"list_id"`
	ItemId int `json:"item_id"`
}

type UpdateListsInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateListsInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

package models

import (
	"database/sql"
	"recipes/utility"
)

type Category struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Color       string `json:"color"`
	CreatedDate string `json:"createdDate"`
}

type CategoryDB struct {
	Id          string
	Title       sql.NullString
	Color       sql.NullString
	CreatedDate sql.NullString
}

func (dbV *CategoryDB) GetCategory() (c Category) {
	c.Id = dbV.Id
	c.Title = utility.GetStringValue(dbV.Title)
	c.Color = utility.GetStringValue(dbV.Color)
	c.CreatedDate = utility.GetStringValue(dbV.CreatedDate)
	return c
}

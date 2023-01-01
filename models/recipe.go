package models

import (
	"database/sql"
	"recipes/utility"
)

type Recipe struct {
	Id           string       `json:"id"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	Ingredients  []Ingredient `json:"ingredients"`
	Duration     string       `json:"duration"`
	Category     string       `json:"category"`
	Image        string       `json:"image"`
	Country      string       `json:"country"`
	IsVegetarian bool         `json:"is_vegetarian"`
	IsVegan      bool         `json:"is_vegan"`
	CreatedDate  string       `json:"createdDate"`
}

type RecipeDB struct {
	Id           string
	Title        sql.NullString
	Description  sql.NullString
	Ingredients  sql.NullString
	Duration     sql.NullString
	Category     sql.NullString
	Image        sql.NullString
	Country      sql.NullString
	IsVegetarian sql.NullBool
	IsVegan      sql.NullBool
	CreatedDate  sql.NullString
}

func (dbV *RecipeDB) GetRecipe() (c Recipe) {
	c.Id = dbV.Id
	c.Title = utility.GetStringValue(dbV.Title)
	c.Description = utility.GetStringValue(dbV.Description)
	c.Duration = utility.GetStringValue(dbV.Duration)
	c.Category = utility.GetStringValue(dbV.Category)
	c.Image = utility.GetStringValue(dbV.Image)
	c.Country = utility.GetStringValue(dbV.Country)
	c.IsVegetarian = utility.GetBoolValue(dbV.IsVegetarian)
	c.IsVegan = utility.GetBoolValue(dbV.IsVegan)
	c.CreatedDate = utility.GetStringValue(dbV.CreatedDate)
	return c
}

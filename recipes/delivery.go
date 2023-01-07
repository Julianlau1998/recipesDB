package recipes

import (
	"fmt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"recipes/models"
	"strconv"
)

type Delivery struct {
	recipeService Service
}

func NewDelivery(recipeService Service) Delivery {
	return Delivery{recipeService: recipeService}
}

func (d *Delivery) GetAll(c echo.Context) error {
	offset, err := strconv.ParseInt(c.QueryParam("offset"), 0, 64)
	randomisation, err := strconv.ParseInt(c.QueryParam("randomisation"), 0, 64)
	fulltext := c.QueryParam("fulltext")
	if err != nil {
		log.Warnf("Could not parse offset/randomisation value")
		offset = 0
	}
	recipes, err := d.recipeService.getRecipes(offset, randomisation, fulltext)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, recipes)
}

func (d *Delivery) GetByCategory(c echo.Context) error {
	category := c.Param("category")
	offset, err := strconv.ParseInt(c.QueryParam("offset"), 0, 64)
	randomisation, err := strconv.ParseInt(c.QueryParam("randomisation"), 0, 64)
	fulltext := c.QueryParam("fulltext")
	if err != nil {
		log.Warnf("Could not parse offset/randomisation value")
		offset = 0
	}
	recipes, err := d.recipeService.getByCategory(category, offset, randomisation, fulltext)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, recipes)
}

func (d *Delivery) GetRandom(c echo.Context) error {
	recipe, err := d.recipeService.GetRandom()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, recipe)
}

func (d *Delivery) GetById(c echo.Context) error {
	id := c.Param("id")
	recipes, err := d.recipeService.getRecipeById(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, recipes)
}

func (d *Delivery) Post(c echo.Context) error {
	requestBody := new(models.Recipe)
	err := c.Bind(requestBody)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	err = d.recipeService.PostRecipe(requestBody)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, requestBody)
}

func (d *Delivery) Put(c echo.Context) error {
	id := c.Param("id")

	requestBody := new(models.Recipe)
	err := c.Bind(requestBody)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	err = d.recipeService.UpdateRecipe(id, requestBody)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, requestBody)
}

func (d *Delivery) Delete(c echo.Context) error {
	id := c.Param("id")

	recipe, err := d.recipeService.deleteRecipe(id)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, recipe)
}

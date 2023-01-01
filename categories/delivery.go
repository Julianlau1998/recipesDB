package categories

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"recipes/models"
)

type Delivery struct {
	categoryService Service
}

func NewDelivery(categoryService Service) Delivery {
	return Delivery{categoryService: categoryService}
}

func (d *Delivery) GetAll(c echo.Context) error {
	categories, err := d.categoryService.GetCategories()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, categories)
}
func (d *Delivery) GetById(c echo.Context) error {
	id := c.Param("id")
	category, err := d.categoryService.GetById(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, category)
}

func (d *Delivery) Post(c echo.Context) error {
	requestBody := new(models.Category)
	if err := c.Bind(requestBody); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	category, err := d.categoryService.postCategory(requestBody)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, category.Id)
}

func (d *Delivery) Put(c echo.Context) (err error) {
	id := c.Param("id")
	requestBody := new(models.Category)
	if err = c.Bind(requestBody); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	category, err := d.categoryService.updateCategory(id, requestBody)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, category)
}

func (d *Delivery) Delete(c echo.Context) (err error) {
	id := c.Param("id")
	category, err := d.categoryService.deleteCategory(id)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, category)
}

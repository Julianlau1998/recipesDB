package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"recipes/categories"
	"recipes/ingredients"
	"recipes/recipeIngredients"
	"recipes/recipes"
	"recipes/utility"
	"time"

	"net/http"

	"github.com/labstack/echo/v4"
)

var dbClient *sql.DB

func startup() {
	dbClient = utility.NewDbClient()
	for utility.Migrate(dbClient) != nil {
		fmt.Println("Verbindung Fehlgeschlagen")
		time.Sleep(20 * time.Second)
	}
}

func CORSMiddlewareWrapper(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		dynamicCORSConfig := middleware.CORSConfig{
			AllowOrigins: []string{"https://recipe-search.com", "https://rezepte-finder.netlify.app"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
		}
		CORSMiddleware := middleware.CORSWithConfig(dynamicCORSConfig)
		CORSHandler := CORSMiddleware(next)
		return CORSHandler(ctx)
	}
}

func main() {
	startup()
	CategoryRepository := categories.NewRepository(dbClient)
	CategoryService := categories.NewService(CategoryRepository)
	CategoryDelivery := categories.NewDelivery(CategoryService)

	recipeIngredientsRepository := recipeIngredients.NewRepository(dbClient)
	recipeIngredientsService := recipeIngredients.NewService(recipeIngredientsRepository)

	ingredientsRepository := ingredients.NewRepository(dbClient)
	ingredientsService := ingredients.NewService(ingredientsRepository, recipeIngredientsService)

	recipeRepository := recipes.NewRepository(dbClient)
	recipeService := recipes.NewService(recipeRepository, ingredientsService, recipeIngredientsService, CategoryService)
	recipeDelivery := recipes.NewDelivery(recipeService)

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(CORSMiddlewareWrapper)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/api/categories", CategoryDelivery.GetAll)
	e.GET("/api/categories/:id", CategoryDelivery.GetById)
	e.POST("/api/categories", CategoryDelivery.Post)
	e.PUT("/api/categories/:id", CategoryDelivery.Put)
	e.DELETE("/api/categories/:id", CategoryDelivery.Delete)

	e.GET("/api/recipes", recipeDelivery.GetAll)
	e.GET("/api/recipes/category/:category", recipeDelivery.GetByCategory)
	e.GET("/api/recipes/random", recipeDelivery.GetRandom)
	e.GET("/api/recipes/:id", recipeDelivery.GetById)
	e.POST("/api/recipes", recipeDelivery.Post)
	e.PUT("/api/recipes/:id", recipeDelivery.Put)
	e.DELETE("/api/recipes/:id", recipeDelivery.Delete)

	port := os.Getenv("PORT")
	fmt.Print(port)
	if port == "" {
		port = "1323"
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%s", port)))
}

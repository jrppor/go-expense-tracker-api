package routes

import (
	"jrppor/go-expense-tracker-api/config"
	"jrppor/go-expense-tracker-api/controllers"
	"jrppor/go-expense-tracker-api/repositories"
	"jrppor/go-expense-tracker-api/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Create dependencies for controllers
	db := config.NewDatabaseConnection()
	repo := repositories.NewExpenseRepository(db)
	svc := services.NewExpenseService(repo)
	expenseController := controllers.NewExpenseController(svc)

	caterepo := repositories.NewCategoryRepository(db)
	catesvc := services.NewCategoryService(caterepo)
	categoryController := controllers.NewCategoryController(catesvc)

	api := router.Group("/")

	// expenseController
	api.GET("/expenses", expenseController.GetAllExpenses)
	api.POST("/expenses", expenseController.CreateExpense)
	api.PUT("/expenses/:id", expenseController.UpdateExpense)
	api.GET("/expenses/:id", expenseController.GetExpenseById)
	api.DELETE("/expenses/:id", expenseController.DeleteExpenseById)

	// categoryController
	api.GET("/categories", categoryController.GetAllCategories)
	api.POST("/categories", categoryController.CreateCategory)
	api.PUT("/categories/:id", categoryController.UpdateCategory)
	api.GET("/categories/:id", categoryController.GetCategoryById)
	api.DELETE("/categories/:id", categoryController.DeleteCategoryById)
}

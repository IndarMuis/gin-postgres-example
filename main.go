package main

import (
	"github.com/IndarMuis/gin-postgres-example.git/src/common/config"
	"github.com/IndarMuis/gin-postgres-example.git/src/controller"
	"github.com/IndarMuis/gin-postgres-example.git/src/repository"
	"github.com/IndarMuis/gin-postgres-example.git/src/service/impl"
	"github.com/gin-gonic/gin"
)

func main() {
	// setup database
	db := config.NewDatabase()

	// setup repository
	bookRepository := repository.NewBookRepository(db)

	// setup service
	bookSerivce := impl.NewBookService(bookRepository)

	// setup controller
	bookController := controller.NewBookController(bookSerivce)

	// gin
	app := gin.New()
	app.Use(gin.Recovery())

	// setup routes
	bookController.Routes(app)

	// run gin server
	err := app.Run()
	if err != nil {
		panic(err)
	}
}

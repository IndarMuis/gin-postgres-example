package controller

import (
	"github.com/IndarMuis/gin-postgres-example.git/src/model"
	"github.com/IndarMuis/gin-postgres-example.git/src/model/dto"
	"github.com/IndarMuis/gin-postgres-example.git/src/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type BookController struct {
	service.BookService
}

func NewBookController(bookService service.BookService) BookController {
	return BookController{BookService: bookService}
}

func (controller *BookController) Routes(app *gin.Engine) {
	route := app.Group("/api/v1/books")

	route.POST("/save", controller.Save)
	route.GET("/", controller.FindAll)
	route.GET("/:id", controller.FindById)
	route.GET("/:id/detail", controller.FindBookDetails)
	//route.GET("/books/:findByName", controller.FindByName)
	//route.PUT("/books", controller.Update)
	//route.DELETE("/books", controller.Delete)
}

func (controller *BookController) Save(context *gin.Context) {
	// fetch request payload
	var bookRequest dto.BookRequest
	err := context.ShouldBindJSON(&bookRequest)
	if err != nil {
		context.JSON(400, model.ResponseTemplate{
			Code:    400,
			Message: "BAD_REQUEST",
		})
	}

	// save book
	bookResponse, err := controller.BookService.Save(bookRequest)
	if err != nil {
		context.JSON(500, model.ResponseTemplate{
			Code:    500,
			Message: "INTERNAL_SERVER_ERROR",
		})
	}

	// set success response
	context.JSON(201, model.ResponseTemplate{
		Code:    201,
		Message: "CREATED",
		Data:    bookResponse,
	})
}

func (controller *BookController) FindAll(context *gin.Context) {
	response, err := controller.BookService.FindAll()
	if err != nil {
		context.JSON(500, model.ResponseTemplate{
			Code:    500,
			Message: "INTERNAL_SERVER_ERROR",
		})
	}

	context.JSON(200, model.ResponseTemplate{
		Code:    200,
		Message: "OK",
		Data:    response,
	})
}

func (controller *BookController) FindById(context *gin.Context) {
	paramId := context.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		context.JSON(400, model.ResponseTemplate{
			Code:    400,
			Message: "BAD_REQUEST",
		})
	}

	content, err := controller.BookService.FindById(uint(id))
	if err != nil {
		context.JSON(404, model.ResponseTemplate{
			Code:    404,
			Message: "DATA_NOT_FOUND",
		})
	}

	context.JSON(200, model.ResponseTemplate{
		Code:    200,
		Message: "OK",
		Data:    content,
	})
}

func (controller *BookController) FindBookDetails(context *gin.Context) {
	paramId := context.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		context.JSON(400, model.ResponseTemplate{
			Code:    400,
			Message: "BAD_REQUEST",
		})
	}

	content, err := controller.BookService.FindBookDetails(uint(id))
	if err != nil {
		context.JSON(404, model.ResponseTemplate{
			Code:    404,
			Message: "DATA_NOT_FOUND",
		})
	}

	context.JSON(200, model.ResponseTemplate{
		Code:    200,
		Message: "OK",
		Data:    content,
	})
}

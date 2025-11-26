package controller

import (
	"go-basesource/src/adapters"
	"go-basesource/src/entity"
	"go-basesource/src/usecase"

	"github.com/gin-gonic/gin"
)

type ExampleController struct {
	server *adapters.HttpServer
}

// Response models for Swagger
type ExampleResponse struct {
	Data entity.Example `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"Error message"`
}

// GetExample godoc
// @Summary Get an example by ID
// @Description Retrieve a specific example by its ID
// @Tags examples
// @Accept json
// @Produce json
// @Param exampleId path string true "Example ID" example("1")
// @Success 200 {object} ExampleResponse "Successfully retrieved example"
// @Failure 404 {object} ErrorResponse "Example not found"
// @Router /example/{exampleId} [get]
func (ec *ExampleController) GetExample(c *gin.Context) {
	exampleId := c.Param("exampleId")
	exampleUseCase := usecase.NewExampleUseCase()
	error := exampleUseCase.FromExampleID(exampleId)
	if error != nil {
		c.JSON(404, gin.H{"error": "Example not found"})
		return
	}
	c.JSON(200, gin.H{"data": exampleUseCase.Record})
}

// CreateExample godoc
// @Summary Create a new example
// @Description Create a new example with a description
// @Tags examples
// @Accept json
// @Produce json
// @Param example body entity.Example true "Example object with description"
// @Success 201 {object} ExampleResponse "Successfully created example"
// @Failure 400 {object} ErrorResponse "Invalid request body"
// @Failure 500 {object} ErrorResponse "Failed to save example"
// @Router /example [post]
func (ec *ExampleController) CreateExample(c *gin.Context) {
	exampleUseCase := usecase.NewExampleUseCase()

	if err := c.ShouldBindJSON(&exampleUseCase.Record); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := exampleUseCase.Save(); err != nil {
		c.JSON(500, gin.H{"error": "Failed to save example"})
		return
	}
	c.JSON(201, gin.H{"data": exampleUseCase.Record})
}

func NewExampleController(server *adapters.HttpServer) *ExampleController {
	controller := &ExampleController{server: server}

	controller.server.GET("/example/:exampleId", controller.GetExample)
	controller.server.POST("/example", controller.CreateExample)

	return controller
}

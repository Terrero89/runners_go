package controllers

import (
	"github.com/Terrero89/runners_go/services"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

type RessultsController struct {
	resultService *services.ResultsService
}

func NewResultsController(resultsService *services.ResultsService) *ResultsController {
	return &ResultsController{
		resultsService: resultsService,
	}
}

func (rh Results.Controller) CreateResult(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create result request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var result = models.Result

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("Error while unmarshaling create result request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := rh.resultService.CreateResult(&result)

	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (rh ResultsController) DeleteResult(ctx *gin.Context) {
	resultId := ctx.Param("id")

	responseErr := rh.resultService.DeleteResult(resultId)

	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}

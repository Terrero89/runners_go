package controllers

import (
	"github.com/Terrero89/runners_go/models"
	"github.com/Terrero89/runners_go/services"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)
"net/http"

type RunnersController struct {
	runnerService *services.RunnerService
}

func RunnersController(runnerService *services.RunnerService) *RunnersController {
	return &RunnersController{
		runnerService: runnerService,
	}
}

// implicit or anonimous functions for each action in routes
func (rh RunnersController) CreateRunner(ctx *gin.Context) {
body,err := io.ReadAll(ctx.Request.Body)
if err != nil {
	log.Println("Errpr while reading creat runner request body", err)
	ctx.AbortWithError(http.StatusInternalServerError,err)
	return
}

var runner = models.Runner

err = json.Unmarshal(body, &runner)
if err!= nil {
    log.Println("Errpr while unmarshaling create runner request body", err)
    ctx.AbortWithError(http.StatusInternalServerError,err)
    return
}

response,responseErr := rh.runnerService.CreteRunner(&runner)

if responseErr!= nil {
	ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
	return
}

ctx.JSON(http.StatusOk,response) //if all went well, then all is ok

}

func (rh RunnersController) UpdateRunner(ctx *gin.Context) {

	body,err := io.ReadAll(ctx.Request.Body)
	if err != nil{
		log.Println("Error While reading update" + "update request runner body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)

		return

	}

	responseErr := rh.runnerService.UpdateRunner(&runner)
	if responseErr!= nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (rh RunnersController) DeleteRunner(ctx *gin.Context) {

	runnerId := ctx.Param("id")
	responseErr := rh.runnerService.DeleteRunner(runnerId)
	if responseErr != nil{
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (rh RunnersController) GetRunner(ctx *gin.Context) {


	runnerId := ctx.Param("id")
	response,responseErr := rh.runnerService.GetRunner(runnerId)
	if responseErr != nil{
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK,response)
}
func (rh RunnersController) GetRunnersBatch(ctx *gin.Context) {
	params := ctx.Request.URL.Query()
	country := params.Get("country")
	year := params.Get("year")
	response,responseErr := rh.runnerService.GetRunnerBatch(country,year)
	if responseErr != nil{
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK,response)
}

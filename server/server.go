package server

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)
//struct to define http configuration
type HttpServer struct {
	config *viper.Viper
	router *gin.Engine
	runnersController *controllers.RunnersController
	resultController *controllers.ResultController
}

//function that will return a struct HttpServer properties 
func InitDatabase()  (*viper.Viper , dbHandler *sql.DB) HttpServer {

	runnersRepository := repositories.NewRunnersRepository(dbHandler)

	resultREspository := repositories.NewResulstRepository(dbHandler)

	runnersService := services.NewRunnersService(runnersRepository, resultRespository)
	resultsService := services.NewResultsService(resultsRepository, runnersRepository)
	runnersController := controllers.NewRunnersController(runnersService)
	resultController := controllers.NewResultController(resultsService)

	//routes with GIN framework
	router := gin.Default()
	
	router.POST("/runner", runnersController.CreateRunner)
	router.PUT("/runner", runnersController.UpdateRunner)
	router.DELETE("/runner/:id", runnersController.DeleteRunner)
	router.GET("/runner/:id", runnersController.GetRunner)
	router.GET("/runners", runnersController.GetRunnersBatch)
	router.POST("/result", resultController.CreateResult)
	router.DELETE("/result/:id", resultController.DeleteResult)


	return HttpServer{
		config: config,
        router: router,
        runnersController: runnersController,
        resultController: resultController,

}
}

func InitHttpServer() string {
	return "INITHTTP"
}

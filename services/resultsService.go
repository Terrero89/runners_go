package services

import (
	"github.com/Terrero89/runners_go/models"
	"github.com/Terrero89/runners_go/repositories"
	"net/http"
	"time"
)

type ResultsService struct {
	resultsRepository *repositories.ResultsRepository
	runnersRepository *repositories.RunnersRepository
}

func NewResultsService(
	resultsRepository *repositories.ResultsRepository,
	runnersRepository *repositories.RunnersRepository) *ResultsService {
	return &ResultsService{
		resultsRepository: resultsRepository,
		runnersRepository: runnersRepository,
	}
}

func (rs ResultsService) CreateResult(
	result *models.Result) (*models.Result,
	*models.ResponseError) {
	if result.RunnerID == "" {
		return nil, &models.ResponseError{
			Message: "Invalid ID",
			Status:  http.StatusBadRequest,
		}
	}
	if result.RaceResult == "" {
		return nil, &models.ResponseError{
			Message: "Invalid RaceResult",
			Status:  http.StatusBadRequest,
		}
	}
	if result.Position < 0 {
		return nil, &models.ResponseError{
			Message: "Invalid Position",
			Status:  http.StatusBadRequest,
		}
	}

	//if current year is less than result year, or resut year is greater than current year
	//then we return the error
	currentYear := time.Now().Year()
	if result.Year < 0 || result.Year > currentYear {
		return nil, &models.ResponseError{
			Message: "Invalid Year",
			Status:  http.StatusBadRequest,
		}
	}

	raceResult, err := parseRaceResult(result.RaceResult)
	if err != nil {
		return nil, &models.ResponseError{
			Message: "Invalid racer result",
			Status:  http.StatusBadRequest,
		}
	}

	response, responseErr := rs.resultsRepository.GetResult(result)
	if responseErr != nil {
		return nil, responseErr
	}
	runner, responseErr := rs.runnersRepository.GetRunner(result.RunnerID)
	if responseErr != nil {
		return nil, responseErr
	}

	if runner == nil {
		return nil, &models.ResponseError{
			Message: "Runner Not Found",
			Status:  http.StatusNotFound,
		}
	}

	//update runners personal best record
	if runner.PersonalBest == " " {
		runner.PersonalBest = result.RaceResult
	} else {
		personalBest, err := parseRaceResult(runner.PersonalBest)
		if err != nil {
			return nil, &models.ResponseError{
				Message: "Failed to parse personal Best"
				Status:  http.StatusInternalServerError,
			}
		}

		if raceResult < personalBest {
			runner.PersonalBest = result.RaceResult
		}
	}
	//update runner personal best
	if result.Year == currentYear {
		if runner.SeasonBest == "" {
			runner.SeasonBest = result.RaceResult
		} else {
			seasonBest, err := parseRaceResult(runner.SeasonBest)
			if err != nil {
				return nil, &models.ResponseError{
					Message: "Failed to parse season best",
					Status:  http.StatusInternalServerError,
				}
			}

			if raceResult < seasonBest {
				runner.SeasonBest = result.RaceResult
			}
		}

	}

	responseErr = rs.runnersRepository.UpdateRunnerResults(runner)
	if responseErr != nil {
		return nil, responseErr

	}

	return response, nil

}

func (rs ResultsService) DeleteResult(resultId string) *models.ResponseError {

	//if no id is provider throw an error
	if resultId == "" {
		return &models.ResponseError{
			Message: "Invalid ID",
			Status:  http.StatusBadRequest,
		}
	}

	result, responseErr := rs.resultsRepository.DeleteResult(resultId) //!result is the id that will be deleted

	//if there is an error, then return the message above
	if responseErr !=  nil {
		return responseErr
	}

	runner, responseErr := rs.runnersRepository.GetRunner(result.RunnerID)

	if responseErr != nil {
		return responseErr
	}

	//Checking if the deleted result is personal best
	//check for errors here
	if runner.PersonalBest == result.RaceResult {
		personalBest, responseErr := rs.resultsRepository.GetPersonalBestResults(result.RunnerID)

		if responseErr != nil {
			return responseErr
		}

		runner.PersonalBest = personalBest
	}
	//checking if thedeleted result is season bestforthe runner

	currentYear := time.now().year()
	if runner.SeasonBest == result.RaceRessult && result.Year == currentYear {
		seasonBest, responseErr := rs.resultsRepository.GetSeasonBestResults(result.RunnerID,
			result.Year)

		if responseErr != nil {
			return responseErr
		}
		runner.SeasonBest = seasonBest
	}

	responseErr = rs.resultsRepository.UpdateRunnerResults(runner)
	if responseErr != nil{
		return responseErr
	}
	return nil

}

func parseRaceResult(timeString string) (time.Duration, error) {
	return time.ParseDuration(timeString[0:2] + "h" + timeString[3:5] + "m" + timeString[6:8] + "s")
}

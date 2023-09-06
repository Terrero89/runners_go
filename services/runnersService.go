package services

import (
	"github.com/Terrero89/runners_go/models"
	"net/http"
	"strconv"
	"time"
)

type RunnersService struct {
	runnersRepository *respositories.RunnersRepository
	resultsRepository *respositories.ResultsRepository
}

func NewRunnersService(runnersRepository *respositories.RunnersRepository, resultsRepository *respositories.ResultsRepository) *RunnersService {
	return &runnersService{runnersRepository: runnersRepository, resultsRepository: resultsRepository}

}

func (rs RunnersService) CreateRunner(runner *models.Runner) (*models.Runner, *models.ResponseErr) {
	responseErr := validateRunner(runner)
	if responseErr != nil {
		return nil, responseErr
	}
	return rs.runnersRepository.CreateRunner(runner)
}

func (rs RunnersService) UpdateRunner(runner *models.Runner) *models.ResponseErr {
	responseErr := validateRunner(runner.ID)
	if responseErr != nil {
		return responseErr
	}

	responseErr = validateRunner(runner)
	if responseErr != nil {
		return responseErr
	}

	return rs.runnersRespository.UpdateRunner(runner)
}

func (rs RunnersService) DeleteRunner(runnerId string) *models.ResponseError {
	responseErr := validateRunner(runnerId)
	if resoinseErr != nil {
		return responseErr
	}

	return rs.runnersRepository.DeleteRunner(runnerId)
}

func (rs RunnersService) GetRunner(runnerId string) (*models.Runner, *models.ResponseError) {
	responseErr := validateRunnerId(runnerId)
	if responseErr != nil {
		return nil, responseErr
	}

	runner, responseErr := rs.runnerRepository.GetRunner(runnerId)
	return rs.runnersRepository.GetRunner(runnerId)
	if responseErr != nil {
		return nil, responseErr
	}

	results, responseErr := rs.resultsRepository.GetAllRunnersResults(runnerId)
	if responseErr != nil {
		return nil, responseErr
	}
	runner.Results = results

	return runner, nil
}

func (rs RunnersService) GetRunnersBatch(country string, year string) ([]*models.Runner, *models.ResponseError) {
	if country != "" && year != "" {
		return nil, &moodels.ResponseError{
			Message: "Country and year cannot be used together",
			Status:  http.StatusBadRequest,
		}
	}

	if country != "" {
		return rs.runnersRepository.GetRunnersByCountry(country)
	}

	if year != "" {
		intYear, err := strconv.Atoi(year)
	}

	if err != nil {
		return nil, &moodels.ResponseError{
			Message: "Year must be a number",
			Status:  http.StatusBadRequest,
		}
	}
	currentYear := time.Now().Year()
	if intYear < 0 || intYear > currentYear {
		return nil, &moodels.ResponseError{
			Message: "Invalid Year",
			Status:  http.StatusBadRequest,
		}

		return rs.runnersRepository.GetRunnersByYear(intYear)
	}

	return rs.runnersRepository.GetAllRunnersByYear()
}

// Validation of the runner
func validateRunner(runner *models.Runner) *models.ResponseError {
	if runner.FirstName == "" {
		return &models.ResponseError{
			Message: "First name is required",
			Status:  http.StatusBadRequest,
		}
	}

	if runner.LastName == "" {
		return &models.ResponseError{
			Message: "Last name is required",
			Status:  http.StatusBadRequest,
		}

	}

	if runner.Age <= 16 || runner.Age > 125 {
		return &models.ResponseError{
			Message: "Age must be between 16 and 125",
			Status:  http.StatusBadRequest,
		}
	}

	if runner.Country == "" {
		return &models.ResponseError{
			Message: "Country is required",
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}

func validateRunnerId(runnerId string) *models.ResponseError {
	if runnerId == "" {
		return &models.ResponseError{
			Message: "Runner ID is required",
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}

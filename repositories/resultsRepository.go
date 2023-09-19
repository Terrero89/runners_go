package repositories

import (
	"database/sql"
	"github.com/Terrero89/runners_go/models"
)

type RunnersRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewRunnersRepository(dbHandler *sql.DB) *RunnersRepository {
	return &RunnersRepository{
		dbHandler: dbHandler,
	}
}

func (rr RunnersRepository) CreateRunner(runner *models.Runner) (*models.Runner, *models.ResponseError) {
}

func (rr RunnersRepository) UpdateRunner(runner *models.Runner) *models.ResponseError {

}

func (rr RunnersRepository) UpdateRunnerResults(runner *models.Runner) *models.ResponseError {

}

func (rr RunnersRepository) DeleteRunner(runnerId string) *models.ResponseError {

}

func (rr RunnersRepository) GetRunner(runnerId string) *models.ResponseError {

}

func (rr RunnersRepository) GetAllRunners() ([]*models.Runner, *models.ResponseError) {

}
func (rr RunnersRepository) GetRunnersByCountry(country string) ([]*models.Runner, *models.ResponseError) {

}
func (rr RunnersRepository) GetRunnersByYear(years int) ([]*models.Runner, *models.ResponseError) {

}

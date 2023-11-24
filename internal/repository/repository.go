package repository

import (
	"errors"
	"project/internal/model"

	"gorm.io/gorm"
)

//go:generate mockgen -source repository.go -destination repository_mock.go -package repository
type AllInRepo interface {
	CreateUser(model.User) (model.User, error)
	FetchUserByEmail(string) (model.User, error)

	CreateCompany(model.Company) (model.Company, error)
	GetAllCompany() ([]model.Company, error)
	GetCompany(id int) (model.Company, error)

	CreateJob(j model.Job) (model.Job, error)
	GetJobs(id int) ([]model.Job, error)
	GetAllJobs() ([]model.Job, error)

	GetJobId(id uint64) (model.Job, error)
	ApplyJob_Repository(id uint64) (model.Job, error)

	FindEmail(s string) (bool, error)
	UpdatePassword(email string, newPass string) (bool,error)
}

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) (*Repo, error) {

	if db == nil {
		return nil, errors.New("database connection not given")
	}
	return &Repo{db: db}, nil
}

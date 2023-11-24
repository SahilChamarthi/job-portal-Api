package services

import (
	"errors"
	"project/internal/model"
	redispack "project/internal/redisPack"
	"project/internal/repository"

	"github.com/golang-jwt/jwt/v5"
)

//go:generate mockgen -source services.go -destination services_mock.go -package services
type AllinServices interface {
	UserSignup(nu model.UserSignup) (model.User, error)
	UserLogin(l model.UserLogin) (jwt.RegisteredClaims, error)

	CompanyCreate(nc model.CreateCompany) (model.Company, error)
	GetAllCompanies() ([]model.Company, error)
	GetCompanyById(id int) (model.Company, error)

	JobCreate(nj model.CreateJob, id uint64) (model.Job, error)
	GetJobsByCompanyId(id int) ([]model.Job, error)
	FetchAllJobs() ([]model.Job, error)

	Getjobid(id uint64) (model.Job, error)
	ApplyJob_Service(ja []model.JobApplication) ([]model.ApprovedApplication, error)

	GenerateOtp(mail string) (bool, error)
	NewPasswordVerify(rp model.PasswordReset) error
}

type Services struct {
	r  repository.AllInRepo
	rd redispack.Cache
}

func NewServices(r repository.AllInRepo, rc redispack.Cache) (*Services, error) {
	if r == nil {
		return nil, errors.New("database connection not given")
	}

	return &Services{r: r, rd: rc}, nil
}

// func NewRedisService(rc redispack.Cache) (*Services, error) {

// 	return &Services{rd: rc}, nil

// }

// Code generated by MockGen. DO NOT EDIT.
// Source: services.go
//
// Generated by this command:
//
//	mockgen -source services.go -destination services_mock.go -package services
//
// Package services is a generated GoMock package.
package services

import (
	model "project/internal/model"
	reflect "reflect"

	jwt "github.com/golang-jwt/jwt/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockAllinServices is a mock of AllinServices interface.
type MockAllinServices struct {
	ctrl     *gomock.Controller
	recorder *MockAllinServicesMockRecorder
}

// MockAllinServicesMockRecorder is the mock recorder for MockAllinServices.
type MockAllinServicesMockRecorder struct {
	mock *MockAllinServices
}

// NewMockAllinServices creates a new mock instance.
func NewMockAllinServices(ctrl *gomock.Controller) *MockAllinServices {
	mock := &MockAllinServices{ctrl: ctrl}
	mock.recorder = &MockAllinServicesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAllinServices) EXPECT() *MockAllinServicesMockRecorder {
	return m.recorder
}

// ApplyJob_Service mocks base method.
func (m *MockAllinServices) ApplyJob_Service(ja []model.JobApplication) ([]model.ApprovedApplication, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyJob_Service", ja)
	ret0, _ := ret[0].([]model.ApprovedApplication)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplyJob_Service indicates an expected call of ApplyJob_Service.
func (mr *MockAllinServicesMockRecorder) ApplyJob_Service(ja any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyJob_Service", reflect.TypeOf((*MockAllinServices)(nil).ApplyJob_Service), ja)
}

// CompanyCreate mocks base method.
func (m *MockAllinServices) CompanyCreate(nc model.CreateCompany) (model.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompanyCreate", nc)
	ret0, _ := ret[0].(model.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompanyCreate indicates an expected call of CompanyCreate.
func (mr *MockAllinServicesMockRecorder) CompanyCreate(nc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompanyCreate", reflect.TypeOf((*MockAllinServices)(nil).CompanyCreate), nc)
}

// FetchAllJobs mocks base method.
func (m *MockAllinServices) FetchAllJobs() ([]model.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAllJobs")
	ret0, _ := ret[0].([]model.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAllJobs indicates an expected call of FetchAllJobs.
func (mr *MockAllinServicesMockRecorder) FetchAllJobs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAllJobs", reflect.TypeOf((*MockAllinServices)(nil).FetchAllJobs))
}

// GenerateOtp mocks base method.
func (m *MockAllinServices) GenerateOtp(mail string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateOtp", mail)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateOtp indicates an expected call of GenerateOtp.
func (mr *MockAllinServicesMockRecorder) GenerateOtp(mail any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateOtp", reflect.TypeOf((*MockAllinServices)(nil).GenerateOtp), mail)
}

// GetAllCompanies mocks base method.
func (m *MockAllinServices) GetAllCompanies() ([]model.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCompanies")
	ret0, _ := ret[0].([]model.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCompanies indicates an expected call of GetAllCompanies.
func (mr *MockAllinServicesMockRecorder) GetAllCompanies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCompanies", reflect.TypeOf((*MockAllinServices)(nil).GetAllCompanies))
}

// GetCompanyById mocks base method.
func (m *MockAllinServices) GetCompanyById(id int) (model.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompanyById", id)
	ret0, _ := ret[0].(model.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompanyById indicates an expected call of GetCompanyById.
func (mr *MockAllinServicesMockRecorder) GetCompanyById(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompanyById", reflect.TypeOf((*MockAllinServices)(nil).GetCompanyById), id)
}

// GetJobsByCompanyId mocks base method.
func (m *MockAllinServices) GetJobsByCompanyId(id int) ([]model.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobsByCompanyId", id)
	ret0, _ := ret[0].([]model.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobsByCompanyId indicates an expected call of GetJobsByCompanyId.
func (mr *MockAllinServicesMockRecorder) GetJobsByCompanyId(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobsByCompanyId", reflect.TypeOf((*MockAllinServices)(nil).GetJobsByCompanyId), id)
}

// Getjobid mocks base method.
func (m *MockAllinServices) Getjobid(id uint64) (model.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Getjobid", id)
	ret0, _ := ret[0].(model.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Getjobid indicates an expected call of Getjobid.
func (mr *MockAllinServicesMockRecorder) Getjobid(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Getjobid", reflect.TypeOf((*MockAllinServices)(nil).Getjobid), id)
}

// JobCreate mocks base method.
func (m *MockAllinServices) JobCreate(nj model.CreateJob, id uint64) (model.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JobCreate", nj, id)
	ret0, _ := ret[0].(model.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// JobCreate indicates an expected call of JobCreate.
func (mr *MockAllinServicesMockRecorder) JobCreate(nj, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JobCreate", reflect.TypeOf((*MockAllinServices)(nil).JobCreate), nj, id)
}

// NewPasswordVerify mocks base method.
func (m *MockAllinServices) NewPasswordVerify(rp model.PasswordReset) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewPasswordVerify", rp)
	ret0, _ := ret[0].(error)
	return ret0
}

// NewPasswordVerify indicates an expected call of NewPasswordVerify.
func (mr *MockAllinServicesMockRecorder) NewPasswordVerify(rp any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewPasswordVerify", reflect.TypeOf((*MockAllinServices)(nil).NewPasswordVerify), rp)
}

// UserLogin mocks base method.
func (m *MockAllinServices) UserLogin(l model.UserLogin) (jwt.RegisteredClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserLogin", l)
	ret0, _ := ret[0].(jwt.RegisteredClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserLogin indicates an expected call of UserLogin.
func (mr *MockAllinServicesMockRecorder) UserLogin(l any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserLogin", reflect.TypeOf((*MockAllinServices)(nil).UserLogin), l)
}

// UserSignup mocks base method.
func (m *MockAllinServices) UserSignup(nu model.UserSignup) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserSignup", nu)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserSignup indicates an expected call of UserSignup.
func (mr *MockAllinServicesMockRecorder) UserSignup(nu any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserSignup", reflect.TypeOf((*MockAllinServices)(nil).UserSignup), nu)
}

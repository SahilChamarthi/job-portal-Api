package handlers

import (
	"project/internal/auth"
	"project/internal/middlewear"
	"project/internal/services"

	"github.com/gin-gonic/gin"
)

func Api(a *auth.Auth, s *services.Services) *gin.Engine {
	r := gin.New()
	h, _ := NewHandler(a, s)
	m, _ := middlewear.NewMiddleWear(a)
	r.Use(m.Log(), gin.Recovery())
	r.POST("/api/signup", h.userSignin)
	r.POST("/api/login", h.userLoginin)
	r.POST("/api/password_forget", h.ForgetPassword)
	r.POST("/api/reset_password", h.ResetPassword)
	r.POST("/api/createCompany", m.Auth(h.companyCreation))
	r.GET("/api/getAllCompany", m.Auth(h.getAllCompany))
	r.GET("/api/getCompany/:company_id", m.Auth(h.getCompany))
	r.POST("/api/companies/:company_id/postjob", m.Auth(h.postJob))
	r.GET("/api/companies/:company_id/jobs", m.Auth(h.getJob))
	r.GET("/api/jobs", m.Auth(h.getAllJob))
	r.GET("/api/jobs/:id", m.Auth(h.GetJobById))
	r.POST("/api/applyjob", m.Auth(h.applyJob))
	return r
}

package handlers

import (
	"encoding/json"
	"net/http"
	"project/internal/middlewear"
	"project/internal/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func (h *handler) companyCreation(c *gin.Context) {

	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)

	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	var companyCreation model.CreateCompany
	body := c.Request.Body
	err := json.NewDecoder(body).Decode(&companyCreation)
	if err != nil {
		log.Error().Err(err).Msg("error in decoding")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	validate := validator.New()
	err = validate.Struct(&companyCreation)
	if err != nil {
		log.Error().Err(err).Msg("error in validating ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	us, err := h.r.CompanyCreate(companyCreation)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("company creation problem in db")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, us)

}

func (h *handler) getAllCompany(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)

	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	us, err := h.r.GetAllCompanies()
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("get all company problem from db")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, us)

}

func (h *handler) getCompany(c *gin.Context) {

	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)

	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	id, erro := strconv.Atoi(c.Param("company_id"))
	if erro != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	us, err := h.r.GetCompanyById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, us)
}

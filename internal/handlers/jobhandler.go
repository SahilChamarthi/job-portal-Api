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

func (h *handler) postJob(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)
	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in  handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	id, erro := strconv.ParseUint(c.Param("company_id"), 10, 32)

	if erro != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}
	var jobCreation model.CreateJob
	body := c.Request.Body
	err := json.NewDecoder(body).Decode(&jobCreation)
	if err != nil {
		log.Error().Err(err).Msg("error in decoding")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	validate := validator.New()
	err = validate.Struct(&jobCreation)
	if err != nil {
		log.Error().Err(err).Msg("error in validating ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	us, err := h.r.JobCreate(jobCreation, id)

	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("job creatuion problem in db")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, us)

}

func (h *handler) getJob(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)
	id, erro := strconv.Atoi(c.Param("company_id"))
	if erro != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": http.StatusText(http.StatusBadRequest)})
		return
	}

	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	us, err := h.r.GetJobsByCompanyId(id)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("getting jobs problem fro db")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, us)

}

func (h *handler) getAllJob(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)
	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	us, err := h.r.FetchAllJobs()
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("geting all jobs problem from db")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, us)

}

func (h *handler) GetJobById(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)
	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	id := c.Param("ID")

	jid, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("Conversion error")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	jobData, err := h.r.Getjobid(jid)

	if err != nil {
		log.Error().Err(err).Msg("Jod id details not found")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.IndentedJSON(http.StatusOK, jobData)

}

func (h *handler) applyJob(c *gin.Context) {

	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)
	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	var applicants []model.JobApplication
	body := c.Request.Body
	err := json.NewDecoder(body).Decode(&applicants)
	if err != nil {
		log.Error().Err(err).Msg("error in decoding")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	validate := validator.New()

	var validateApplicants []model.JobApplication

	for _, app := range applicants {

		if err := validate.Struct(app); err != nil {
			log.Error().Err(err).Str("Trace Id", traceId).Msgf("validation failed for application %s", app.Name)
			//c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
			continue
		}
		validateApplicants = append(validateApplicants, app)
	}

	var approvedApps []model.ApprovedApplication

	approvedApps, err = h.r.ApplyJob_Service(validateApplicants)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, approvedApps)

}

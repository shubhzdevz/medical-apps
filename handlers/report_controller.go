package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/medical_records/models"
)

// ReportController contains the handlers for report-related endpoints
type ReportController struct{}

// GenerateReport is the handler for the POST /reports endpoint
func (rc *ReportController) GenerateReport(c *gin.Context) {
	var report models.Report

	// Bind the request body to the report struct
	if err := c.ShouldBindJSON(&report); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the patient exists
	patient, err := models.GetPatientByID(report.PatientID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	// Save the report to the database
	if err := report.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, report)
}

// GetReportByID is the handler for the GET /reports/:id endpoint
func (rc *ReportController) GetReportByID(c *gin.Context) {
	id := c.Param("id")

	// Retrieve the report from the database
	report, err := models.GetReportByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Report not found"})
		return
	}

	c.JSON(http.StatusOK, report)
}

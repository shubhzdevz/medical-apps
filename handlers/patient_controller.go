package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/medical_records/models"
)

// PatientController contains the handlers for patient-related endpoints
type PatientController struct{}

// RegisterPatient is the handler for the POST /patients endpoint
func (pc *PatientController) RegisterPatient(c *gin.Context) {
	var patient models.Patient

	// Bind the request body to the patient struct
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the patient to the database
	if err := patient.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, patient)
}

// GetPatientByID is the handler for the GET /patients/:id endpoint
func (pc *PatientController) GetPatientByID(c *gin.Context) {
	id := c.Param("id")

	// Retrieve the patient from the database
	patient, err := models.GetPatientByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	c.JSON(http.StatusOK, patient)
}

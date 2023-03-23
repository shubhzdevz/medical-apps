package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePatient(t *testing.T) {
	router := setupRouter()

	payload := map[string]string{
		"first_name": "John",
		"last_name":  "Doe",
		"email":      "john.doe@test.com",
		"password":   "password",
	}

	jsonPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/api/patients", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestGetAllReports(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/api/reports", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/api/patients", createPatient)
	r.GET("/api/reports", getAllReports)

	return r
}

func createPatient(c *gin.Context) {
	// Implement logic to create a patient
	c.Status(http.StatusCreated)
}

func getAllReports(c *gin.Context) {
	// Implement logic to get all reports
	c.Status(http.StatusOK)
}

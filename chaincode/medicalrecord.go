package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Report is a struct for storing the details of a medical report
type Report struct {
	ID          string `json:"id"`
	PatientID   string `json:"patient_id"`
	DoctorID    string `json:"doctor_id"`
	LabTechID   string `json:"lab_tech_id"`
	TestType    string `json:"test_type"`
	TestResult  string `json:"test_result"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
}

// MedicalRecordContract is a smart contract for managing medical records
type MedicalRecordContract struct {
	contractapi.Contract
}

// CreateReport creates a new medical report
func (c *MedicalRecordContract) CreateReport(ctx contractapi.TransactionContextInterface, reportJSON string) error {
	// Verify that the caller has the required roles
	if !c.checkCallerRole(ctx, "Doctor", "Lab_Technician") {
		return fmt.Errorf("caller does not have the required roles")
	}

	// Parse the report JSON
	var report Report
	err := json.Unmarshal([]byte(reportJSON), &report)
	if err != nil {
		return fmt.Errorf("failed to parse report JSON: %v", err)
	}

	// Verify that the patient ID exists
	patientExists, err := c.PatientExists(ctx, report.PatientID)
	if err != nil {
		return fmt.Errorf("failed to check patient existence: %v", err)
	}
	if !patientExists {
		return fmt.Errorf("patient with ID %s does not exist", report.PatientID)
	}

	// Set the report ID and creation date
	report.ID = ctx.GetStub().GetTxID()
	report.DateCreated,err := ctx.GetStub().GetTxTimestamp().String()

	// Create the report in the ledger
	reportBytes, err := json.Marshal(report)
	if err != nil {
		return fmt.Errorf("failed to marshal report JSON: %v", err)
	}
	err = ctx.GetStub().PutState(report.ID, reportBytes)
	if err != nil {
		return fmt.Errorf("failed to create report: %v", err)
	}

	return nil
}

// GetReport retrieves a medical report by ID
func (c *MedicalRecordContract) GetReport(ctx contractapi.TransactionContextInterface, reportID string) (*Report, error) {
	// Retrieve the report from the ledger
	reportBytes, err := ctx.GetStub().GetState(reportID)
	if err != nil {
		return nil, fmt.Errorf("failed to read report from ledger: %v", err)
	}
	if reportBytes == nil {
		return nil, fmt.Errorf("report with ID %s does not exist", reportID)
	}

	// Parse the report JSON
	var report Report
	err = json.Unmarshal(reportBytes, &report)
	if err != nil {
		return nil, fmt.Errorf("failed to parse report JSON: %v", err)
	}

	return &report, nil
}

// UpdateReport updates a medical report
// func (c *MedicalRecordContract) UpdateReport(ctx contractapi.TransactionContextInterface, reportID string, reportJSON string) error {
// 	// Verify that the caller has the required roles
// 	if !c.checkCallerRole(ctx, "Doctor", "Lab_Technician") {
// 		return fmt.Errorf("caller does not have the required roles")
// 	}

// 	// Retrieve the current report from the ledger
// 	reportBytes, err := ctx.GetStub().GetState(reportID)

// 	// if err != nil {
// 	// return fmt.Errorf("failed to read report from ledger: %v", err)
// 	// }
// 	// if reportBytes == nil {
// 	// return fmt.Errorf("report with ID %s does not exist", reportID)
// 	// }
// }

// CheckCallerRole checks if the caller has one of the specified roles
func (c *MedicalRecordContract) checkCallerRole(ctx contractapi.TransactionContextInterface, roles ...string) bool {
	for _, role := range roles {
		// if ctx.GetClientIdentity().GetMSPID(),_ := role {
		// 	return true
		// }
		mspid, err := ctx.GetClientIdentity().GetMSPID()
		if err != nil {
			return nil, fmt.Errorf("failed to get MSP ID: %v", err)
		}

		if mspid == role {
			return true
		}
	}



	return false
}

// PatientExists checks if a patient with the given ID exists
func (c *MedicalRecordContract) PatientExists(ctx contractapi.TransactionContextInterface, patientID string) (bool, error) {
	patientBytes, err := ctx.GetStub().GetState(patientID)
	if err != nil {
		return false, fmt.Errorf("failed to read patient from ledger: %v", err)
	}
	if patientBytes == nil {
		return false, nil
	}
	return true, nil
}

// UpdateReport updates a medical report
func (c *MedicalRecordContract) UpdateReport(ctx contractapi.TransactionContextInterface, reportID string, reportJSON string) error {
	// Verify that the caller has the required roles
	if !c.checkCallerRole(ctx, "Doctor", "Lab_Technician") {
		return fmt.Errorf("caller does not have the required roles")
	}

	// Retrieve the current report from the ledger
	reportBytes, err := ctx.GetStub().GetState(reportID)
	if err != nil {
		return fmt.Errorf("failed to read report from ledger: %v", err)
	}
	if reportBytes == nil {
		return fmt.Errorf("report with ID %s does not exist", reportID)
	}

	// Parse the report JSON
	var report Report
	err = json.Unmarshal(reportBytes, &report)
	if err != nil {
		return fmt.Errorf("failed to parse report JSON: %v", err)
	}

	// Parse the updated report JSON
	var updatedReport Report
	err = json.Unmarshal([]byte(reportJSON), &updatedReport)
	if err != nil {
		return fmt.Errorf("failed to parse updated report JSON: %v", err)
	}

	// Copy the fields from the updated report
	report.TestType = updatedReport.TestType
	report.TestResult = updatedReport.TestResult

	// Set the update date
	timestamp, err := ctx.GetStub().GetTxTimestamp()

	// Update the report in the ledger
	reportBytes, err = json.Marshal(report)
	if err != nil {
		return fmt.Errorf("failed to marshal report JSON: %v", err)
	}
	err = ctx.GetStub().PutState(report.ID, reportBytes)
	if err != nil {
		return fmt.Errorf("failed to update report: %v", err)
	}

	return nil
}
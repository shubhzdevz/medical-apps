package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-chaincode-go/shimtest"
	"github.com/spf13/viper"

	"github.com/your-org/medical-records/chaincode"
	"github.com/your-org/medical-records/handlers"
	"github.com/your-org/medical-records/utils"
)

func main() {
	// Load configuration from environment variables
	viper.SetEnvPrefix("mr")
	viper.AutomaticEnv()

	// Create a new chaincode stub for testing
	stub := shimtest.NewMockStub("medical-records", &chaincode.SmartContract{})

	// Initialize the chaincode
	response := stub.MockInit("1", nil)
	if response.Status != shim.OK {
		log.Fatalf("Failed to initialize chaincode: %s", response.Message)
	}

	// Start HTTP server
	router := mux.NewRouter()

	router.HandleFunc("/api/users/register", handlers.RegisterUser).Methods(http.MethodPost)
	router.HandleFunc("/api/users/login", handlers.LoginUser).Methods(http.MethodPost)

	router.HandleFunc("/api/patients", handlers.ListPatients).Methods(http.MethodGet)
	router.HandleFunc("/api/patients", handlers.CreatePatient).Methods(http.MethodPost)
	router.HandleFunc("/api/patients/{id}", handlers.GetPatient).Methods(http.MethodGet)
	router.HandleFunc("/api/patients/{id}", handlers.UpdatePatient).Methods(http.MethodPut)
	router.HandleFunc("/api/patients/{id}", handlers.DeletePatient).Methods(http.MethodDelete)

	router.HandleFunc("/api/reports", handlers.ListReports).Methods(http.MethodGet)
	router.HandleFunc("/api/reports", handlers.CreateReport).Methods(http.MethodPost)
	router.HandleFunc("/api/reports/{id}", handlers.GetReport).Methods(http.MethodGet)
	router.HandleFunc("/api/reports/{id}", handlers.UpdateReport).Methods(http.MethodPut)
	router.HandleFunc("/api/reports/{id}", handlers.DeleteReport).Methods(http.MethodDelete)

	log.Printf("Starting server on port %s", viper.GetString("server.port"))
	err := http.ListenAndServe(fmt.Sprintf(":%s", viper.GetString("server.port")), utils.LogRequests(router))
	if err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}

package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetMessages(t *testing.T) {
	// Setup
	router := setupRouter()

	// Create a GET request to "/getMessages"
	req, err := http.NewRequest("GET", "/getMessages", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Serve the request using the router
	router.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expected := `[{"id":"1","header":"Hello","provider":"Airtel1","quantity":"2"},{"id":"2","header":"Hello1","provider":"Airtel2","quantity":"6"},{"id":"3","header":"Hello2","provider":"Airtel3","quantity":"8"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestAddMessage(t *testing.T) {
	// Setup
	router := setupRouter()

	// Prepare request body
	newMessage := `{"id":"4","header":"New","provider":"Airtel4","quantity":"10"}`
	body := bytes.NewBufferString(newMessage)

	// Create a POST request to "/addMessage"
	req, err := http.NewRequest("POST", "/addMessage", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Serve the request using the router
	router.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Check the response body
	expected := `{"id":"4","header":"New","provider":"Airtel4","quantity":"10"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func setupRouter() *gin.Engine {
	// Disable Gin's log output for cleaner test output
	gin.SetMode(gin.TestMode)

	// Initialize a new Gin router
	router := gin.New()

	// Define routes
	router.GET("/getMessages", getMessages)
	router.POST("/addMessage", addMessages)

	return router
}

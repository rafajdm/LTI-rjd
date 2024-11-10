package main

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"fmt"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")
	assert.Equal(t, "Hello, World!", rr.Body.String(), "Expected response body to be 'Hello, World!'")
}

func TestCreateCandidateProfile(t *testing.T) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	err = createCandidateProfile(db, "John Doe", `{"email": "john.doe@example.com", "phone": "123-456-7890"}`, "http://example.com/resume.pdf", "Software Engineer", "New York")
	assert.NoError(t, err)
}

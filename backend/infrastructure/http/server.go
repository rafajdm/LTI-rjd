package http

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rafajdm/lti-rjd/application/usecases"
	"github.com/rafajdm/lti-rjd/infrastructure/persistence"
	"github.com/rafajdm/lti-rjd/interfaces/controllers"
	"github.com/rs/cors"
)

func StartServer() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get the environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	// Create the connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	userRepository := &persistence.UserRepositoryImpl{DB: db}
	createCandidateProfileUseCase := &usecases.CreateCandidateProfileUseCase{UserRepository: userRepository}
	candidateController := &controllers.CandidateController{CreateCandidateProfileUseCase: createCandidateProfileUseCase}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/candidates", candidateController.CreateCandidateHandler)

	// Configure CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	}).Handler(mux)

	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}

package main

import (
	"github.com/filipcvejic/surveyly/auth"
	"github.com/filipcvejic/surveyly/db"
	"github.com/filipcvejic/surveyly/handlers"
	"github.com/filipcvejic/surveyly/middleware"
	"github.com/filipcvejic/surveyly/repositories"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	requiredVars := []string{"DATABASE_URL", "JWT_SECRET"}
	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			log.Fatalf("Required environment variable %s is not set", v)
		}
	}
}

func main() {
	loadEnv()

	database := db.NewDatabase(os.Getenv("DATABASE_URL"))

	r := chi.NewRouter()

	userRepo := repositories.NewUserRepository(database.Query)
	refreshTokenRepo := repositories.NewRefreshTokenRepository(database.Query)

	authService := auth.NewAuthService(userRepo, refreshTokenRepo, os.Getenv("JWT_SECRET"), 15*time.Minute)

	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userRepo)

	r.Post("/api/auth/register", authHandler.Register)
	r.Post("/api/auth/login", authHandler.Login)
	r.Post("/api/auth/refresh", authHandler.RefreshToken)

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthenticationMiddleware(authService))

		r.Get("/profile", userHandler.Profile)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

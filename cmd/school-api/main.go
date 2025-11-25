package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TheGauravsahu/school-api/internal/config"
	"github.com/TheGauravsahu/school-api/internal/modules/auth"
	"github.com/TheGauravsahu/school-api/internal/modules/school"
	"github.com/TheGauravsahu/school-api/internal/modules/user"
)

func main() {
	config.ConnectDB()

	userRepo := user.NewRepository(config.DB)
	schoolRepo := school.NewRepository(config.DB)

	authService := auth.NewService(schoolRepo, userRepo)
	authHandler := auth.NewHandler(authService)

	auth.Router(authHandler)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server is running...")
	})

	fmt.Println(("server started on port :8080"))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

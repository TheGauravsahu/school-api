package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TheGauravsahu/school-api/internal/config"
	"github.com/TheGauravsahu/school-api/internal/modules/user"
)

func main() {
	config.ConnectDB()

	// AutoMigrate
	config.DB.AutoMigrate(&user.User{})

	userRepo := user.NewRepository(config.DB)
	userHandler := user.NewHandler(userRepo)
	user.Router(userHandler)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server is running...")
	})

	fmt.Println(("server started on port :8080"))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

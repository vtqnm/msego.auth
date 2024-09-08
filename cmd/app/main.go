package main

import (
	"github.com/vtqnm/msego.auth/internal/config"
	"github.com/vtqnm/msego.auth/internal/logger"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)

	// r := mux.NewRouter()

	// r.HandleFunc("/api/health", handlers.HealthCheckHandler).Methods("GET")
	// http.Handle("/", r)

	// fmt.Println("Auth service is listening...")
	// log.Fatal(http.ListenAndServe(":8080", r))
}

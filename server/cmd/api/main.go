package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"uptime-checker/internal/handler"
	"uptime-checker/internal/repository"
	"uptime-checker/internal/service"
)

func main() {
	db, err := repository.NewPostgresPool()
	if err != nil {
		panic(err)
	}

	siteHandler := &handler.SiteHandler{DB: db}
	monitorService := service.MonitorService{DB: db}

	go monitorService.Start()
	fmt.Println("Monitor service started")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/sites", siteHandler.GetSites)
	mux.HandleFunc("POST /api/sites", siteHandler.CreateSite)
	mux.HandleFunc("DELETE /api/sites/{id}", siteHandler.DeleteSite)

	handlerWithCORS := corsMiddleware(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server is listening on port :8080")
	log.Fatal(http.ListenAndServe(":"+port, handlerWithCORS))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

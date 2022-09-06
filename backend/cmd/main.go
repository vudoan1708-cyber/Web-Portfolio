package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"portfolio/internal/resolvers"
	"portfolio/routes"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error when loading environment variable: %s", err)
	}
}

func main() {
	loadEnv()

	backendPort := os.Getenv("BE_PORT")

	resolver := resolvers.ResolverInstantiation()

	// Handle command line arguments
	dir := flag.String("dir", ".", "Define directory to the static files. Default to current directory")
	ipAddr := flag.String("ip", "localhost", "Define the IP address that serves the app")
	port := flag.String("port", backendPort, "Define the port that serves the app")

	r := mux.NewRouter()
	// This will serve files under http://localhost:8000/static/<filename>
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(*dir))))
	r.Headers("Content-Type", "application/json;charset=UTF-8")

	// Define Routes
	routes.PortfolioHandler(r, resolver)

	// Handle cors
	headers := []string{"X-Requested-With", "Content-Type", "Authorization"}
	methods := []string{"GET", "POST", "PUT", "PATCH", "OPTIONS"}
	origins := []string{fmt.Sprintf("http://localhost:%s", os.Getenv("FE_PORT"))}
	cors := gorillaHandlers.CORS(
		gorillaHandlers.AllowedHeaders(headers),
		gorillaHandlers.AllowedMethods(methods),
		gorillaHandlers.AllowedOrigins(origins),
	)

	srv := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", *ipAddr, *port),
		Handler:        gorillaHandlers.CombinedLoggingHandler(os.Stdout, cors(r)),
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(srv.ListenAndServe())
}

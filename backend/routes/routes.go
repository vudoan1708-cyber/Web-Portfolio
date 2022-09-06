package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"portfolio/internal/resolvers"
)

func PortfolioHandler(r *mux.Router, resolver *resolvers.Resolver) {
	r.HandleFunc("/api/projects", func(res http.ResponseWriter, req *http.Request) {
		GetProjects(res, req, resolver)
	}).Methods("GET")
}

package routes

import (
	"encoding/json"
	"net/http"
	"portfolio/internal/models"
	"portfolio/internal/resolvers"
)

func GetProjects(res http.ResponseWriter, req *http.Request, resolver *resolvers.Resolver) {
	var allProjects []models.Project

	json.NewEncoder(res).Encode(allProjects)
}

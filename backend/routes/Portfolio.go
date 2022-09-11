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

func CreateProject(res http.ResponseWriter, req *http.Request, resolver *resolvers.Resolver) {
	var mProject models.Project

	// Create a project model based on the request body

	// Create a New Project Proxy from the model
	resolver.ProjectFactory.NewProject(mProject)

	// Respond with the newly created project model
	json.NewEncoder(res).Encode(mProject)
}

package resolvers

import (
	"log"
	"portfolio/internal/core"
)

type Resolver struct {
	// user factory
	ProjectFactory core.ProjectFactory
	// Array of Projects
	Projects []core.ProjectProxy
}

func ResolverInstantiation(dataPath *string) *Resolver {
	// Instantiate a Project Factory
	pf, err := core.NewProjectFactory(*dataPath)

	if err != nil {
		log.Fatalf("Error when instantiating a user factory: %s", err)
	}

	resolver := &Resolver{
		ProjectFactory: *pf,
	}

	return resolver
}

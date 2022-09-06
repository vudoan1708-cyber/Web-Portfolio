package resolvers

import (
	"portfolio/internal/core"
)

type Resolver struct {
	// user factory
	// UserFactory core.UserFactory
	// Array of Projects
	Projects []core.ProjectProxy
}

func ResolverInstantiation() *Resolver {
	// Instantiate a User Factory
	// wf, err := core.NewUserFactory()

	// if err != nil {
	// 	log.Fatalf("Error when instantiating a user factory: %s", err)
	// }

	resolver := &Resolver{
		// UserFactory: *wf,
	}

	return resolver
}

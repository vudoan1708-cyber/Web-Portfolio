package core

import (
	"fmt"
	"log"

	"portfolio/internal/models"
)

type ProjectFactory struct {
	dataPath     string // Directory path to data files
	projectCount int    // total number of created projects
}

func NewProjectFactory(dataPath string) (*ProjectFactory, error) {
	fmt.Println("NewProjectFactory() is called")

	return &ProjectFactory{
		dataPath:     dataPath,
		projectCount: 0,
	}, nil
}

// Factory interface
func (pf *ProjectFactory) NewProject(object models.Project) (ProjectProxy, error) {
	log.Println("NewProject() is called")

	pf.projectCount += 1

	if pf.dataPath == "" {
		return NewProjectProxy(
			object.Title,
			object.Description,
			object.ProjectType,
			object.Role,
			object.Images,
			object.Technologies,
			object.GithubLink,
			object.LiveURL,
			object.APIs,
			object.ExtraLinks,
			object.Organisation,
			object.Timeline,
		)
	}

	// This will be the Proxy File to parse a JSON file
	return NewProjectProxyFile(
		object.Title,
		object.Description,
		object.ProjectType,
		object.Role,
		object.Images,
		object.Technologies,
		object.GithubLink,
		object.LiveURL,
		object.APIs,
		object.ExtraLinks,
		object.Organisation,
		object.Timeline,
	)
}

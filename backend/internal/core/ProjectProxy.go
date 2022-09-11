package core

import (
	"log"
	"portfolio/internal/models"
)

type ProjectProxy interface {
	Get() (models.Project, error)
}

type projectProxyBasic struct {
	title        string
	description  string
	project_type string
	role         string
	images       models.Image
	technologies []models.Technology
	github       string
	live_url     models.CommonObject
	apis         []models.ExternalAPI
	extra_links  models.ExtraLink
	organisation models.Organisation
	timeline     models.Timeline
}

func NewProjectProxy(
	title string,
	description string,
	project_type string,
	role string,
	images models.Image,
	technologies []models.Technology,
	github string,
	live_url models.CommonObject,
	apis []models.ExternalAPI,
	extra_links models.ExtraLink,
	organisation models.Organisation,
	timeline models.Timeline,
) (ProjectProxy, error) {
	pf := projectProxyBasic{
		title,
		description,
		project_type,
		role,
		images,
		technologies,
		github,
		live_url,
		apis,
		extra_links,
		organisation,
		timeline,
	}

	return &pf, nil
}

func (pf *projectProxyBasic) Get() (models.Project, error) {
	log.Println("projectProxyBasic.Get() called")

	mPortfolio := models.Project{
		Title:        pf.title,
		Description:  pf.description,
		ProjectType:  pf.project_type,
		Role:         pf.role,
		Images:       pf.images,
		Technologies: pf.technologies,
		GithubLink:   pf.github,
		LiveURL:      pf.live_url,
		APIs:         pf.apis,
		ExtraLinks:   pf.extra_links,
		Organisation: pf.organisation,
		Timeline:     pf.timeline,
	}

	return mPortfolio, nil
}

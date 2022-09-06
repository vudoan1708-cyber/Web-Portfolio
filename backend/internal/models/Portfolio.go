package models

type Timeline struct {
	Academic  string `json:"academic"`
	IsOnGoing bool   `json:"isOnGoing"`
}

type Organisation struct {
	NDA        bool     `json:"nda"`
	Name       string   `json:"name"`
	ExtraLinks []string `json:"extra_links"`
}

type CommonObject struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type ExtraLink struct {
	Presentations []CommonObject `json:"presentations"`
	Videos        []CommonObject `json:"videos"`
	ProjectLogs   string         `json:"project_log"`
	Reports       []CommonObject `json:"reports"`
}

type ExternalAPI struct {
	Name  string `json:"name"`
	Image string `json:"img"`
	Link  string `json:"link"`
}

type Technology struct {
	Name  string `json:"name"`
	Image string `json:"img"`
	Link  string `json:"link"`
}

type ProgressImage struct {
	Title  string   `json:"title"`
	Images []string `json:"imgs"`
}

type Image struct {
	Thumbnail     string          `json:"thumbnail"`
	ProgressImage []ProgressImage `json:"progress"`
}

type Project struct {
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	ProjectType  string        `json:"project_type"`
	Role         string        `json:"role"`
	Images       Image         `json:"images"`
	Technologies []Technology  `json:"technologies"`
	GithubLink   string        `json:"github"`
	LiveURL      CommonObject  `json:"live_url"`
	APIs         []ExternalAPI `json:"apis"`
	ExtraLinks   ExtraLink     `json:"extra_links"`
	Organisation Organisation  `json:"organisation"`
	Timeline     Timeline      `json:"timeline"`
}

type Portfolio struct {
	Apps   []Project `json:"apps"`
	Design []Project `json:"design"`
}

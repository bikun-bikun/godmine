package redmine

type projectResult struct {
	Project Project `json:"project"`
}

type projectsResult struct {
	Projects []*Project `json:"projects"`
}

type Project struct {
	ID              int      `json:"id"`
	Name            string   `json:"name"`
	Identifier      string   `json:"identifier"`
	Description     string   `json:"description"`
	Parent          IdName   `json:"parent"`
	Homepage        string   `json:"homepage"`
	Status          int      `json:"status"`
	IsPublic        bool     `json:"is_public"`
	Trackers        []IdName `json:"trackers"`
	IssueCategories []IdName `json:"issue_categories"`
	EnabledModules  []IdName `json:"enabled_modules"`
	CreatedOn       string   `json:"created_on"`
	UpdatedOn       string   `json:"updated_on"`
}

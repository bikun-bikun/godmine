package redmine

import "net/http"

type issueResult struct {
	Issue Issue `json:"issue" yaml:"issue"`
}

type issuesResult struct {
	Issues     []*Issue `json:"issues"`
	TotalCount int      `json:"total_count"`
	Offset     int      `json:"offset"`
	Limit      int      `json:"limit"`
}

type Issue struct {
	Id                  int         `yaml,json:"id"`
	Project             *IdName     `yaml,json:"project"`
	Tracker             *IdName     `yaml,json:"tracker"`
	Status              *IdName     `yaml,json:"status"`
	Priority            *IdName     `yaml,json:"priority"`
	Author              *IdName     `yaml,json:"author"`
	AssignedTo          *IdName     `yaml,json:"assigned_to"`
	Parent              *Id         `yaml,json:"parent"`
	Subject             string      `yaml,json:"subject"`
	Description         string      `yaml,json:"description"`
	StartDate           string      `yaml,json:"start_date"`
	DueDate             string      `yaml,json:"due_date"`
	DoneRatio           int         `yaml,json:"done_date"`
	IsPrivate           bool        `yaml,json:"is_private"`
	EstimatedHours      int         `yaml,json:"estimated_hours"`
	TotalEstimatedHours int         `yaml,json:"total_estimated_hours"`
	CreatedOn           string      `yaml,json:"created_on"`
	UpdatedOn           string      `yaml,json:"updated_on"`
	ClosedOn            string      `yaml,json:"closed_on"`
	Relations           []*Relation `yaml,json:"relations"`
	Journals            []*Journal  `yaml,json:"journals"`
}

type Relation struct {
	Id           int    `json:"id"`
	IssueId      int    `json:"issue_id"`
	IssueToId    int    `json:"issue_to_id"`
	RelationType string `json:"relation_type"`
	Delay        string `json:"delay"`
}

type Journal struct {
	Id           int              `json:"id"`
	User         IdName           `json:"user"`
	Notes        string           `json:"notes"`
	CreatedOn    string           `json:"created_on"`
	PrivateNotes bool             `json:"private_notes"`
	Details      []*JournalDetail `json:"details"`
}

type JournalDetail struct {
	Property string `json:"property"`
	Name     string `json:"name"`
	OldValue string `json:"old_value"`
	NewValue string `json:"new_value"`
}

func (c *Client) getIssues(projectID interface{}) ([]*Issue, error) {
	res, err := c.Get()
	url := c.endpoint + "/issues.json"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-Redmine-API-Key", c.apikey)
	req.Header.Add("Content-Type", "application/json")

	res, _ := c.Do(req)

	return nil, nil
}

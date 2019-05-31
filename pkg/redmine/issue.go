package redmine

import (
	"encoding/json"
	"net/http"
)

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
	Id                  int         `yaml:"id" json:"id"`
	Project             *IdName     `yaml:"project" json:"project"`
	Tracker             *IdName     `yaml:"tracker" json:"tracker"`
	Status              *IdName     `yaml:"status" json:"status"`
	Priority            *IdName     `yaml:"priority" json:"priority"`
	Author              *IdName     `yaml:"author" json:"author"`
	AssignedTo          *IdName     `json:"assigned_to"`
	AssignedMembers     []int       `yaml:"members"`
	Parent              *Id         `yaml:"parent" json:"parent"`
	Subject             string      `yaml:"subject" json:"subject"`
	Description         string      `yaml:"description" json:"description"`
	StartDate           string      `yaml:"start_date" json:"start_date"`
	DueDate             string      `yaml:"due_date" json:"due_date"`
	DoneRatio           int         `yaml:"done_date" json:"done_date"`
	IsPrivate           bool        `yaml:"is_private" json:"is_private"`
	EstimatedHours      float32     `yaml:"estimated_hours" json:"estimated_hours"`
	TotalEstimatedHours float32     `yaml:"total_estimated_hours" json:"total_estimated_hours"`
	CreatedOn           string      `yaml:"created_on" json:"created_on"`
	UpdatedOn           string      `yaml:"updated_on" json:"updated_on"`
	ClosedOn            string      `yaml:"closed_on" json:"closed_on"`
	Relations           []*Relation `yaml:"relations" json:"relations"`
	Journals            []*Journal  `yaml:"journals" json:"journals"`
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

func (c *Client) GetIssues(projectID interface{}) ([]*Issue, error) {
	url := c.endpoint + "/issues.json"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-Redmine-API-Key", c.apikey)
	req.Header.Add("Content-Type", "application/json")

	res, _ := c.Do(req)

	var ir issuesResult
	decoder := json.NewDecoder(res.Body)
	err := decoder.Decode(&ir)
	if err != nil {
		return nil, err
	}

	return ir.Issues, nil
}

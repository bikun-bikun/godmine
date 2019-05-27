package redmine

type issuesResult struct {
	Issues     []*Issue
	TotalCount int
	Offset     int
	Limit      int
}

type Issue struct {
	Id                  int         `json:"id"`
	Project             *IdName     `json:"project"`
	Tracker             *IdName     `json:"tracker"`
	Status              *IdName     `json:"status"`
	Priority            *IdName     `json:"priority"`
	Author              *IdName     `json:"author"`
	AssignedTo          *IdName     `json:"assigned_to"`
	Parent              *Id         `json:"parent"`
	Subject             string      `json:"subject"`
	Description         string      `json:"description"`
	StartDate           string      `json:"start_date"`
	DueDate             string      `json:"due_date"`
	DoneRatio           int         `json:"done_date"`
	IsPrivate           bool        `json:"is_private"`
	EstimatedHours      int         `json:"estimated_hours"`
	TotalEstimatedHours int         `json:"total_estimated_hours"`
	CreatedOn           string      `json:"created_on"`
	UpdatedOn           string      `json:"updated_on"`
	ClosedOn            string      `json:"closed_on"`
	Relations           []*Relation `json:"relations"`
	Journals            []*Journal  `json:"journals"`
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

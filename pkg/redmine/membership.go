package redmine

type membershipsResult struct {
	Memberships []Membership `json:"memberships"`
	TotalCount  int          `json:"total_count"`
	Limit       int          `json:"limit"`
	Offset      int          `json:"offset"`
}

type Membership struct {
	ID      int       `json:"id"`
	Project IdName    `json:"project"`
	User    IdName    `json:"user"`
	Roles   []*IdName `json:"roles"`
}

func (c *Client) getMemberships(project string) ([]*Membership, error) {
	return nil, nil
}

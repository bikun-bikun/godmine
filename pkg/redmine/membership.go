package redmine

import (
	"encoding/json"
	"net/http"
)

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

func (c *Client) GetMemberships(project string) ([]Membership, error) {
	url := c.endpoint + "/memberships.json"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-Redmine-API-Key", c.apikey)
	req.Header.Add("Content-Type", "application/json")

	res, _ := c.Do(req)
	var mr membershipsResult
	decoder := json.NewDecoder(res.Body)
	defer res.Body.Close()
	err := decoder.Decode(&mr)
	if err != nil {
		return nil, err
	}

	return mr.Memberships, nil
}

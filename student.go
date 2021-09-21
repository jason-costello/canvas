package canvas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) GetStudents() ([]Student, error) {
students, err := getStudents(c)
if err != nil {
return []Student{}, nil
}

var allStudents []Student
for _, s := range students {
allStudents = append(allStudents, Student{
ID:          s.ID,
Name:        s.Name,
})
}
return allStudents, nil
}


type Students []Student

func UnmarshalStudents(data []byte) (Students, error) {
	var r Students
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Students) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Student struct {
	ID                            int    `json:"id"`
	Name                          string `json:"name"`
	CreatedAt                     string `json:"created_at"`
	SortableName                  string `json:"sortable_name"`
	ShortName                     string `json:"short_name"`
	ObservationLinkRootAccountIDS []int  `json:"observation_link_root_account_ids"`
}

func getStudents(c *Client) (Students, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/users/%s/observees", c.baseURL, c.observerID), nil)
	if err != nil {
		return Students{}, err
	}


	req.Header.Add("Authorization", c.authToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Students{}, err
	}

	defer resp.Body.Close()

	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Students{}, err
	}

	return UnmarshalStudents(d)
}


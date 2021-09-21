package canvas

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) GetGrades(studentID, courseID string) error {
	path := fmt.Sprintf("%s/api/v1/courses/%s/users/%s?include[]=enrollments&", c.baseURL, courseID, studentID)
	// path := "/api/v1/users/:user_id/courses/:course_id/assignments"
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", c.authToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(d))

	return nil

}



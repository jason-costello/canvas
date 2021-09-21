package canvas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


func (c *Client) GetSubmissions(studentID, courseID, assignmentID string) (Submission, error) {

	// /api/v1/courses/:course_id/students/submissions


	path := fmt.Sprintf("%s/api/v1/courses/%s/assignments/%s/submissions/%s", c.baseURL, courseID,assignmentID,studentID)

	// path := fmt.Sprintf("%s/api/v1/courses/%s/assignments/%s/submissions", c.baseURL, courseID, assignmentID)
	//path := fmt.Sprintf("%s/api/v1/courses/%s/assignments/%s/submissions", c.baseURL, courseID, assignmentID)
	// path := "/api/v1/users/:user_id/courses/:course_id/assignments"
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return Submission{}, err
	}

	req.Header.Add("Authorization", c.authToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Submission{}, err
	}

	defer resp.Body.Close()

	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Submission{}, err
	}


 	return UnmarshalSubmission(d)


}

func UnmarshalSubmission(data []byte) (Submission, error) {
	var r Submission
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Submission) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Submission struct {
	AssignmentID                  int64       `json:"assignment_id"`
	ID                            int64       `json:"id"`
	UserID                        int64       `json:"user_id"`
	Body                          interface{} `json:"body"`
	URL                           string      `json:"url"`
	Grade                         string      `json:"grade"`
	Score                         float64       `json:"score"`
	SubmittedAt                   string      `json:"submitted_at"`
	SubmissionType                string      `json:"submission_type"`
	WorkflowState                 string      `json:"workflow_state"`
	GradeMatchesCurrentSubmission bool        `json:"grade_matches_current_submission"`
	GradedAt                      string      `json:"graded_at"`
	GraderID                      int64       `json:"grader_id"`
	Attempt                       int64       `json:"attempt"`
	CachedDueDate                 string      `json:"cached_due_date"`
	Excused                       interface{} `json:"excused"`
	LatePolicyStatus              interface{} `json:"late_policy_status"`
	PointsDeducted                interface{} `json:"points_deducted"`
	GradingPeriodID               int64       `json:"grading_period_id"`
	ExtraAttempts                 interface{} `json:"extra_attempts"`
	PostedAt                      string      `json:"posted_at"`
	Late                          bool        `json:"late"`
	Missing                       bool        `json:"missing"`
	SecondsLate                   int64       `json:"seconds_late"`
	EnteredGrade                  string      `json:"entered_grade"`
	EnteredScore                  float64       `json:"entered_score"`
	PreviewURL                    string      `json:"preview_url"`
	ExternalToolURL               string      `json:"external_tool_url"`
}

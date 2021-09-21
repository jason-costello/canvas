package canvas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func (c *Client) GetAssignments(studentID, courseID int) ([]Assignment, error){
	return getAssignments(c, studentID, courseID)

}


func getAssignments(c *Client, studentID, courseID int) (Assignments, error) {

	path := fmt.Sprintf("%s/api/v1/users/%v/courses/%v/assignments", c.baseURL, studentID, courseID)
	// path := "/api/v1/users/:user_id/courses/:course_id/assignments"
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return Assignments{}, err
	}

	req.Header.Add("Authorization", c.authToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Assignments{}, err
	}

	defer resp.Body.Close()

	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Assignments{}, err
	}

	return UnmarshalAssignments(d)

}


type Assignments []Assignment


func UnmarshalAssignments(data []byte) (Assignments, error) {
	var r Assignments
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Assignments) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Assignment struct {
	ID                              int         `json:"id"`
	Description                     string      `json:"description"`
	DueAt                           time.Time   `json:"due_at"`
	UnlockAt                        interface{} `json:"unlock_at"`
	LockAt                          time.Time   `json:"lock_at"`
	PointsPossible                  float64     `json:"points_possible"`
	GradingType                     string      `json:"grading_type"`
	AssignmentGroupID               int         `json:"assignment_group_id"`
	GradingStandardID               interface{} `json:"grading_standard_id"`
	CreatedAt                       time.Time   `json:"created_at"`
	UpdatedAt                       time.Time   `json:"updated_at"`
	PeerReviews                     bool        `json:"peer_reviews"`
	AutomaticPeerReviews            bool        `json:"automatic_peer_reviews"`
	Position                        int         `json:"position"`
	GradeGroupStudentsIndividually  bool        `json:"grade_group_students_individually"`
	AnonymousPeerReviews            bool        `json:"anonymous_peer_reviews"`
	GroupCategoryID                 interface{} `json:"group_category_id"`
	PostToSis                       bool        `json:"post_to_sis"`
	ModeratedGrading                bool        `json:"moderated_grading"`
	OmitFromFinalGrade              bool        `json:"omit_from_final_grade"`
	IntraGroupPeerReviews           bool        `json:"intra_group_peer_reviews"`
	AnonymousInstructorAnnotations  bool        `json:"anonymous_instructor_annotations"`
	AnonymousGrading                bool        `json:"anonymous_grading"`
	GradersAnonymousToGraders       bool        `json:"graders_anonymous_to_graders"`
	GraderCount                     int         `json:"grader_count"`
	GraderCommentsVisibleToGraders  bool        `json:"grader_comments_visible_to_graders"`
	FinalGraderID                   interface{} `json:"final_grader_id"`
	GraderNamesVisibleToFinalGrader bool        `json:"grader_names_visible_to_final_grader"`
	AllowedAttempts                 int         `json:"allowed_attempts"`
	SecureParams                    string      `json:"secure_params"`
	CourseID                        int         `json:"course_id"`
	Name                            string      `json:"name"`
	SubmissionTypes                 []string    `json:"submission_types"`
	HasSubmittedSubmissions         bool        `json:"has_submitted_submissions"`
	DueDateRequired                 bool        `json:"due_date_required"`
	MaxNameLength                   int         `json:"max_name_length"`
	InClosedGradingPeriod           bool        `json:"in_closed_grading_period"`
	IsQuizAssignment                bool        `json:"is_quiz_assignment"`
	CanDuplicate                    bool        `json:"can_duplicate"`
	OriginalCourseID                int         `json:"original_course_id"`
	OriginalAssignmentID            int         `json:"original_assignment_id"`
	OriginalAssignmentName          string      `json:"original_assignment_name"`
	OriginalQuizID                  interface{} `json:"original_quiz_id"`
	WorkflowState                   string      `json:"workflow_state"`
	ImportantDates                  bool        `json:"important_dates"`
	IsQuizLtiAssignment             bool        `json:"is_quiz_lti_assignment,omitempty"`
	FrozenAttributes                []string    `json:"frozen_attributes,omitempty"`
	ExternalToolTagAttributes       struct {
		URL            string      `json:"url"`
		NewTab         bool        `json:"new_tab"`
		ResourceLinkID string      `json:"resource_link_id"`
		ExternalData   string      `json:"external_data"`
		ContentType    string      `json:"content_type"`
		ContentID      int         `json:"content_id"`
		CustomParams   interface{} `json:"custom_params"`
	} `json:"external_tool_tag_attributes,omitempty"`
	Muted                  bool   `json:"muted"`
	HTMLURL                string `json:"html_url"`
	URL                    string `json:"url,omitempty"`
	Published              bool   `json:"published"`
	OnlyVisibleToOverrides bool   `json:"only_visible_to_overrides"`
	LockedForUser          bool   `json:"locked_for_user"`
	SubmissionsDownloadURL string `json:"submissions_download_url"`
	PostManually           bool   `json:"post_manually"`
	AnonymizeStudents      bool   `json:"anonymize_students"`
	RequireLockdownBrowser bool   `json:"require_lockdown_browser"`
	LockInfo               struct {
		LockAt      time.Time `json:"lock_at"`
		CanView     bool      `json:"can_view"`
		AssetString string    `json:"asset_string"`
	} `json:"lock_info,omitempty"`
	AllowedExtensions []string `json:"allowed_extensions,omitempty"`
	LockExplanation   string   `json:"lock_explanation,omitempty"`
}



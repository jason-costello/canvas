package canvas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) GetCourses(studentID string) (Courses, error) {

	// path := "/api/v1/users/:user_id/courses/:course_id/assignments"
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/users/%s/courses?enrollment_state=active&include[]=total_scores", c.baseURL, studentID), nil)
	if err != nil {
		return Courses{}, err
	}

	req.Header.Add("Authorization", c.authToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Courses{}, err
	}

	defer resp.Body.Close()

	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Courses{}, err
	}


	return UnmarshalCourses(d)

}

type Courses []Course

func UnmarshalCourses(data []byte) (Courses, error) {
	var r Courses
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Courses) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Course struct {
	ID                               int          `json:"id"`
	RootAccountID                    *int         `json:"root_account_id,omitempty"`
	AccountID                        *int         `json:"account_id,omitempty"`
	Name                             *string      `json:"name,omitempty"`
	EnrollmentTermID                 *int         `json:"enrollment_term_id,omitempty"`
	UUID                             *string      `json:"uuid,omitempty"`
	StartAt                          *string      `json:"start_at,omitempty"`
	GradingStandardID                interface{}  `json:"grading_standard_id"`
	IsPublic                         *bool        `json:"is_public"`
	CreatedAt                        *string      `json:"created_at,omitempty"`
	CourseCode                       *string      `json:"course_code,omitempty"`
	DefaultView                      *string      `json:"default_view,omitempty"`
	License                          *string      `json:"license"`
	GradePassbackSetting             interface{}  `json:"grade_passback_setting"`
	EndAt                            *string      `json:"end_at"`
	PublicSyllabus                   *bool        `json:"public_syllabus,omitempty"`
	PublicSyllabusToAuth             *bool        `json:"public_syllabus_to_auth,omitempty"`
	StorageQuotaMB                   *int64       `json:"storage_quota_mb,omitempty"`
	IsPublicToAuthUsers              *bool        `json:"is_public_to_auth_users,omitempty"`
	HomeroomCourse                   *bool        `json:"homeroom_course,omitempty"`
	CourseColor                      interface{}  `json:"course_color"`
	FriendlyName                     interface{}  `json:"friendly_name"`
	ApplyAssignmentGroupWeights      *bool        `json:"apply_assignment_group_weights,omitempty"`
	Calendar                         *Calendar    `json:"calendar,omitempty"`
	TimeZone                         *string      `json:"time_zone,omitempty"`
	Blueprint                        *bool        `json:"blueprint,omitempty"`
	Template                         *bool        `json:"template,omitempty"`
	Enrollments                      []Enrollment `json:"enrollments,omitempty"`
	HideFinalGrades                  *bool        `json:"hide_final_grades,omitempty"`
	WorkflowState                    *string      `json:"workflow_state,omitempty"`
	RestrictEnrollmentsToCourseDates *bool        `json:"restrict_enrollments_to_course_dates,omitempty"`
	OverriddenCourseVisibility       *string      `json:"overridden_course_visibility,omitempty"`
	AccessRestrictedByDate           *bool        `json:"access_restricted_by_date,omitempty"`
}

type Calendar struct {
	ICS string `json:"ics"`
}

type Enrollment struct {
	Type                           string `json:"type"`
	Role                           string `json:"role"`
	RoleID                         int    `json:"role_id"`
	UserID                         int    `json:"user_id"`
	EnrollmentState                string `json:"enrollment_state"`
	LimitPrivilegesToCourseSection bool   `json:"limit_privileges_to_course_section"`
}



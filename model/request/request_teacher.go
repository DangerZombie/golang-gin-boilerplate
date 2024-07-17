package request

// swagger:parameters TeacherCreateRequest
type TeacherCreateRequest struct {
	// in:body
	Body TeacherCreateRequestBody `json:"body"`
}

type TeacherCreateRequestBody struct {
	// User Id of the teacher
	// in: string
	// required: true
	UserId string `json:"user_id"`

	// Job Title Id of the teacher
	// in: string
	// required: true
	JobTitleId string `json:"job_title_id"`

	// Status of the teacher
	// in: string
	// required: true
	Status string `json:"status"`

	// Long time of experience in years
	// in: int
	Experience int `json:"experience"`

	// Degree of the teacher
	// in: string
	Degree string `json:"degree"`

	Issuer string `json:"-"`
}

// swagger:parameters TeacherListRequest
type TeacherListRequest struct {
	// Page of the list
	// in: int
	Page int `json:"page"`

	// Limit row of each page
	// in: int
	Limit int `json:"limit"`

	// Sorting by column
	// in: string
	Sort string `json:"sort"`

	// Direction of sorting
	// in: string
	Dir string `json:"dir"`

	// Filtering name
	// in: string
	Name string `json:"name"`
}

// swagger:parameters TeacherDetailRequest
type TeacherDetailRequest struct {
	// Id of the teacher
	// in: path
	Id string `json:"id"`
}

// swagger:parameters TeacherUpdateRequest
type TeacherUpdateRequest struct {
	// Id of the teacher
	// in: path
	Id string `json:"id"`

	// in:body
	Body TeacherUpdateRequestBody `json:"body"`
}

type TeacherUpdateRequestBody struct {
	// Job Title Id
	// in: string
	JobTitleId *string `json:"job_title_id,omitempty"`

	// Status of the teacher
	// in: string
	Status *string `json:"status,omitempty"`

	// Experience of the teacher
	// in: int
	Experience *int `json:"experience,omitempty"`

	// Degree of the teacher
	// in: string
	Degree *string `json:"degree,omitempty"`
}

// swagger:parameters TeacherDeleteRequest
type TeacherDeleteRequest struct {
	// Id of teacher
	// in: path
	Id string `json:"id"`
}

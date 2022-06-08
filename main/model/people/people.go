package people

type People struct {
	PersonId       int    `json:"personId" db:"person_id"`
	EmployeeNumber string `json:"employeeNumber" db:"employee_number"`
	FirstName      string `json:"firstName" db:"first_name"`
	LastName       string `json:"lastName" db:"last_name"`
	Gender         string `json:"gender" db:"gender"`
	MaritalStatus  string `json:"maritalStatus" db:"marital_status"`
	DateOfBirth    string `json:"dateOfBirth" db:"date_of_birth"`
	JoiningDate    string `json:"joiningDate" db:"joining_date"`
	JobId          int    `json:"jobId" db:"job_id"`
	DepartmentId   int    `json:"departmentId" db:"department_id"`
	LocationId     int    `json:"locationId" db:"location_id"`
}

type Repo interface {
	FindAllPeople() ([]People, error)
	FindOnePerson(personId int) (*People, error)
}

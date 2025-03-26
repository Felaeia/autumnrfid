package repositories

import (
	"database/sql"
	"fmt"
	"time"
)

type Student struct {
	Student_Number      string
	DepartmentID        int
	FirstName           string
	LastName            string
	MiddleName          string
	YearLevel           int
	Program             string
	Birthday            string
	Contact_Number      int64
	Email               string
	Block               string
	System_First_Access string
	System_Last_Access  string
}

// Bills Data Structure
type Bills struct {
	assessment_Number       int //Assessment Table PK
	StudentNumber           string
	PaymentNumber           int //Payment Table PK
	TuitionAmt              float64
	MiscFeesName            string
	MiscFeesAmt             float64
	DiscountAmt             float64
	InitialPayment          float64
	RemainingBalance        float64
	FullPaymentBeforePrelim float64
	PerExamFee              float64
	PaymentStatus           string
	Date                    time.Time
	Description             string
	Amount                  float64
	Status                  string
}

// Grades Data Structure
type Grades struct {
	Grade_ID            int
	StudentNumber       string
	AcademicYear        string
	FirstSemesterGrade  float64
	SecondSemesterGrade float64
}

type RFIDRepository struct {
	dbClient *DatabaseClient
}

func NewRFIDRepository(dbClient *DatabaseClient) *RFIDRepository {
	return &RFIDRepository{dbClient: dbClient}
}

// Query all students data in the db
func (r *RFIDRepository) GetStudentByRFID(student_Number string) (*Student, error) {
	query := `
		SELECT student_Number, department_ID, first_Name, last_Name, middle_Name, year_Level, program, birthday, contact_Number, email, block, system_First_Access, system_Last_Access
		FROM Students
		WHERE student_Number = ?
	`

	fmt.Printf("Executing query with student ID: %s\n", student_Number)

	student := &Student{}
	err := r.dbClient.DB.QueryRow(query, student_Number).Scan(
		&student.Student_Number,
		&student.DepartmentID,
		&student.FirstName,
		&student.LastName,
		&student.MiddleName,
		&student.YearLevel,
		&student.Program,
		&student.Birthday,
		&student.Contact_Number,
		&student.Email,
		&student.Block,
		&student.System_First_Access,
		&student.System_Last_Access,
	)

	if err == sql.ErrNoRows {
		fmt.Printf("No student found with ID: %s\n", student_Number)
		return nil, nil
	}
	if err != nil {
		fmt.Printf("Error querying student: %v\n", err)
		return nil, fmt.Errorf("error querying student: %v", err)
	}

	return student, nil
}

func (r *RFIDRepository) GetStudentBillsByRFID(student_Number string) (*Bills, error) {
	query := `
		SELECT assessment_Number, student_Number, payment_Number, tuition_Amt, misc_Fees_Name, misc_Fees_Amt, 
		discount_Amt, initial_Payment, remaining_Balance, full_pmt_if_b4_prelim, per_Exam_Fee, payment_Status, payment_Number,
		date, description, amount, status
		FROM Bills
		WHERE student_Number = ?
	`

	fmt.Printf("Executing query with student ID: %s\n", student_Number)

	student := &Student{}
	err := r.dbClient.DB.QueryRow(query, student_Number).Scan(
		&student.Student_Number,
		&student.DepartmentID,
		&student.FirstName,
		&student.LastName,
		&student.MiddleName,
		&student.YearLevel,
		&student.Program,
		&student.Birthday,
		&student.Contact_Number,
		&student.Email,
		&student.Block,
		&student.System_First_Access,
		&student.System_Last_Access,
	)

	if err == sql.ErrNoRows {
		fmt.Printf("No student found with ID: %s\n", student_Number)
		return nil, nil
	}
	if err != nil {
		fmt.Printf("Error querying student: %v\n", err)
		return nil, fmt.Errorf("error querying student: %v", err)
	}

	return student, nil
}

func (r *RFIDRepository) GetStudentGradesByRFID(student_Number string) (*Grades, error) {
	return nil, nil
}

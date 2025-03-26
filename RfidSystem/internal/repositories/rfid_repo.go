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
	AssessmentNumber        int //Assessment Table PK
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
	PaymentDate             time.Time
	PaymentDescription      string
	PaymentAmount           float64
	PaymentHistoryStatus    string
}

// Grades Data Structure
type Grades struct {
	GradeID             int
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
		FROM students
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

// Query all student bills data in the db
func (r *RFIDRepository) GetStudentBillsByRFID(student_Number string) (*Bills, error) {
	query := `
	SELECT 
		assess.assessment_Number,
		assess.student_Number,
		assess.payment_Number,
		assess.tuition_Amt,
		assess.misc_Fees_Name,
		assess.misc_Fees_Amt,
		assess.discount_Amt,
		assess.initial_Payment,
		assess.remaining_Balance,
		assess.full_pmt_if_b4_prelim,
		assess.per_Exam_Fee,
		assess.payment_Status,
		payHist.date,
		payHist.description,
		payHist.amount,
		payHist.status
	FROM assessment AS assess
	LEFT JOIN payment_history AS payHist ON assess.payment_Number = payHist.payment_Number
	WHERE assess.student_Number = ?

	`

	fmt.Printf("Executing query with student ID: %s\n", student_Number)

	bill := &Bills{}
	err := r.dbClient.DB.QueryRow(query, student_Number).Scan(
		//Assessment Table
		&bill.AssessmentNumber,
		&bill.StudentNumber,
		&bill.PaymentNumber,
		&bill.TuitionAmt,
		&bill.MiscFeesName,
		&bill.MiscFeesAmt,
		&bill.DiscountAmt,
		&bill.InitialPayment,
		&bill.RemainingBalance,
		&bill.FullPaymentBeforePrelim,
		&bill.PerExamFee,
		&bill.PaymentStatus,
		//Payment History Table
		&bill.PaymentDate,
		&bill.PaymentDescription,
		&bill.PaymentAmount,
		&bill.PaymentHistoryStatus,
	)

	if err == sql.ErrNoRows {
		fmt.Printf("No student found with ID: %s\n", student_Number)
		return nil, nil
	}
	if err != nil {
		fmt.Printf("Error querying student: %v\n", err)
		return nil, fmt.Errorf("error querying student: %v", err)
	}

	return bill, nil
}

func (r *RFIDRepository) GetStudentGradesByRFID(student_Number string) (*Grades, error) {
	query := `
		SELECT 
			grade_ID, 
			student_Number, 
			academic_Year,
			first_Semester_Grade,
			second_Semester_Grade
		FROM grades
		WHERE student_Number = ?
	`

	fmt.Printf("Executing query with student ID: %s\n", student_Number)

	grades := &Grades{}
	err := r.dbClient.DB.QueryRow(query, student_Number).Scan(
		&grades.GradeID,
		&grades.StudentNumber,
		&grades.AcademicYear,
		&grades.FirstSemesterGrade,
		&grades.SecondSemesterGrade,
	)

	if err == sql.ErrNoRows {
		fmt.Printf("No student found with ID: %s\n", student_Number)
		return nil, nil
	}
	if err != nil {
		fmt.Printf("Error querying student: %v\n", err)
		return nil, fmt.Errorf("error querying student: %v", err)
	}

	return grades, nil
}

package handlers

// import (
// 	"html/template"
// 	"net/http"
// )

// func (h *Handler) ViewStudentPage(w http.ResponseWriter, r *http.Request) {
// 	studentID := r.URL.Query().Get("id") // e.g., /student?id=2023-0001

// 	student, err := h.repo.GetStudentByRFID(studentID)
// 	if err != nil {
// 		http.Error(w, "Student not found", http.StatusNotFound)
// 		return
// 	}

// 	data := struct {
// 		Student *Student
// 	}{
// 		Student: student,
// 	}

// 	tmpl := template.Must(template.ParseFiles(
// 		"ui/html/pages/student_page.html",           // main layout
// 		"ui/html/partials/student_information.html", // the fragment
// 	))

// 	err = tmpl.ExecuteTemplate(w, "student_page", data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

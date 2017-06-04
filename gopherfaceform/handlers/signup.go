package handlers

import (
	"net/http"
)

type SignUpForm struct {
	FieldNames []string
	Fields     map[string]string
	Errors     map[string]string
}

// DisplaySignUpForm displays the Sign Up form
func DisplaySignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm) {
	RenderTemplate(w, "./templates/signup.html", s)
}

func PopulateFormFields(r *http.Request, s *SignUpForm) {

	for _, fieldName := range s.FieldNames {
		s.Fields[fieldName] = r.FormValue(fieldName)
	}

}

// ValidateSignUpForm validates the Sign Up form's fields
func ValidateSignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm) {

	PopulateFormFields(r, s)
	// Check if username was filled out

	// Check if first name was filled out

	// Check if last name was filled out

	// Check if e-mail address was filled out

	// Check username syntax

	// Check e-mail address syntax

	// Check if passord and confirm password field values match

}

// ProcessSignUpForm
func ProcessSignUpForm(w http.ResponseWriter, r *http.Request) {

}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {

	s := SignUpForm{}
	s.FieldNames = []string{"username", "firstName", "lastName", "email"}
	s.Fields = make(map[string]string)
	s.Errors = make(map[string]string)
	//	s.Errors[""] = "test"

	switch r.Method {

	case "GET":
		DisplaySignUpForm(w, r, &s)
	case "POST":
		ValidateSignUpForm(w, r)
	default:
		DisplaySignUpForm(w, r, &s)
	}

}

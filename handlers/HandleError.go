package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// var errorTemplate *template.Template

// func init() {
// 	var err error
// 	errorTemplate, err = template.ParseFiles("./templates/error.html")
// 	if err != nil {
// 		log.Fatalf("failed to parse error template: %v", err)
// 	}
// }

// // ErrorPage renders the error.html template with the provided error message.
//
//	func HandleError(w http.ResponseWriter, message string, code int) {
//		w.WriteHeader(code) // Set the appropriate HTTP status code
//		if execErr := errorTemplate.Execute(w, map[string]string{
//			"ErrorMessage": message,
//		}); execErr != nil {
//			log.Printf("failed to execute error template: %v", execErr)
//			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//		}
//	}
func responseAlreadyWritten(w http.ResponseWriter) bool {
	return w.Header().Get("Content-Type") != ""
}

func HandleError(w http.ResponseWriter, err error, statusCode int, templateName string) {
	log.Println(err)

	if !responseAlreadyWritten(w) {
		w.WriteHeader(statusCode)
	}
	tpl, tplErr := template.ParseFiles("templates/" + templateName)
	if tplErr != nil {
		log.Println("Error parsing error template file:", tplErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if execErr := tpl.Execute(w, nil); execErr != nil {
		log.Println("Error executing error template:", execErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

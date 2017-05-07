// Simple example of creating and using a template in Go
package main

import (
	"fmt"
	"gocodelab/socialmedia"
	"html/template"
	"net/http"
)

// Handler for displaying a social media post
func displaySocialMediaPostHandler(w http.ResponseWriter, r *http.Request) {

	myPost := socialmedia.NewPost("EngineerKamesh", socialmedia.Moods["thrilled"], "Go is awesome!", "Check out the Go web site!", "https://golang.org", "/images/gogopher.png", "", []string{"go", "golang", "programming language"})
	fmt.Printf("myPost: %+v\n", myPost)
	renderTemplate(w, "./templates/socialmediapost.html", myPost)
}

// Template rendering function
func renderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {

	t, _ := template.ParseFiles(templateFile)
	t.Execute(w, templateData.(*socialmedia.Post))
}

func main() {

	http.HandleFunc("/display-social-media-post", displaySocialMediaPostHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
}

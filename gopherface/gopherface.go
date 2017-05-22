package main

import (
	//"gocodelab/gopherface/endpoints"
	"gocodelab/gopherface/handlers"
	//"gocodelab/gopherface/middleware"
	"net/http"

	//ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	WEBSERVERPORT = ":8080"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("GET,POST")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")
	r.HandleFunc("/feed", handlers.FeedHandler).Methods("GET")
	r.HandleFunc("/friends", handlers.FriendsHandler).Methods("GET")
	r.HandleFunc("/find", handlers.FindHandler).Methods("GET,POST")
	r.HandleFunc("/profile", handlers.MyProfileHandler).Methods("GET")
	r.HandleFunc("/profile/{username}", handlers.ProfileHandler)
	r.HandleFunc("/triggerpanic", handlers.TriggerPanicHandler)
	r.HandleFunc("/foo", handlers.FooHandler)

	//r.HandleFunc("/restapi/socialmediapost/{username}", endpoints.FetchPostsEndpoint).Methods("GET")
	//r.HandleFunc("/restapi/socialmediapost/{postid}", endpoints.CreatePostEndpoint).Methods("POST")
	//r.HandleFunc("/restapi/socialmediapost/{postid}", endpoints.UpdatePostEndpoint).Methods("PUT")
	//r.HandleFunc("/restapi/socialmediapost/{postid}", endpoints.DeletePostEndpoint).Methods("DELETE")

	http.Handle("/", r)
	//http.Handle("/", ghandlers.LoggingHandler(os.Stdout, r))
	//http.Handle("/", middleware.PanicRecoveryHandler(ghandlers.LoggingHandler(os.Stdout, r)))
	//http.Handle("/", middleware.ContextExampleHandler(middleware.PanicRecoveryHandler(ghandlers.LoggingHandler(os.Stdout, r))))

	http.ListenAndServe(WEBSERVERPORT, nil)

}

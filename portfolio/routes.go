package main

import "net/http"

func registerRoutes() {
	http.HandleFunc("/profile", profileHandler)
	http.HandleFunc("/skills", skillsHandler)
	http.HandleFunc("/socials", socialsHandler)
	http.HandleFunc("/summary", summaryHandler)
	http.HandleFunc("/", homeHandler)
}
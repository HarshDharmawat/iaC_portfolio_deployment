package main

import (
	"encoding/json"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	w.Write([]byte(`
	====================================
		 Harsh's API Server
	====================================

	Thank you for visiting! :-)

	Available endpoints:
	- /profile
	- /skills
	- /socials
	- /summary

	Built with Go :-)
	====================================
	`))
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(myIntro)
}

func skillsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"skills": myIntro.Skills,
	})
}

func socialsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(myIntro.Socials)
}

func summaryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"summary": myIntro.Description,
		"role":    myIntro.Role,
	})
}

package main

import (
	"fmt"
	"github.com/dlorch/errors.fail/config"
	"github.com/dlorch/errors.fail/session"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	probe := session.ReadBoolOrDefault("http_probe", true)

	if probe {
		fmt.Fprintf(w, "<h1>200 OK</h1>\n")
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "<h1>500 Internal Server Error</h1>\n")
	}

	fmt.Fprintf(w, "<p>Change your settings at <a href=\"https://%s/?%s\">errors.fail</a>.</p>", config.CookieDomain, session.SessionID)
}

func main() {
	http.HandleFunc("/", session.WithSession(handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

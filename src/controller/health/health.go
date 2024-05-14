package health

import (
	"encoding/json"
	"fmt"
	"html/template"
	"httpServer/src/middlewares"
	"httpServer/src/middlewares/logging"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the health endpoint\n")
}

func Traffic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(logging.Logs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln("There was an error encoding the initialized struct")
	}
}

func ShowHtml(w http.ResponseWriter, file string, data interface{}) {
	tmpl := template.Must(template.ParseFiles(file))
	tmpl.Execute(w, data)
}

func healthHtml(w http.ResponseWriter, r *http.Request) {
	ShowHtml(w, "src/controllers/health/health.html", nil)
}

func trafficHtml(w http.ResponseWriter, r *http.Request) {
	ShowHtml(w, "src/controllers/health/traffic.html", logging.Logs)
}

func InitController(r *mux.Router) {
	logsRouter := r.PathPrefix("/health").Subrouter()
	logsRouter.HandleFunc("", middlewares.Chain(Health)).Methods("GET")
	logsRouter.HandleFunc("/html", middlewares.Chain(healthHtml)).Methods("GET")
	logsRouter.HandleFunc("/traffic", middlewares.Chain(Traffic)).Methods("GET")
	logsRouter.HandleFunc("/traffic/html", middlewares.Chain(trafficHtml)).Methods("GET")
}

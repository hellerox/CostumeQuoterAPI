package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "carlos"
	DB_PASSWORD = "carlos"
	DB_NAME     = "carlos"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func main() {
	a := App{}
	var err error
	dbinfo := "postgres://" + DB_USER + ":" + DB_PASSWORD + "@192.168.2.19/" + DB_NAME
	a.DB, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/material", a.queryMaterial).Methods("GET")
	a.Router.HandleFunc("/material", a.createMaterial).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", a.Router))

}

func (a *App) queryMaterial(w http.ResponseWriter, r *http.Request) {

	ms, err := getMaterials(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, ms)
}

func (a *App) createMaterial(w http.ResponseWriter, r *http.Request) {
	var m Material
	decoder := json.NewDecoder(r.Body)
	fmt.Println("Recibí el body:", r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	if err := m.createMaterial(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, m)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	fmt.Println(payload)
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

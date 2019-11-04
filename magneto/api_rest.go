package main

import (
	"net/http"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var host = "http://localhost/"
var port = "port"
var connectionString = "user:password@tcp(127.0.0.1:3306)/magneto?charset=utf8&parseTime=True&loc=Local"

func main() {
	var router *mux.Router
	router = mux.NewRouter().StrictSlash(true)

	router.Handle("/mutant", http.HandlerFunc(mutant)).Methods("POST")
	router.Handle("/stats", http.HandlerFunc(stats)).Methods("GET")

	http.ListenAndServe(":"+port, router)
}

func mutant(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", connectionString)
	defer db.Close()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "DB error")
		return
	}
	
	decoder := json.NewDecoder(r.Body)
	var sequence sequence
	err = decoder.Decode(&sequence)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Sequence error")
		return
	}

	data, err := json.Marshal(sequence.DNA)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error decoding sequence")
		return
	}

	var id int
	var result bool
	db.QueryRow("SELECT id, is_mutant FROM sequence WHERE dna = ?;", data).Scan(&id, &result)
	if id != 0 {
		if result {
			sequence.RESULT = "Is mutant"
			respondWithJSON(w, http.StatusOK, sequence)
			return
		}
		sequence.RESULT = "Is human"
		respondWithJSON(w, http.StatusForbidden, sequence)
		return
	}
	statement, err := db.Prepare("INSERT INTO sequence (dna, is_mutant) VALUES(?,?);")

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Query error.")
		return
	}
	defer statement.Close()

	resultTest := isMutant(sequence.DNA[:])
	res, err := statement.Exec(data, resultTest)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Function error.")
		return
	}

	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 1 {
		if resultTest {
			sequence.RESULT = "Is mutant"
			respondWithJSON(w, http.StatusOK, sequence)
		} else {
			sequence.RESULT = "Is human"
			respondWithJSON(w, http.StatusForbidden, sequence)
		}
	}
}

func stats(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", connectionString)
	defer db.Close()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "DB error")
		return
	}
	var countMutantDna sql.NullInt32
	var countHumanDna sql.NullInt32
	var ratio sql.NullFloat64
	query := `
		SELECT 
		COUNT(CASE WHEN s.is_mutant = '1' THEN 1 ELSE NULL END) AS countMutantDna,
		COUNT(CASE WHEN s.is_mutant = '0' THEN 1 ELSE NULL END) AS countHumanDna,
		COUNT(CASE WHEN s.is_mutant = '1' THEN 1 ELSE NULL END)/COUNT(CASE WHEN s.is_mutant = '0' THEN 1 ELSE NULL END) AS ratio
		FROM sequence s`
	err = db.QueryRow(query).Scan(&countMutantDna, &countHumanDna, &ratio)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Query error")
		return
	}
	var stats stat
	stats.CountMutantDna = countMutantDna.Int32
	stats.CountHumanDna = countHumanDna.Int32
	stats.Ratio = ratio.Float64
	respondWithJSON(w, http.StatusOK, stats)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

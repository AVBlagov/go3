package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type solution struct {
	a       int `json: "a"`
	b       int `json: "b"`
	c       int `json: "c"`
	n_roots int `json: "n_roots"`
}

var sol []solution

func getLast(w http.ResponseWriter, r *http.Request) {
	s := sol[len(sol)-1]
	fmt.Println(s)
	json.NewEncoder(w).Encode(s)

}

func postValue(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	a, _ := strconv.Atoi(vars["a"])
	b, _ := strconv.Atoi(vars["b"])
	c, _ := strconv.Atoi(vars["c"])

	d := (b * b) - (4 * a * c)

	n := 0
	if d > 0 {
		n = 2
	} else if d == 0 {
		n = 1
	} else {
		n = 0
	}

	var s = solution{a, b, c, n}

	sol = append(sol, s)

	fmt.Println(sol)
	json.NewEncoder(w).Encode(sol)
}

func main() {
	route := mux.NewRouter().StrictSlash(true)
	route.HandleFunc("/solution", getLast).Methods("GET")
	route.HandleFunc("/solve/{a}/{b}/{c}", postValue).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", route))
}

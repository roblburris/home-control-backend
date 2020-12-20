package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// A struct is not necessary for storing on/off but once we get
// more complex stuff it will be useful
type ledState struct {
	state int
	test  string
}

var finState = ledState{state: 1000000, test: "abcdef"}

func getState(w http.ResponseWriter, req *http.Request) {
	// json.NewEncoder(w).Encode(finState)
	io.WriteString(w, strconv.Itoa(finState.state))
	log.Println(strconv.Itoa(finState.state))
}

// func setState(w http.ResponseWriter, req *http.Request) {
// 	finState.state = mux.Vars()
// 	log.Println(strconv.Itoa(finState.state))
// }

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/state", getState).Methods("GET")
	// router.HandleFunc("/set/{newState}", setState).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))

}

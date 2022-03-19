package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rahmaniali-ir/shorter/shorter"
)

type errorResponse struct {
	message string
}

func makeShorter(w http.ResponseWriter, req *http.Request) {

	params, ok := req.URL.Query()["link"]
    
	if !ok || len(params[0]) < 1 {
		fmt.Printf("\nInvalid request")
		fmt.Fprintf(w, "Invalid request!\n")
		return
	}

	short := shorter.MakeShorter(params[0])

	// log
	fmt.Printf("\n\"%s\" shortened to \"%s\"", params[0], short)

	fmt.Fprintf(w, short)
}

func makeLonger(w http.ResponseWriter, req *http.Request) {

	params, ok := req.URL.Query()["link"]
    
	if !ok || len(params[0]) < 1 {
		fmt.Printf("\nInvalid request")
		fmt.Fprintf(w, "Invalid request!\n")
		return
	}

	long, err := shorter.MakeLonger(params[0])

	if err != nil {
		fmt.Printf("\nFailed to longen \"%s\"!", params[0])
		
		var message = make(map[string]string)
		message["message"] = "No such link exists!"
		response, jsonErr := json.Marshal(message)
		
		if jsonErr != nil {
			fmt.Printf("\nFailed stringify json \"%v\"!", message)
			fmt.Fprintf(w, "Error!")
			} else {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(response))
		}

		return
	}

	// log
	fmt.Printf("\n\"%s\" longened to \"%s\"", params[0], long)
	fmt.Printf("\n%v", shorter.Records)

	fmt.Fprintf(w, long)
}

func HandleAPI() {
	http.HandleFunc("/shorten", makeShorter)
	http.HandleFunc("/longen", makeLonger)

	println("Listening on port 8090")
	http.ListenAndServe(":8090", nil)
}

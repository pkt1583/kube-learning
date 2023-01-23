// Access values from environment variables
// Path: main.go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/caarlos0/env"
	"github.com/gorilla/mux"
)

type config struct {
	DownstreamUrl string `env:"DOWNSTREAM_URL" envDefault:"https://localhost:8080"`
	PodName       string `env:"POD_NAME" envDefault:"myservice-local"`
}

func main() {
	var e config
	err := env.Parse(&e)
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "OK")

	}).Methods("GET")

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var result = make(map[string]interface{})
		result["podName"] = e.PodName
		fmt.Printf("accessing downstream url %s \n", e.DownstreamUrl)
		response, err := http.Get(e.DownstreamUrl)
		if err == nil {
			fmt.Printf("the response is %v  \n", response)
			all, er := io.ReadAll(response.Body)
			if er != nil {
				fmt.Printf("the error in reading response %s  \n", err)
			}
			fmt.Printf("the response body %s  \n", all)
			if err != nil {
				fmt.Printf("the error is %s  \n", err)
			}
			result["response"] = string(all)
		}
		result["error"] = err
		json.NewEncoder(w).Encode(result)

	})
	http.Handle("/", router)

	http.ListenAndServe(":8081", nil)

}

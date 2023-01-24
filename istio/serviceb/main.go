// Access values from environment variables
// Path: main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

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
		w.Write([]byte("OK from " + e.PodName))
	})

	router.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("accessing downstream url %s \n", e.DownstreamUrl)
		var result = make(map[string]interface{})
		result["podName"] = e.PodName
		result["message"] = "Hello World"
		json.NewEncoder(w).Encode(result)

	})
	http.Handle("/", router)

	http.ListenAndServe(":8081", nil)

}

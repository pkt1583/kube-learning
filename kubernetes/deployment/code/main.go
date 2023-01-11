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
	Message  string `env:"MESSAGE" envDefault:"No Message Set"`
	Userid   string `env:"USER_ID" envDefault:"No user id set"`
	Password string `env:"PASSWORD" envDefault:"No password set"`
	PodName  string `env:"POD_NAME" envDefault:"No pod name set"`
}

func main() {
	var e config
	err := env.Parse(&e)
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "OK")
	})

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var result = make(map[string]interface{})
		result["podName"] = e.PodName
		result["message"] = fmt.Sprintf("Hello %s", e.Message)
		result["userid"] = e.Userid
		result["password"] = e.Password
		result["error"] = err
		json.NewEncoder(w).Encode(result)

	})

	http.Handle("/", router)

	http.ListenAndServe(":8081", nil)

}

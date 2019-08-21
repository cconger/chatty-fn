package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var letters = []byte("abcdefghijklmnopqrstuvwxyz")

func randString(l int) string {
	b := make([]byte, l, l)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func noisy200(w http.ResponseWriter, r *http.Request) {
	lb := make(map[string]string)
	for i := 0; i < 100; i++ {
		lb[randString(5)] = randString(20)
	}

	for i := 0; i < 3; i++ {
		message, err := json.Marshal(lb)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		log.Print(string(message))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	log.Print("Starting application server")
	fmt.Fprintf(os.Stderr, "This is a warning to stderr\n")
	mux := http.NewServeMux()
	mux.HandleFunc("/", noisy200)

	s := &http.Server{
		Addr:           ":5000",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}

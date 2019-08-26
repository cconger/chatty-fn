package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type ReqLogEntry struct {
	Method        string
	URL           string
	Proto         string
	Host          string
	RemoteAddr    string
	ContentLength int64
	Body          string
	Headers       map[string][]string
	ReceivedAt    time.Time
}

// FromRequest returns a ReqLogEntry describing the assocated http.Request
func FromRequest(r *http.Request) *ReqLogEntry {
	entry := &ReqLogEntry{
		Method:        r.Method,
		URL:           r.URL.String(),
		Proto:         r.Proto,
		Host:          r.Host,
		RemoteAddr:    r.RemoteAddr,
		ContentLength: r.ContentLength,
		ReceivedAt:    time.Now(),
		Headers:       r.Header,
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		body = []byte{}
	}
	entry.Body = string(body)
	return entry
}

func noisy200(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request")
	message, err := json.Marshal(FromRequest(r))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	log.Println("Request:", string(message))

	// Purposely not a new line on this one
	fmt.Fprintf(os.Stdout, "This is me writing to stdout on request: ")
	fmt.Fprintln(os.Stdout, string(message))

	fmt.Println("I liked what I got... so I'm gonna 200 OK")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	// Do not include my own timestamps
	log.SetFlags(0)

	log.Print("Starting application server")
	fmt.Fprintf(os.Stdout, "This is a warning to stdout\n")
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

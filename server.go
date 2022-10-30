package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "encoding/json"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
	}
	log.Print(string(body))
}

func main() {

	var port int
	flag.IntVar(&port, "p", 8080, "Port")

	var url string
	flag.StringVar(&url, "u", "", "Url for accepting requests")

	var logFile string
	flag.StringVar(&logFile, "l", "", "log directory")

	flag.Parse()

	if url == "" {
		log.Panic("Url required, use the '-u' flag")
	}

	log.Printf("URL: %s\n", url)

	http.HandleFunc(url, webhookHandler)

	addr := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

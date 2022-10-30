package main

import (
	"crypto/rand"
	"encoding/base64"
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

	var err error

	var port int
	flag.IntVar(&port, "port", 8080, "Port")

	var path string
	flag.StringVar(&path, "path", "", "Path for accepting requests")

	var logFile string
	flag.StringVar(&logFile, "l", "", "log directory")

	flag.Parse()

	if path == "" {
		// if no custom path was provided - then generate a random one...
		randomBytes := make([]byte, 20)
		_, err = rand.Read(randomBytes)
		if err != nil {
			log.Panic(err)
		}
		path = fmt.Sprintf("/%s", base64.StdEncoding.EncodeToString(randomBytes))
	}

	fmt.Printf("URL: <host>:%d%s\n", port, path)

	http.HandleFunc(path, webhookHandler)

	addr := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

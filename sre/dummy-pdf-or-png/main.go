package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	s := fmt.Sprintf("Health OK!")
	io.WriteString(w, s)
	log.Println("OK")
}

func serveRandomFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println("ERROR: Expected a GET request")
		http.Error(w, "Not supported", 403)
		return
	}

	err := error(nil)
	if r.URL.String() != "/" {
		_, err = strconv.Atoi(r.URL.String()[1:])
	}
	if err != nil {
		log.Println("ERROR: Expected an integer path suffix")
		http.Error(w, "Need an integer", 400)
		return
	}

	rnd := rand.Intn(10)
	if rnd < 5 {
		http.ServeFile(w, r, "./dummy.png")
		return
	} else {
		http.ServeFile(w, r, "./dummy.pdf")
		return
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", serveRandomFile)
	http.HandleFunc("/health", healthCheck)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

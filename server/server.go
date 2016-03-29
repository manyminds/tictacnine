package server

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type server struct {
	handler http.Handler
}

func init() {
	rand.Seed(time.Now().UnixNano())
	filename = randString(10) + ".txt"
	log.Printf("Writing to %s\n", filename)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func logMove(w http.ResponseWriter, r *http.Request) {
	color := r.FormValue("color")
	x := r.FormValue("x")
	y := r.FormValue("y")

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Printf(err.Error())
	}

	defer f.Close()

	lineCount++
	if _, err = f.WriteString(fmt.Sprintf("%d,[%s],%s,%s\n", lineCount, color, x, y)); err != nil {
		log.Printf(err.Error())
	}
}

var filename = ""
var lineCount = 0

//NewServer returns a new configured http handler
func NewServer(dist string) http.Handler {
	mux := http.NewServeMux()
	fileHandler := http.FileServer(http.Dir(dist))
	mux.HandleFunc("/log", logMove)
	mux.Handle("/", fileHandler)
	s := server{}
	s.handler = mux
	return &s
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handler.ServeHTTP(w, r)
}

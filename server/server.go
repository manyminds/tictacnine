package server

import (
	"net/http"

	"github.com/manyminds/tictacnine/game"
)

type server struct {
	handler http.Handler
}

//NewServer returns a new configured http handler
func NewServer(dist string) http.Handler {
	mux := http.NewServeMux()
	fileHandler := http.FileServer(http.Dir(dist))
	mux.Handle("/", fileHandler)
	g := game.NewGame()
	mux.Handle("/game/", g)
	s := server{}
	s.handler = mux
	return &s
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handler.ServeHTTP(w, r)
}

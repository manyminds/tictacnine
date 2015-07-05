package server

import "net/http"

type server struct {
	handler http.Handler
}

//NewServer returns a new configured http handler
func NewServer(dist string) http.Handler {
	mux := http.NewServeMux()
	fileHandler := http.FileServer(http.Dir(dist))
	mux.Handle("/", fileHandler)
	s := server{}
	s.handler = mux
	return &s
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handler.ServeHTTP(w, r)
}

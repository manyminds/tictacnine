package main

import (
	"fmt"
	"net/http"

	"github.com/lucas-clemente/go-http-logger"
	"github.com/manyminds/tictacnine/server"
)

func main() {
	s := server.NewServer("./")
	port := 13337
	fmt.Printf("Server started on port %d.\n", port)
	err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", port), logger.Logger(s))
	if err != nil {
		panic(err)
	}

}
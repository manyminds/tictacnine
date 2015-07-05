package main

import (
	"fmt"
	"os"
	"path"

	godo "gopkg.in/godo.v1"
)

func tasks(p *godo.Project) {
	wd, _ := os.Getwd()
	ws := path.Join(wd, "Godeps/_workspace")
	godo.Env = fmt.Sprintf("GOPATH=%s::$GOPATH", ws)

	p.Task("build", godo.D{}, func() error {
		return godo.Run("go build -o ../bin/ttn", godo.In{"cmd/"})
	}).Watch("**/*.go")

	p.Task("server", godo.D{"build"}, func() {
		godo.Start("main.go", godo.M{"$in": "cmd"})
	}).Watch("**/*.go", "*.{go,json}").
		Debounce(3000)
}

func main() {
	godo.Godo(tasks)
}
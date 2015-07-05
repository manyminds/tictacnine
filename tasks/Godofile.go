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

	p.Task("frontend", godo.D{}, func() error {
		godo.Bash("mkdir -p frontend/css/dist/octicons")
		godo.Bash("mkdir -p frontend/css/dist/bootstrap")
		godo.Bash("cp bower_components/bootstrap/dist/css/bootstrap.min.css frontend/css/dist/bootstrap/")
		return godo.Bash("cp -r bower_components/octicons/octicons/ frontend/css/dist/octicons/")
	})

	p.Task("backend", godo.D{"frontend", "build"}, func() {
		godo.Start("main.go", godo.M{"$in": "cmd"})
	}).Watch("**/*.go", "*.{go,json}").
		Debounce(3000)
}

func main() {
	godo.Godo(tasks)
}
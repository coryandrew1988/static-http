package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"log"
	"os"
	"path/filepath"
)

func cwd() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func argOrDflt(i int, dflt string) string {
	if len(os.Args) > i {
		return os.Args[i]
	} else {
		return dflt
	}
}

func main() {
	dir := argOrDflt(1, "static")
	//port := argOrDflt(2, "3000")

	fmt.Printf("The current working directory is \"%s\".\r\n", cwd())
	fmt.Printf("The static file directory is \"%s\".\r\n", dir)
	//fmt.Printf("The port is \"%s\".\r\n", port)

	m := martini.Classic()

	//m.RunOnAddr(":" + port)

	m.Use(martini.Static(dir, martini.StaticOptions{Fallback: "/index.html"}))

	m.Run()
}

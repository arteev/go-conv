package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func usage() {
	fmt.Printf("cat file | %s > file_new\n", filepath.Base(os.Args[0]))
}

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		usage()
		os.Exit(1)
	}
	tr := transform.NewReader(os.Stdin, charmap.Windows1251.NewEncoder())
	_, err = io.Copy(os.Stdout, tr)
	if err != nil {
		log.Fatal(err)
	}
}

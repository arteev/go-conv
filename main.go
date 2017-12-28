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
	fmt.Printf("%s file > file_new\n", filepath.Base(os.Args[0]))
	fmt.Printf("cat file | %s > file_new\n", filepath.Base(os.Args[0]))
	os.Exit(1)
}

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}
	var r io.Reader
	if fi.Mode()&os.ModeNamedPipe == 0 {
		if len(os.Args) > 1 {
			f, err := os.Open(os.Args[1])
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			r = f
		} else {
			usage()
		}
	} else {
		r = os.Stdin
	}

	tr := transform.NewReader(r, charmap.Windows1251.NewEncoder())
	_, err = io.Copy(os.Stdout, tr)
	if err != nil {
		log.Fatal(err)
	}
}

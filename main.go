package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

var (
	toUtf8 = flag.Bool("utf8", false, "win1251 -> utf8")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\t%s [options] file > file_new\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\tcat file | %s [options] > file_new\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {

	flag.Usage = usage
	flag.Parse()

	fi, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}
	var r io.Reader
	if fi.Mode()&os.ModeNamedPipe == 0 {
		if len(os.Args) > 1 {
			f, err := os.OpenFile(os.Args[len(os.Args)-1], os.O_RDONLY, 0)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			r = f
		} else {
			flag.Usage()
		}
	} else {
		r = os.Stdin
	}

	var tr *transform.Reader
	if *toUtf8 {
		tr = transform.NewReader(r, charmap.Windows1251.NewDecoder())
	} else {
		tr = transform.NewReader(r, charmap.Windows1251.NewEncoder())
	}

	_, err = io.Copy(os.Stdout, tr)
	if err != nil {
		log.Fatal(err)
	}
}

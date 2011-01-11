package main

import "google/urlshortener"
import "flag"
import "fmt"
import "os"

var expand = flag.Bool("e", false, "expand URL")

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "usage: urlshortener [-e] URL\n")
		os.Exit(2)
	}
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
	}

	var u string
	var e os.Error
	if *expand {
		u, e = urlshorter.ExpandURL(args[0])
	} else {
		u, e = urlshorter.ShortenURL(args[0])
	}
	if e != nil {
		panic(e.String())
	} else {
		println(u)
	}
}

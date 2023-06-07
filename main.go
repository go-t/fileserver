package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var (
	flagV    = flag.Bool("v", true, "verbose")
	flagBind = flag.String("bind", ":8000", "http listen address")
	flagRoot = flag.String("root", ".", "document root dir")
)

func main() {
	flag.Parse()
	
	if *flagV {
		fmt.Printf("run file server on address %s dir %s\n", *flagBind, *flagRoot)
	}

	var handler http.Handler = http.FileServer(http.Dir(*flagRoot))
	if *flagV {
		next := handler
		handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			fmt.Printf("%s - %s %s - %v\n", r.RemoteAddr, r.Method, r.RequestURI, time.Since(start))
		})
	}

	err := http.ListenAndServe(*flagBind, handler)
	if err != nil {
		panic(err)
	}
}

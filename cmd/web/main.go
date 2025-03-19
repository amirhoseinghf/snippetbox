package main

import (
	"flag"
	"net/http"

	"github.com/charmbracelet/log"
)

func main() {

	addr := flag.String("addr", ":4000", "Http network address")
	flag.Parse()

	// infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	srv := &http.Server{
		Addr:    *addr,
		Handler: mux,
	}

	log.Infof("Starting Server on %s", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

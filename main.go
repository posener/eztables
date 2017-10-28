package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/posener/eztables/html"
	"github.com/posener/eztables/table"
)

var (
	addr = flag.String("addr", ":4242", "Listening address")
)

func main() {
	flag.Parse()

	_, port, err := net.SplitHostPort(*addr)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Listening on %s", *addr)
	log.Printf("You can browse to http://localhost:%s", port)

	err = http.ListenAndServe(*addr, handler())
	if err != nil {
		log.Panic(err)
	}
}

func handler() http.Handler {
	m := mux.NewRouter()
	m.Methods(http.MethodGet).Path("/chains/{chain}").HandlerFunc(chain)
	m.Methods(http.MethodGet).Path("/").HandlerFunc(index)
	return m
}

func index(w http.ResponseWriter, r *http.Request) {
	showTable(w, "")
}

func chain(w http.ResponseWriter, r *http.Request) {
	showTable(w, mux.Vars(r)["chain"])
}

func showTable(w http.ResponseWriter, chain string) {
	t, err := table.Load(chain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	html.Write(w, t)
}

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

	err := table.Test()
	if err != nil {
		log.Fatal(err)
	}

	_, port, err := net.SplitHostPort(*addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on %s", *addr)
	log.Printf("You can browse to http://localhost:%s", port)

	err = http.ListenAndServe(*addr, handler())
	if err != nil {
		log.Fatal(err)
	}
}

func handler() http.Handler {
	m := mux.NewRouter()
	m.Methods(http.MethodGet).Path("/tables/{table}").HandlerFunc(get)
	m.Methods(http.MethodGet).Path("/").HandlerFunc(index)
	return m
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/tables/filter", http.StatusFound)
}

func get(w http.ResponseWriter, r *http.Request) {
	tables, err := table.Load()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(tables) == 0 {
		http.Error(w, "Could not load tables", http.StatusNotFound)
		return
	}

	tableName := mux.Vars(r)["table"]
	var cur table.Table
	others := make([]table.Table, 0, len(tables)-1)

	for _, t := range tables {
		if t.Name == tableName {
			cur = t
		} else {
			others = append(others, t)
		}
	}
	html.Write(w, cur, others)
}

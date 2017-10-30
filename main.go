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
	}

	tableName := mux.Vars(r)["table"]
	var cur table.Table
	names := make([]string, 0, len(tables)-1)

	for _, t := range tables {
		if t.Name == tableName {
			cur = t
		} else {
			names = append(names, t.Name)
		}
	}
	html.Write(w, cur, names)
}

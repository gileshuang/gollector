package main

import (
	"flag"
	"log"
	"net/http"
	"sync"

	"github.com/gileshuang/gollector/lib/model"
)

var (
	servHTTP     string
	allHostsInfo map[string]*model.HostInfo
	lockHostInfo *sync.RWMutex
)

func init() {
	flag.StringVar(&servHTTP, "http", ":8880",
		"HTTP service address.")
	flag.Parse()

	allHostsInfo = make(map[string]*model.HostInfo)
	lockHostInfo = new(sync.RWMutex)
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/hosts/", handleHosts)
	http.HandleFunc("/update/", handleUpdate)
	http.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("static"))))

	log.Println("Starting server at", servHTTP)
	err := http.ListenAndServe(servHTTP, nil)
	if err != nil {
		log.Println("ListenAndServe: ", err)
		return
	}
}

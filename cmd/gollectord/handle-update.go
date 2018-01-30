package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gileshuang/gollector/lib/model"
)

func handleUpdate(w http.ResponseWriter, r *http.Request) {
	var (
		pathSplits []string
	)

	// Check request.Method
	if r.Method != http.MethodGet && r.Method != http.MethodPost &&
		r.PostForm.Get("method") != "delete" {
		t, _ := template.ParseFiles("template/rawtext.gtpl")
		w.WriteHeader(http.StatusMethodNotAllowed)
		t.Execute(w, "Invalid http method.")
		return
	}

	// Check request.URL.Path
	pathSplits = make([]string, 3, 3)
	pathSplits = strings.Split(r.URL.Path, "/")
	if len(pathSplits) < 3 || pathSplits[2] == "" {
		t, _ := template.ParseFiles("template/rawtext.gtpl")
		w.WriteHeader(http.StatusBadRequest)
		t.Execute(w, "Hostname required.")
		return
	}
	if len(pathSplits) > 3 && pathSplits[3] != "" {
		t, _ := template.ParseFiles("template/rawtext.gtpl")
		w.WriteHeader(http.StatusNotFound)
		t.Execute(w, "There should not be a parameter after hostname.")
		return
	}
	hostName := pathSplits[2]

	err := r.ParseForm()
	if err != nil {
		log.Println("handleUpdate=>ParseForm", err)
		t, _ := template.ParseFiles("template/rawtext.gtpl")
		w.WriteHeader(http.StatusBadRequest)
		t.Execute(w, "Parse post form error, please check the post data.")
		return
	}

	if _, ok := allHostsInfo[hostName]; !ok {
		allHostsInfo[hostName] = new(model.HostInfo)
	}
	allHostsInfo[hostName].HostName = hostName
	if allHostsInfo[hostName].Info == nil {
		allHostsInfo[hostName].Info = make(map[string]*model.AtomInfo)
	}
	for k := range r.PostForm {
		if _, ok := allHostsInfo[hostName].Info[k]; !ok {
			allHostsInfo[hostName].Info[k] = new(model.AtomInfo)
		}
		allHostsInfo[hostName].Info[k].Value = r.PostFormValue(k)
		allHostsInfo[hostName].Info[k].UpdateTime = time.Now().UTC()
	}

	t, _ := template.ParseFiles("template/rawtext.gtpl")
	w.WriteHeader(http.StatusOK)
	t.Execute(w, "POST success.")

	// DEBUG
	log.Println("hostinfo:", allHostsInfo[hostName])
	log.Println("atominfo-a:", allHostsInfo[hostName].Info["a"])
	log.Println("atominfo-b:", allHostsInfo[hostName].Info["b"])
	return
}

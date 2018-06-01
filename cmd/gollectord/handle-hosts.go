package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gileshuang/gollector/lib/model"
)

func handleHosts(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		t, _ := template.ParseFiles("template/rawtext.html")
		w.WriteHeader(http.StatusMethodNotAllowed)
		t.Execute(w, "Invalid http method.")
		return
	}

	var (
		res        map[string]*model.HostInfo
		bSrchName  bool
		bSrchKey   bool
		bSrchValue bool
	)
	res = make(map[string]*model.HostInfo)
	bSrchName = false
	bSrchKey = false
	bSrchValue = false

	// Filter
	keyWord := r.PostForm.Get("keyword")
	for _, rField := range r.PostForm["field"] {
		switch rField {
		case "hostname":
			bSrchName = true
		case "key":
			bSrchKey = true
		case "value":
			bSrchValue = true
		}
	}
	lockHostInfo.RLock()
	for k, v := range allHostsInfo {
		if bSrchName == true && strings.Contains(k, keyWord) {
			res[k] = allHostsInfo[k]
		} else {
			for ki, vi := range v.Info {
				if (bSrchKey == true && strings.Contains(ki, keyWord)) ||
					(bSrchValue == true && strings.Contains(vi.Value, keyWord)) {
					if _, ok := res[k]; !ok {
						res[k] = new(model.HostInfo)
						res[k].HostName = k
						res[k].Info = make(map[string]*model.AtomInfo)
					}
					res[k].Info[ki] = allHostsInfo[k].Info[ki]
				}
			}
		}
	}
	lockHostInfo.RUnlock()

	t, err := template.ParseFiles("template/hosts.html")
	if err != nil {
		log.Println("handleHosts.template.ParseFiles:", err)
	}
	t.Execute(w, res)
	return
}

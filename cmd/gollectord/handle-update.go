package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
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
		t, _ := template.ParseFiles("template/rawtext.html")
		w.WriteHeader(http.StatusMethodNotAllowed)
		t.Execute(w, "Invalid http method.")
		return
	}

	// Check request.URL.Path
	pathSplits = make([]string, 3, 3)
	pathSplits = strings.Split(r.URL.Path, "/")
	if len(pathSplits) < 3 || pathSplits[2] == "" {
		t, _ := template.ParseFiles("template/rawtext.html")
		w.WriteHeader(http.StatusBadRequest)
		t.Execute(w, "Hostname required.")
		return
	}
	if len(pathSplits) > 3 && pathSplits[3] != "" {
		t, _ := template.ParseFiles("template/rawtext.html")
		w.WriteHeader(http.StatusNotFound)
		t.Execute(w, "There should not be a parameter after hostname.")
		return
	}
	hostName := pathSplits[2]

	defer r.Body.Close()
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("handleUpdate=>ReadRequestBody", err)
		t, _ := template.ParseFiles("template/rawtext.html")
		w.WriteHeader(http.StatusBadRequest)
		t.Execute(w, "Read request body error.")
		return
	}

	hInfo := new(model.HostInfo)
	hInfo.HostName = hostName
	jsonErr := json.Unmarshal(rBody, hInfo)
	if jsonErr != nil {
		log.Println("handleUpdate=>JsonUnmarshal", jsonErr)
		t, _ := template.ParseFiles("template/rawtext.html")
		w.WriteHeader(http.StatusBadRequest)
		t.Execute(w, "Parse json failed, please post request as given json format.")
		return
	}

	if _, ok := allHostsInfo[hostName]; ok {
		for k, v := range hInfo.Info {
			v.UpdateTime = time.Now().UTC()
			allHostsInfo[hostName].Info[k] = v
		}
	} else {
		for k := range hInfo.Info {
			hInfo.Info[k].UpdateTime = time.Now().UTC()
		}
		allHostsInfo[hostName] = hInfo
	}

	t, _ := template.ParseFiles("template/rawtext.html")
	w.WriteHeader(http.StatusOK)
	t.Execute(w, "POST success.")

	return
}

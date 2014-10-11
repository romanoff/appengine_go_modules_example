package app

import (
	"fmt"
	"net/http"
	"appengine"
	"appengine/urlfetch"
	"io/ioutil"
)

func init() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/callmobile", mobileHandler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world from app")
}

func mobileHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	hostname, err := appengine.ModuleHostname(c, "mobile-frontend", "", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	client := urlfetch.Client(c)
	resp, err := client.Get("http://"+hostname+"/")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%v", string(body))
}

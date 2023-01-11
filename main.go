package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type post struct {
	UserID int    `json:"userID"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

//raw html
var form = `
<h1>Post #{{.ID}}</h1>
<div>{{printf "User %d" .UserID}}</div>
<div>{{printf "Title is %s" .Title}}</div>
<div>{{printf "Body is  %s" .Body}}</div>`

//use printf to actually print strings with formatting

func handler(w http.ResponseWriter, r *http.Request) {
	const baseUrl = "https://jsonplaceholder.typicode.com/"
	resp, err := http.Get(baseUrl + r.URL.Path[1:])

	if err != nil {
		//signal an error with http.Error
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	//ensure to close the body that was open else server runs out of sockets
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		//signal an error with http.Error
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	var item post

	err = json.Unmarshal(body, &item)

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	_template := template.New("sample")
	_template.Parse(form)
	_template.Execute(w, item)

}

func main() {
	http.HandleFunc("/", handler) // route

	log.Fatal(http.ListenAndServe(":8080", nil)) // => http://localhost:8080
}

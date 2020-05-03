package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Result struct {
	Id      int
	Subject string
	Time    int
}

var Results map[string][]*Result

func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "hello")
}

func main() {
	// Results = make(map[string][]*Result)

	// result1 := Result{Id: 1, Subject: "Math", Time: 2}
	// result2 := Result{Id: 2, Subject: "Science", Time: 6}
	// Results["hoge"] = append(Results["hoge"], &result1)
	// Results["hoge"] = append(Results["hoge"], &result2)

	// for _, result := range Results["hoge"] {
	// 	fmt.Println(result)
	// }

	// result := Results["hoge"]
	// output, _ := json.MarshalIndent(&result, "", "\t\t")
	// fmt.Println(output)

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/record/", handleRequest)
	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGet(w, r)
	case "POST":
		handlePost(w, r)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	result := Results["hoge"]
	output, _ := json.MarshalIndent(&result, "", "\t\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)

	r.Body.Read(body)
	var result Result
	json.Unmarshal(body, &result)

	Results["hoge"] = append(Results["hoge"], &result)
	w.WriteHeader(200)
}

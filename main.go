package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Id      int
	Subject string
	Time    int
}

var Results map[string][]*Result

func record(user string, result Result) {
	Results[user] = append(Results[user], &result)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func recordRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "5")
}

func main() {
	Results = make(map[string][]*Result)

	result1 := Result{Id: 1, Subject: "Math", Time: 2}
	fmt.Println(result1)
	record("hoge", result1)
	record("hoge", result1)

	for _, result := range Results["hoge"] {
		fmt.Println(result)
	}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/recordRequest", recordRequest)
	server.ListenAndServe()

}

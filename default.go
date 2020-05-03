package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type StudyHistory struct {
	UserId     int         `json:"userId"`
	UserName   string      `json:"userName"`
	StudyInfos []StudyInfo `json:"studyInfos"`
}

type StudyInfo struct {
	StudyId   int    `json:"studyId"`
	SubId     int    `json:"subId"`
	StudyTime int    `json:"studyTime"`
	DateTime  string `json:"dateTime"`
}

var studyHistory StudyHistory

func main() {

	jsonFile, err := os.Open("default.json")
	defer jsonFile.Close()
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}
	fmt.Println(string(jsonData))

	json.Unmarshal(jsonData, &studyHistory)
	fmt.Println(studyHistory)

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
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
	output, _ := json.MarshalIndent(&studyHistory, "", "\t\t")
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Println(string(body))

	var studyInfo StudyInfo
	json.Unmarshal(body, &studyInfo)
	fmt.Println(studyInfo)
	w.WriteHeader(200)
	return
}

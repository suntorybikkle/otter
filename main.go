package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type MyTime struct {
	*time.Time
}

func (myTime MyTime) format() string {
	return myTime.Time.Format("2006-01-02 15:04:05")
}

func (myTime *MyTime) UnmarshalJSON(data []byte) (err error) {
	t, err := time.Parse("\"2006-01-02 15:04:05\"", string(data))
	*myTime = MyTime{&t}
	return
}

type StudyHistory struct {
	UserId     int         `json:"userId"`
	UserName   string      `json:"userName"`
	StudyInfos []StudyInfo `json:"studyInfos"`
}

type StudyInfo struct {
	Id        int `json:"studyId"`
	UserId    int
	StudyId   int
	SubjectId int `json:"subId"`
	StudyTime int `json:"studyTime"`
	// DateTime  string `json:"dateTime"`
	DateTime MyTime `json:"dateTime"`
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/record/", handleRequest)
	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	studyInfos, err := GetAllStudyInfo(1)
	if err != nil {
		log.Println(err)
		return
	}

	studyHistory := StudyHistory{UserId: 1, UserName: "ktguy"}
	for _, studyInfo := range studyInfos {
		studyHistory.StudyInfos = append(studyHistory.StudyInfos, studyInfo)
	}

	output, _ := json.MarshalIndent(&studyHistory, "", "\t\t")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(output)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	log.Println(string(body))

	var studyInfo StudyInfo
	json.Unmarshal(body, &studyInfo)
	studyInfo.UserId = 1
	log.Println(studyInfo)
	studyInfo.Create()

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(200)
	return
}

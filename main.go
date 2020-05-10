package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type StudyInfo struct {
	// これがエンティティ?
	Id        int
	UserId    int
	StudyId   int
	SubjectId int
	StudyTime int
	DateTime  time.Time
}

type StudyReportJson struct {
	UserId     int             `json:"userId"`
	UserName   string          `json:"userName"`
	StudyInfos []StudyInfoJson `json:"studyInfos"`
}

func (studyReportJson *StudyReportJson) retrieve(studyInfos []StudyInfo) {
	for _, studyInfo := range studyInfos {
		var studyInfoJson StudyInfoJson
		studyInfoJson.retrieve(studyInfo)
		studyReportJson.StudyInfos = append(studyReportJson.StudyInfos, studyInfoJson)
	}
}

type StudyInfoJson struct {
	Id        int    `json:"studyId"`
	SubjectId int    `json:"subId"`
	StudyTime int    `json:"studyTime"`
	DateTime  string `json:"dateTime"`
}

func (studyInfoJson *StudyInfoJson) retrieve(studyInfo StudyInfo) {
	studyInfoJson.Id = studyInfo.Id
	studyInfoJson.SubjectId = studyInfo.SubjectId
	studyInfoJson.StudyTime = studyInfo.StudyTime
	studyInfoJson.DateTime = studyInfo.DateTime.Format("2006-01-02 15:04:05")
}

type StudyPostJson struct {
	UserId    int    `json:"userId"`
	SubjectId int    `json:"subId"`
	StudyTime int    `json:"studyTime"`
	DateTime  string `json:"dateTime"`
}

func (studyPostJson StudyPostJson) convert() (studyInfo StudyInfo) {
	studyInfo.UserId = 1
	studyInfo.SubjectId = studyPostJson.SubjectId
	studyInfo.StudyTime = studyPostJson.StudyTime
	studyInfo.DateTime, _ = time.Parse("2006-01-02 15:04:05", studyPostJson.DateTime)
	return
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
	studyReportJson := StudyReportJson{UserId: 1, UserName: "ktguy"}
	studyReportJson.retrieve(studyInfos)

	output, _ := json.MarshalIndent(&studyReportJson, "", "\t\t")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(output)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var studyPostJson StudyPostJson
	json.Unmarshal(body, &studyPostJson)

	studyInfo := studyPostJson.convert()
	studyInfo.Create()

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(200)
	return
}

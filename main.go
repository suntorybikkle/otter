package main

import (
	"encoding/json"
	"log"
	"net/http"
	"test/domain/model"
	"test/interfaces"
	"time"
)

type StudyReportJson struct {
	UserId     int             `json:"userId"`
	UserName   string          `json:"userName"`
	StudyInfos []StudyInfoJson `json:"studyInfos"`
}

type StudyInfoJson struct {
	Id        int    `json:"studyId"`
	SubjectId int    `json:"subId"`
	StudyTime int    `json:"studyTime"`
	DateTime  string `json:"dateTime"`
}

func ReportResponseAdapter(studyInfos []model.StudyInfo) (studyReportJson StudyReportJson) {
	var studyInfosJson []StudyInfoJson
	for _, studyInfo := range studyInfos {
		studyInfosJson = append(studyInfosJson, StudyInfoJson{
			Id:        studyInfo.Id,
			SubjectId: studyInfo.SubjectId,
			StudyTime: studyInfo.StudyTime,
			DateTime:  studyInfo.DateTime.Format("2006-01-02 15:04:05"),
		})
	}
	studyReportJson = StudyReportJson{UserId: 1, UserName: "ktguy", StudyInfos: studyInfosJson}
	return
}

type StudyPostJson struct {
	UserId    int    `json:"userId"`
	SubjectId int    `json:"subId"`
	StudyTime int    `json:"studyTime"`
	DateTime  string `json:"dateTime"`
}

func StudyPostRequestAdapter(studyPostJson StudyPostJson) (studyInfo model.StudyInfo) {
	studyDate, err := time.Parse("2006-01-02 15:04:05", studyPostJson.DateTime)
	if err != nil {
		log.Println(err)
	}
	studyInfo = model.StudyInfo{
		UserId:    1,
		SubjectId: studyPostJson.SubjectId,
		StudyTime: studyPostJson.StudyTime,
		DateTime:  studyDate,
	}
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
	studyInfos, err := interfaces.GetAllStudyInfo(1)
	if err != nil {
		log.Println(err)
		return
	}

	studyReportJson := ReportResponseAdapter(studyInfos)
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

	studyInfo := StudyPostRequestAdapter(studyPostJson)
	// studyInfo.Create()
	interfaces.Create(studyInfo)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(200)
	return
}

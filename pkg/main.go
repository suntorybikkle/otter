package main

import (
	"encoding/json"
	"log"
	"net/http"
	"otter/pkg/infrastructure/postgresql"
	"otter/pkg/interfaces/controller"
	"otter/pkg/interfaces/gateway"
	"otter/pkg/usecases/interactor"
)

var studyInfoController controller.StudyInfoController

func main() {
	dbHandler := postgresql.NewPsqlHandler("user=ldb dbname=ldb password=ldb sslmode=disable")

	studyInfoController.StudyInfoInteractor = interactor.StudyInfoInteractor{
		StudyInfoRepository: &gateway.StudyInfoRepository{
			SqlHandler: dbHandler,
		},
	}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/record/", handleRequest)
	_ = server.ListenAndServe()
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
	studyReportJson, err := studyInfoController.GetStudyReport(1)
	if err != nil {
		log.Println(err)
		return
	}
	output, _ := json.MarshalIndent(&studyReportJson, "", "\t\t")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, _ = w.Write(output)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	_, _ = r.Body.Read(body)

	var studyPostJson controller.StudyPostJson
	_ = json.Unmarshal(body, &studyPostJson)

	err = studyInfoController.PostStudy(studyPostJson)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(200)
	return
}

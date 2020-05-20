package interfaces

import (
	"encoding/json"
	"log"
	"otter/pkg/domain"
	"otter/pkg/usecases/interactor"
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

type StudyPostJson struct {
	UserId    int    `json:"userId"`
	SubjectId int    `json:"subId"`
	StudyTime int    `json:"studyTime"`
	DateTime  string `json:"dateTime"`
}

func StudyPostRequestAdapter(studyPostJson StudyPostJson) (studyInfo domain.StudyInfo) {
	studyDate, err := time.Parse("2006-01-02 15:04:05", studyPostJson.DateTime)
	if err != nil {
		log.Println(err)
	}
	studyInfo = domain.StudyInfo{
		UserId:    1,
		SubjectId: studyPostJson.SubjectId,
		StudyTime: studyPostJson.StudyTime,
		DateTime:  studyDate,
	}
	return
}

var UseCase interactor.StudyInfoInteractor

func main() {
	dbHandler := infrastructure.NewPsqlHandler("user=ldb dbname=ldb password=ldb sslmode=disable")

	UseCase = interactor.StudyInfoInteractor{
		StudyInfoRepository: &interfaces.StudyInfoRepository{
			SqlHandler: dbHandler,
		},
	}

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
	studyReport, err := UseCase.GetStudyReport(1)
	log.Println(studyReport)
	// studyInfos, err := interfaces.GetAllStudyInfo(1)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// studyReportJson := ReportResponseAdapter(studyInfos)
	// output, _ := json.MarshalIndent(&studyReportJson, "", "\t\t")

	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Write(output)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var studyPostJson StudyPostJson
	json.Unmarshal(body, &studyPostJson)

	studyInfo := StudyPostRequestAdapter(studyPostJson)
	UseCase.AddStudyInfo(studyInfo)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(200)
	return
}

package webservice

import (
	"encoding/json"
	"log"
	"net/http"
	"otter/pkg/interfaces/controller"
)

type WebserviceHandler struct {
	StudyInfoController controller.StudyInfoController
}

func NewWebserviceHandler(studyInfoController controller.StudyInfoController) WebserviceHandler {
	return WebserviceHandler{
		StudyInfoController: studyInfoController,
	}
}

func (handler *WebserviceHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handler.handleGet(w, r)
	case "POST":
		err = handler.handlePost(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (handler *WebserviceHandler) handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	studyReportJson, _ := handler.StudyInfoController.GetStudyReport(1)
	if err != nil {
		log.Println(err)
		return
	}
	output, _ := json.MarshalIndent(&studyReportJson, "", "\t\t")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(output)
	return
}

func (handler *WebserviceHandler) handlePost(w http.ResponseWriter, r *http.Request) (err error) {

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var studyPostJson controller.StudyPostJson
	json.Unmarshal(body, &studyPostJson)
	err = handler.StudyInfoController.PostStudy(studyPostJson)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(200)
	return
}

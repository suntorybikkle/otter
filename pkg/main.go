package main

import (
	"net/http"
	"otter/pkg/infrastructure/postgresql"
	"otter/pkg/infrastructure/webservice"
	"otter/pkg/interfaces/controller"
	"otter/pkg/interfaces/gateway"
	"otter/pkg/usecases/interactor"
)

var studyInfoController controller.StudyInfoController

func main() {
	dbHandler := postgresql.NewPsqlHandler("user=ldb dbname=ldb password=ldb sslmode=disable")
	studyInfoRepository := gateway.NewStudyInfoRepository(dbHandler)
	studyInfoInteractor := interactor.NewStudyInfoInteractor(studyInfoRepository)
	studyInfoController := controller.NewStudyInfoController(studyInfoInteractor)
	webserviceHandler := webservice.NewWebserviceHandler(studyInfoController)

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/record/", webserviceHandler.HandleRequest)
	_ = server.ListenAndServe()
}

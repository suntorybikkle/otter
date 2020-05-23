package controller

import (
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
	StudyId   int    `json:"studyId"`
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

type StudyInfoController struct {
	StudyInfoInteractor interactor.StudyInfoInteractor
}

func NewStudyInfoController(studyInfoInteractor interactor.StudyInfoInteractor) StudyInfoController {
	return StudyInfoController{
		StudyInfoInteractor: studyInfoInteractor,
	}
}

func (controller *StudyInfoController) PostStudy(studyPostJson StudyPostJson) (err error) {
	studyDate, err := time.Parse("2006-01-02 15:04:05", studyPostJson.DateTime)
	if err != nil {
		log.Println(err)
		return
	}
	studyInfo := domain.StudyInfo{
		UserId:    studyPostJson.UserId,
		SubjectId: studyPostJson.SubjectId,
		StudyTime: studyPostJson.StudyTime,
		DateTime:  studyDate,
	}
	err = controller.StudyInfoInteractor.AddStudyInfo(studyInfo)
	if err != nil {
		return
	}
	return
}

func (controller *StudyInfoController) GetStudyReport(userId int) (studyReportJson StudyReportJson, err error) {
	studyReport, err := controller.StudyInfoInteractor.GetStudyReport(userId)
	if err != nil {
		log.Println(err)
		return
	}
	studyReportJson = studyReportResponseAdapter(studyReport)
	return
}

func studyReportResponseAdapter(studyReport interactor.StudyReport) (studyReportJson StudyReportJson) {
	var studyInfosJson []StudyInfoJson
	for _, studyInfo := range studyReport.StudyInfos {
		studyInfosJson = append(studyInfosJson, StudyInfoJson{
			StudyId:   studyInfo.StudyId,
			SubjectId: studyInfo.SubjectId,
			StudyTime: studyInfo.StudyTime,
			DateTime:  studyInfo.DateTime.Format("2006-01-02 15:04:05"),
		})
	}
	studyReportJson = StudyReportJson{UserId: studyReport.UserId, UserName: studyReport.UserName, StudyInfos: studyInfosJson}
	return
}

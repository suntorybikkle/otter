package interactor

import (
	"log"
	"otter/pkg/domain"
	"otter/pkg/usecases/repository"
	"time"
)

type StudyReport struct {
	UserId     int
	UserName   string
	StudyInfos []StudyInfo
}

type StudyInfo struct {
	StudyId   int
	SubjectId int
	StudyTime int
	DateTime  time.Time
}

type StudyInfoInteractor struct {
	StudyInfoRepository repository.StudyInfoRepository
}

func (interactor *StudyInfoInteractor) AddStudyInfo(studyInfo domain.StudyInfo) (err error) {
	err = interactor.StudyInfoRepository.Create(studyInfo)
	return
}

func (interactor *StudyInfoInteractor) GetStudyReport(userId int) (studyReport StudyReport, err error) {
	domainStudyInfos, err := interactor.StudyInfoRepository.GetAll(userId)
	var studyInfos []StudyInfo
	for _, domainStudyInfo := range domainStudyInfos {
		studyInfos = append(studyInfos, StudyInfo{
			StudyId:   domainStudyInfo.Id,
			SubjectId: domainStudyInfo.SubjectId,
			StudyTime: domainStudyInfo.StudyTime,
			DateTime:  domainStudyInfo.DateTime,
		})
	}
	return StudyReport{UserId: userId, UserName: "ktguy", StudyInfos: studyInfo}
}

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

func (studyInfoInteractor *StudyInfoInteractor) AddStudyInfo(studyInfo domain.StudyInfo) (err error) {
	_, err = studyInfoInteractor.StudyInfoRepository.Create(studyInfo)
	return
}

func (studyInfoInteractor *StudyInfoInteractor) GetStudyReport(userId int) (studyReport StudyReport, err error) {
	domainStudyInfos, err := studyInfoInteractor.StudyInfoRepository.GetAll(userId)
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

package repository

imoprt (
	"test/domain/model"
)

type StudyInfoRepository interface {
	GetLastStudyInfo(id int) (studyInfo model.StudyInfo, err error) 
	GetAllStudyInfo(userId int) (studyInfos []model.StudyInfo, err error)
	Create(studyInfo model.StudyInfo) (err error)
}

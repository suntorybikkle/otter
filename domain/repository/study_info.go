package repository

imoprt (
	"test/domain/model"
)

type StudyInfoRepository interface {
	GetLast(id int) (*studyInfo model.StudyInfo, err error) 
	GetAll(userId int) (*studyInfos []model.StudyInfo, err error)
	Create(*studyInfo model.StudyInfo) (err error)
}

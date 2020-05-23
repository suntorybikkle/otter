package repository

import (
	"test/domain/model"
)

type StudyInfoRepository interface {
	GetLast(userId int) (*model.StudyInfo, error)
	GetAll(userId int) ([]*model.StudyInfo, error)
	Create(studyInfo *model.StudyInfo) (*model.StudyInfo, error)
	Update(studyInfo *model.StudyInfo) (*model.StudyInfo, error)
}

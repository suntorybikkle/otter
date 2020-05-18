package repository

import (
	"otter/pkg/domain"
	"otter/pkg/usecases/repository"
)

type StudyInfoRepository interface {
	GetLast(userId int) (*domain.StudyInfo, error)
	GetAll(userId int) ([]*domain.StudyInfo, error)
	Create(studyInfo *domain.StudyInfo) (*domain.StudyInfo, error)
	Update(studyInfo *domain.StudyInfo) (*domain.StudyInfo, error)
}

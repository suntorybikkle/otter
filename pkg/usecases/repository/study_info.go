package repository

import (
	"otter/pkg/domain"
)

type StudyInfoRepository interface {
	Create(*domain.StudyInfo) error
	GetAll(int) ([]domain.StudyInfo, error)
	GetLast(int) (domain.StudyInfo, error)
	Update(domain.StudyInfo) error
}

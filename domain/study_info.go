package domain

import (
	"time"
)

type StudyInfo struct {
	Id        int
	UserId    int
	StudyId   int
	SubjectId int
	StudyTime int
	DateTime  time.Time
}

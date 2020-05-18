package domain

import (
	"time"
)

type StudyInfo struct {
	Id        int
	UserId    int
	SubjectId int
	StudyTime int
	DateTime  time.Time
}

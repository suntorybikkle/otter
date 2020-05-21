package gateway

import (
	"otter/pkg/domain"
	"time"
)

type StudyInfoRepository struct {
	SqlHandler
}

func (repo *StudyInfoRepository) Create(studyInfo *domain.StudyInfo) (err error) {
	result, err := repo.SqlHandler.Execute("INSERT INTO study_infos(user_id, subject_id, study_time, date_time) VALUES($1, $2, $3, $4) RETURNING id",
		studyInfo.UserId, studyInfo.SubjectId, studyInfo.StudyTime, studyInfo.DateTime)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	studyInfo.Id = int(id64)
	return
}

func (repo *StudyInfoRepository) GetAll(userId int) (studyInfos []domain.StudyInfo, err error) {
	rows, err := repo.SqlHandler.Query("SELECT id, subject_id, study_time, date_time FROM study_infos WHERE user_id = $1", userId)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int
		var subjectId int
		var studyTime int
		var dateTime time.Time
		if err = rows.Scan(&id, &subjectId, &studyTime, &dateTime); err != nil {
			return
		}
		studyInfo := domain.StudyInfo{
			Id:        id,
			UserId:    userId,
			SubjectId: subjectId,
			StudyTime: studyTime,
			DateTime:  dateTime,
		}
		studyInfos = append(studyInfos, studyInfo)
	}
	return
}

func (repo *StudyInfoRepository) GetLast(userId int) (studyInfo domain.StudyInfo, err error) {
	row, err := repo.SqlHandler.Query("SELECT id, subject_id, study_time, date_time FROM study_infos WHERE user_id = $1 ORDER BY Id DESC LIMIT 1", userId)
	defer row.Close()
	if err != nil {
		return
	}
	var id int
	var subjectId int
	var studyTime int
	var dateTime time.Time
	row.Next()
	if err = row.Scan(&id, &subjectId, &studyTime, &dateTime); err != nil {
		return
	}
	studyInfo.Id = id
	studyInfo.UserId = userId
	studyInfo.SubjectId = subjectId
	studyInfo.StudyTime = studyTime
	studyInfo.DateTime = dateTime
	return
}

func (repo *StudyInfoRepository) Update(studyInfo domain.StudyInfo) (err error) {
	_, err = repo.SqlHandler.Execute("UPDATE study_infos SET user_id = $2, subject_id = $3, study_time = $4, date_time = $5 WHERE id = $1",
		studyInfo.Id, studyInfo.UserId, studyInfo.SubjectId, studyInfo.StudyTime, studyInfo.DateTime)
	return
}

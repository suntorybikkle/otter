package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Record struct {
	Id        int
	UserId    int
	SubjectId int
	StudyTime int
	DateTime  string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=ldb dbname=ldb password=ldb sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetLastRecord(id int) (record Record, err error) {
	record = Record{}
	err = Db.QueryRow("SELECT id, user_id, subject_id, study_time, date_time FROM records WHERE user_id = $1 ORDER BY Id DESC LIMIT 1", id).Scan(
		&record.Id, &record.UserId, &record.SubjectId, &record.StudyTime, &record.DateTime)
	return
}

func GetAllRecord(userId int) (records []Record, err error) {
	rows, err := Db.Query("SELECT id, user_id, subject_id, study_time, date_time FROM records WHERE user_id = $1", userId)
	defer rows.Close()
	if err != nil {
		log.Println(err)
		return
	}
	for rows.Next() {
		record := Record{}
		err = rows.Scan(&record.Id, &record.UserId, &record.SubjectId, &record.StudyTime, &record.DateTime)
		if err != nil {
			log.Println(err)
			return
		}
		records = append(records, record)
	}
	return
}

func (record *Record) Create() (err error) {
	statement := "INSERT INTO records(user_id, subject_id, study_time, date_time) VALUES($1, $2, $3, $4) RETURNING id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(record.UserId, record.SubjectId, record.StudyTime, record.DateTime).Scan(&record.Id)
	return
}

func main() {
	record, _ := GetLastRecord(1)
	fmt.Println(record)
	records, _ := GetAllRecord(1)
	fmt.Println(records)
	record = Record{UserId: 1, SubjectId: 1, StudyTime: 1234, DateTime: "2020-09-09 09:09:09"}
	fmt.Println(record)
	record.Create()
	fmt.Println(record)
	records, _ = GetAllRecord(1)
	fmt.Println(records)
	record, _ = GetLastRecord(1)
	fmt.Println(record)
}

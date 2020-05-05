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
	fmt.Println("init")
	var err error
	Db, err = sql.Open("postgres", "user=ldb dbname=ldb password=ldb sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetRecord(id int) (record Record, err error) {
	record = Record{}
	err = Db.QueryRow("SELECT id, user_id, subject_id, study_time, date_time FROM records WHERE id = $1", id).Scan(
		&record.Id, &record.UserId, &record.SubjectId, &record.StudyTime, &record.DateTime)
	return
}

func GetAllRecord(userId int) (records []Record, err error) {
	rows, err := Db.Query("SELECT id, user_id, subject_id, study_time, date_time FROM records WHERE user_id = $1", userId)
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
	rows.Close()
	return
}

func main() {
	fmt.Println("main")
	record, _ := GetRecord(1)
	fmt.Println(record)
	record, _ = GetRecord(2)
	records, _ := GetAllRecord(1)
	fmt.Println(records)
	for _, record := range records {
		fmt.Println(record)
	}
}

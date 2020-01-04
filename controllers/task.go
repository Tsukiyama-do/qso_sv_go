package task

import (
	"log"
	"strconv"
	"time"

	task "../models"
	//  "log"
)

// idとテキストを保持する構造体
type Task struct {
	ID   int
	Text string
}

func NewTask() Task {
	return Task{}
}

// タスク構造体一覧を返す
func (c Task) GetAll() interface{} {

	// テストデータとして５つタスクを作成
	tasks := make([]*Task, 5)
	for i := 1; i <= 5; i++ {
		tasks[i-1] = &Task{ID: i, Text: "Task Text " + strconv.Itoa(i)}
	}

	return tasks
}

/*
func (c Task) Create(text string) {
    repo := task.NewTaskRepository()
    repo.Create(text)
}
*/
func (c Task) SearchDB(s_callsign string, s_fr string, s_to string) *[]task.QSLstr {
	var f_chk int //  pattern of checking
	const s_fmt = "20060102"
	repo := task.NewTaskRepository()
	if s_callsign == "" {
		f_chk = 1 // search by period
		if s_fr == "" {
			s_fr = time.Now().Format(s_fmt)
			s_to = time.Now().Format(s_fmt)
		} else {
			if s_to == "" {
				s_to = s_fr
			}
		}
	} else {
		f_chk = 0 // search by callsign
	}
	log.Printf("Callsign : %s, From : %s, To : %s,flag :  %d\n", s_callsign, s_fr, s_to, f_chk)
	qsl_sp := repo.Retrieve(s_callsign, s_fr, s_to, f_chk)
	return qsl_sp
}

func (c Task) UploadDB(s_callsign string, s_filename string) error { // to register information of upload to database

	repo := task.NewTaskRepository()
	//  log.Printf("UploadDB : Callsign : %s, Filename : %s\n", s_callsign, s_filename)

	qsl_sp := repo.InsertUpload(s_callsign, s_filename)
	return qsl_sp

}

package task

import (
//  "fmt"
//  "github.com/syndtr/goleveldb/leveldb"
  "time"
//  "strconv"
  "database/sql"
//  "fmt"
  "log"
//  "time"
  "strings"
  _ "github.com/mattn/go-sqlite3"
  "../env"
)


type Task struct {
    ID   int
    Text string
}

type QSLstr struct{
    ID int
    Callsign string
    Datetime string
    Files string
}


type Tasks []Task

type TaskRepository struct {
}

func NewTaskRepository() TaskRepository {
    return TaskRepository{}
}



// データベースから検索する
func (m TaskRepository) Retrieve(con_callsign string, con_fr string, con_to string, f_chk int ) *[]QSLstr {

  var qsls []QSLstr
  var qslv QSLstr
  const t_fmt = "20060102"
  // データベースのコネクションを開く
  db, err := sql.Open("sqlite3", "./qsldb/qsldb_main.db")
  defer db.Close()
  checkErr(err)

  rows, err := db.Query(
      `select * from QSLCARDS order by CALLSIGN; `)

  var i_recno int = 0
  var id int
  var callsign string
  var files string
  datetime := time.Now()
  for rows.Next() {      ///  検索結果をオブジェクトに追加　複数行対応
    err = rows.Scan(&id, &callsign, &datetime, &files)
    checkErr(err)

    if f_chk == 0 {   // search by callsign
      if callsign == strings.ToUpper(con_callsign)  {
        if env.S_mode() { log.Printf("id: %d, callsign: %s, datetime: %v, files: %s\n", id, callsign, datetime, files) }
        qslv.ID = id
        qslv.Callsign = callsign
        qslv.Datetime = datetime.String()
        qslv.Files = files
        qsls = append(qsls, qslv)
        i_recno = i_recno + 1
      }
    } else if f_chk == 1 {   // serch by period
      //

      t_fr, _ := time.Parse(t_fmt, con_fr)
      t_to, _ := time.Parse(t_fmt, con_to)

      if  (datetime.Unix() >= t_fr.Unix()) && (datetime.Unix() <= t_to.Unix()) {
        if env.S_mode() { log.Printf("id: %d, callsign: %s, datetime: %v, files: %s\n", id, callsign, datetime, files) }
        qslv.ID = id
        qslv.Callsign = callsign
        qslv.Datetime = datetime.String()
        qslv.Files = files
        qsls = append(qsls, qslv)
        i_recno = i_recno + 1
      }
    }
  }

  if i_recno == 0 {    //  データが見つからなかった場合
    qslv.ID = 0
    qslv.Callsign = "No callsign found"
    qslv.Datetime = ""
    qslv.Files = ""
    qsls = append(qsls, qslv)
    if env.S_mode() { log.Printf("Not found record") }
  }

  return &qsls

}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

// データベースから検索する
func (m TaskRepository) InsertUpload(con_callsign string, con_file string ) error {

  const t_fmt = "20060102"
  // データベースのコネクションを開く
  db, err := sql.Open("sqlite3", "./qsldb/qsldb_main.db")
  defer db.Close()
    if err != nil { return err }   // エラー時は、エラーを返却して抜ける

  // 文字列作成
  i_timeu := time.Now().Unix()
  d_time := time.Now()

  // データの挿入
  res, err := db.Exec(
    `INSERT INTO UPLOAD_FILES (ID, CALLSIGN, DATETIME, FILES) VALUES (?, ?, ?, ?)`,
    i_timeu, strings.ToUpper(con_callsign), d_time, con_file)
    if err != nil { return err }   // エラー時は、エラーを返却して抜ける

  // 挿入処理の結果からIDを取得
  id, err := res.LastInsertId()
    if err != nil { return err }   // エラー時は、エラーを返却して抜ける

    if env.S_mode() { log.Printf("id is %v \n", id) }

  return nil

}

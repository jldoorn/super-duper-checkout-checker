package room

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Bed struct {
	Id       int
	Room     int
	Wing     string
	Name     string
	IsOut    bool
	TimeOut  int
	Ra1      int
	Ra2      int
	Comments string
}

type Room struct {
	Beds []Bed
}

func GetRoom(num int, db *sql.DB) Room {
	rows, err := db.Query("SELECT * FROM bed WHERE room_num = $1", num)
	if err != nil {
		println("err in db query")
	}
	defer rows.Close()
	beds := []Bed{}
	for rows.Next() {
		var id int
		var room_num int
		var name string
		var ra1 int
		var ra2 int
		var isout bool
		var timeout int
		var comments string
		var wing string
		err = rows.Scan(&id, &room_num, &wing, &name, &ra1, &ra2, &isout, &timeout, &comments)
		if err != nil {
			log.Fatal(err)
		}
		beds = append(beds, Bed{Id: id, Room: room_num, Name: name, Ra1: ra1, Ra2: ra2, IsOut: isout, TimeOut: timeout, Wing: wing, Comments: comments})
	}
	return Room{Beds: beds}
}

func UpdateBed(bed Bed, db *sql.DB) {
	var curTime int64
	if bed.IsOut {
		curTime = time.Now().Unix()
	} else {
		curTime = int64(bed.TimeOut)
	}
	//	db.Exec(`
	//UPDATE bed
	//SET bed_checked_out = $1, ra1_id = $2, ra2_id = $3, bed_time_checked_out = $4, comments = $5
	//WHERE id = $5;
	//`, bed.IsOut, bed.Ra1, bed.Ra2, curTime, bed.Comments)
	fmt.Println(bed)
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("update bed set bed_checked_out=?, ra1_id=?, ra2_id=?, bed_time_checked_out=?, comments=? where id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec(bed.IsOut, bed.Ra1, bed.Ra2, curTime, bed.Comments, bed.Id)
	tx.Commit()
}

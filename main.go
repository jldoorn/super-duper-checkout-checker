package main

import (
	"database/sql"
	"fmt"
	"github.com/jldoorn/super-duper-checkout-checker/room"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	println("hello world")
	db, err := sql.Open("sqlite3", "./beds.db")
	if err != nil {
		log.Fatal(err)
	}
	myBed := room.GetRoom(2, db)
	fmt.Println(myBed)
	//myBed.IsOut = true
	//myBed.Ra1 = 1
	//myBed.Comments = "Room looks good!"
	//
	//room.UpdateBed(myBed, db)
	//fmt.Println(room.GetRoom(2, db).Beds[0])
	//room.GetRoom(237, db)
	//controllers.RegisterControllers()
	//http.ListenAndServe(":3000", nil)
	//
}

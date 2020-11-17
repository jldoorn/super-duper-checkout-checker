package main

import (
	"github.com/jldoorn/super-duper-checkout-checker/controllers"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

func main() {
	println("hello world")

	controllers.RegisterControllers()
	http.ListenAndServe(":3080", nil)
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

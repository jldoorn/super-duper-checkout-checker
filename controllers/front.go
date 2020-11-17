package controllers

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func RegisterControllers() {
	bedDB, err := sql.Open("sqlite3", "./beds.db")
	if err != nil {
		log.Fatal(err)
	}
	rc := newRoomCheckController(bedDB)

	http.Handle("/api/room", rc)
	http.Handle("/api/room/", rc)
}

func enableCors(w *http.ResponseWriter){
	(*w).Header().Set("Access-Control-Allow-Origin","*")
}

func encodeResponseAsJson(data interface{}, w io.Writer){
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
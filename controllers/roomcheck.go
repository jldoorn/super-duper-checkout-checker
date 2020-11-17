package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/jldoorn/super-duper-checkout-checker/room"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type roomCheckController struct {
	db *sql.DB
	roomIDPattern *regexp.Regexp
}

type ApiRequest struct {
	RequestType int // 1 = room lookup 2 = floor lookup
	RequestBody string
}

type roomResponse struct {
	room room.Room
}

func (rr * roomCheckController) ServeHTTP(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	if r.URL.Path == "/api/room" {
		switch r.Method {
		case http.MethodGet:
			rr.getAll(w)
		// case http.MethodPost:
		// 	rr.post(w, r)
		}
	} else {
		matches := rr.roomIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0{
			w.WriteHeader(http.StatusNotFound)
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil{
			w.WriteHeader(http.StatusNotFound)
		}
		switch r.Method {
		case http.MethodGet:
			rr.get(id, w)
		}
	}
}

func (rr*roomCheckController) getAll(w http.ResponseWriter){
	encodeResponseAsJson(room.GetAll(rr.db), w)
}

func (rr * roomCheckController) get(roomNum int, w http.ResponseWriter){
	encodeResponseAsJson(room.GetRoom(roomNum, rr.db), w)
}

func (rr * roomCheckController) parseRequest(w http.ResponseWriter, r * http.Request){
	dec := json.NewDecoder(r.Body)
	var ar ApiRequest
	err := dec.Decode(&ar)
	if err != nil{
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if ar.RequestBody == "" || ar.RequestType < 1{
		log.Println("Bad request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch ar.RequestType{
	case 1:
		rsp,_ := strconv.Atoi(ar.RequestBody)
		encodeResponseAsJson(room.GetRoom(rsp, rr.db), w)
	case 2:
		encodeResponseAsJson(room.GetFloor(ar.RequestBody, rr.db), w)
	}

}

func (rr * roomCheckController) post(w http.ResponseWriter, r * http.Request){
	dec := json.NewDecoder(r.Body)
	var bed room.Bed
	dec.Decode(&bed)
	room.UpdateBed(bed, rr.db)
	encodeResponseAsJson(room.GetRoom(bed.Room, rr.db), w)
}

func newRoomCheckController(db *sql.DB) *roomCheckController{
	return &roomCheckController{db: db, roomIDPattern: regexp.MustCompile(`^/api/room/(\d+)/?`)}
}
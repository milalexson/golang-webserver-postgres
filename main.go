package main

import (
	"fmt"
	"net/http"
	"log"
	"strconv"
	_ "github.com/lib/pq"
	"database/sql"
)


const (  
  host     = "vhapoc.c6d1wvruqzfl.ap-southeast-2.rds.amazonaws.com"
  port     = 5432
  user     = "apcera"
  password = "apc3raPoC"
  dbname   = "postgres"
)


var count int64=1
var cdrs string = "<span></span>"
var MSIDN int64
var IMSI int64
var IMEI int64
var PLAN string
var CALL_TYPE string
var CORRESP_TYPE string
var CORRESP_ISDN int64
var DURATION int64
var TIME string
var DATE string





func Log(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
    })
}

func getCDRS(){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	    "password=%s dbname=%s sslmode=disable",
	    host, port, user, password, dbname)
	  db, err := sql.Open("postgres", psqlInfo)
	  if err != nil {
	    panic(err)
	  }
	  defer db.Close()

	  err = db.Ping()
	  if err != nil {
	    panic(err)
	  }

	//fmt.Println("Successfully connected!")

	sql := "SELECT MSIDN,IMSI,IMEI,PLAN,CALL_TYPE,CORRESP_TYPE,CORRESP_ISDN,DURATION,TIME,DATE from cdr"
	rows, err := db.Query(sql)

	cdrs ="<table border='1' cellpadding='10'><tr style='font-weight:600;'><td>MSIDN</td><td>IMSI</td><td>IMEI</td><td>PLAN</td><td>CALL_TYPE</td><td>CORRESP_TYPE</td><td>CORRESP_ISDN</td><td>DURATION</td><td>TIME</td><td>DATE</td></tr>"

	for rows.Next() {
		if err := rows.Scan(&MSIDN, &IMSI, &IMEI, &PLAN, &CALL_TYPE, &CORRESP_TYPE, &CORRESP_ISDN, &DURATION, &TIME, &DATE); err != nil {
			log.Fatal(err)
		}
		cdrs +="<tr><td>"
		cdrs +=strconv.FormatInt(MSIDN,10)
		cdrs +="</td><td>"
		cdrs +=strconv.FormatInt(IMSI,10)
		cdrs +="</td><td>"
		cdrs +=strconv.FormatInt(IMEI,10)
		cdrs +="</td><td>"
		cdrs +=PLAN
		cdrs +="</td><td>"
		cdrs +=CALL_TYPE
		cdrs +="</td><td>"
		cdrs +=CORRESP_TYPE
		cdrs +="</td><td>"
		cdrs +=strconv.FormatInt(CORRESP_ISDN,10)
		cdrs +="</td><td>"
		cdrs +=strconv.FormatInt(DURATION,10)
		cdrs +="</td><td>"
		cdrs +=TIME
		cdrs +="</td><td>"
		cdrs +=DATE
		cdrs +="</td>"
		cdrs +="</tr>"
		
	}

}

func main() {

	//http.Handle("/", http.FileServer(http.Dir("./docroot")))


   // http.HandleFunc("/count", handler)
   http.HandleFunc("/", handler)

	getCDRS()
	
	error := http.ListenAndServe(":8080", Log(http.DefaultServeMux))
        if error != nil {
                fmt.Println("Error listening on port 8080: ", error)
        }

}

func handler(response http.ResponseWriter, request *http.Request){
	html :=string("<!DOCTYPE html><html><body style='background-color:#fefefe;'><p /><center><span style='font-size:7vh;  color:#888888;'>CDRS</span></center><center>")
    //html +=strconv.FormatInt(count,10)
    html +=cdrs
    html +="</center></body></html>"
    fmt.Fprintf(response, html) 
    count++
    getCDRS()
 
}


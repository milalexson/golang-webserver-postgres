package main

import (
	"fmt"
	"net/http"
	"log"
	"strconv"
)

var count int64=1

func Log(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
    })
}

func main() {

	//http.Handle("/", http.FileServer(http.Dir("./docroot")))


   // http.HandleFunc("/count", handler)
   http.HandleFunc("/", handler)

	err := http.ListenAndServe(":8080", Log(http.DefaultServeMux))
        if err != nil {
                fmt.Println("Error listening on port 8080: ", err)
        }
	
}

func handler(response http.ResponseWriter, request *http.Request){
	html :=string("<!DOCTYPE html><html><body style='background-color:#fefefe;'><p /><center><span style='font-size:10vh;  color:#888888;'>Count</span></center><center><span style='font-size:20vh;  color:#444444;'>")
    html +=strconv.FormatInt(count,10)
    html +="</span></center></body></html>"
    fmt.Fprintf(response, html) 
    count++
 
}


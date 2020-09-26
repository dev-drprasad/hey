package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var lt time.Time

func init() {
	lt = time.Now()
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("request came. last request came before : ", time.Now().Sub(lt))
	lt = time.Now()
	fmt.Fprintf(w, "Hi there")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

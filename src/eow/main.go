package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
	"strings"
	"strconv"
)


func main() {
	tts_init()

	send_testpage(tts_engines())
	
	db_open();
	finalHandler := http.HandlerFunc(final)

	logFile, err := os.OpenFile("server.log",
		os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	http.Handle("/css/", http.StripPrefix("/css",http.FileServer(http.Dir("css"))))
	http.Handle("/images/", http.StripPrefix("/images",http.FileServer(http.Dir("images"))))
	http.Handle("/waves/", http.StripPrefix("/waves",http.FileServer(http.Dir("waves"))))
	http.Handle("/", handlers.LoggingHandler(logFile, finalHandler))
	http.ListenAndServe(":3000", nil)
}

func final(w http.ResponseWriter, r *http.Request) {
	var qid int
	query_string, ok := r.URL.Query()["q"]
	if ok {
		id, err := strconv.Atoi( query_string[0] )
		if  err == nil {
			qid = id
		} else {
			qid = -1
		}
	} else {
		qid = -1;
	}
	// GET + no query string => redirect to GET + generated_query_string
	if  r.Method == "GET" && qid == -1 {
		qid = db_next()
		http.Redirect(w, r, "?q=" + strconv.Itoa(qid), 307)
		return
	}
	// GET + query_string => ( pull from db, update form, make audio )
	// POST + query_string => ( get from form, make audio )
	if r.Method == "GET" {
		send_newpage(w,tts_engines())
	} else {
		r.ParseForm()
		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
		}

	}
}


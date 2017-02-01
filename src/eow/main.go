package main

import (
	//"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	tts_init()

	//send_testpage(tts_engines())

	db_open()

	finalHandler := http.HandlerFunc(final)

	logFile, err := os.OpenFile("server.log",
		os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css"))))
	http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("images"))))
	http.Handle("/waves/", http.StripPrefix("/waves", http.FileServer(http.Dir("waves"))))
	http.Handle("/", handlers.LoggingHandler(logFile, finalHandler))
	http.ListenAndServe(":3000", nil)
}

func parse_form(r *http.Request) (invocation, bool) {
	var inv invocation
	var new_id bool

	r.ParseForm()

	for k, v := range r.Form {
		val := strings.Join(v, "")
		switch k {
		case "share":
			new_id = true
		case "speak":
			new_id = false
		case "espeakVersion":
			inv.TTSVer = val
		case "langVariant":
			inv.Language = val
		case "inputText":
			inv.TextInput = val
			//case "langDictList":
			//case "langDictRules":
			//case "langLangPhoneme":
		}
	}

	return inv, new_id
}

func final(w http.ResponseWriter, r *http.Request) {
	var qid int
	query_string, ok := r.URL.Query()["q"]
	if ok {
		id, err := strconv.Atoi(query_string[0])
		if err == nil {
			qid = id
		} else {
			qid = -1
		}
	} else {
		qid = -1
	}
	//fmt.Println(r.Method, qid)
	// GET + no query string => redirect to GET + generated_query_string
	if qid == -1 {
		inv := db_new()
		http.Redirect(w, r, "?q="+strconv.Itoa(inv.Id), http.StatusFound)
		return
	}
	// GET + query_string => ( pull from db, update form, make audio )
	if r.Method == "GET" {
		inv := db_get(qid)
		//fmt.Println(inv)
		tts_run(&inv)
		send_newpage(w, tts_engines(), inv)
		return
	}
	// POST + query_string => save  and redirect
	if r.Method == "POST" {
		inv, new_id := parse_form(r)
		if new_id {
			inv.Wavepath = ""
			inv.Analysis = ""
			db_clone(&inv)
			http.Redirect(w, r, "?q="+strconv.Itoa(inv.Id), http.StatusFound)
		} else {
			inv.Id = qid
			db_update(inv)
			http.Redirect(w, r, "?q="+strconv.Itoa(inv.Id), http.StatusFound)
		}
		return
	}
	panic(r)
}

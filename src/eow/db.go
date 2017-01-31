package main

import (
	"github.com/boltdb/bolt"
)

var gdb *bolt.DB


type model struct {
	id  int64		// -1 if not present
	

}



func db_open() {
	db, err := bolt.Open("eow.db", 0600, nil)
	if err != nil {
		panic(err)
	}
	gdb = db
}


func db_next() int {

	return 123
}

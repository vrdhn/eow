package main

import (
	"github.com/asdine/storm"
)

var gdb *storm.DB

func db_open() {
	db, _ := storm.Open("eow.db")
	gdb = db
}

func db_new() invocation {
	inv := invocation{}
	gdb.Save(&inv)
	return inv
}

func db_get(id int) invocation {
	var inv invocation
	gdb.One("Id", id, &inv)
	return inv
}

func db_clone(inv *invocation) {
	inv.Id = 0
	gdb.Save(inv)
}

func db_update(inv invocation) {
	gdb.Update(&inv)

}

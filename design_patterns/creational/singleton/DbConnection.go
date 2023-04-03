package main

import (
	"fmt"
	"sync"
	"time"
)

var db *Database
var lock sync.Mutex

type Database struct {
}

func (d *Database) CreateSingleConnection() {
	fmt.Println("Connecting to server...")
	time.Sleep(2 * time.Second)
	fmt.Println("Connected Succesfully")
}

func GetDataBaseInstance() *Database {
	defer lock.Unlock()
	lock.Lock()
	if db == nil {
		fmt.Println("Creating DB Connection")
		db = &Database{}
		db.CreateSingleConnection()
	} else {
		fmt.Println("DB Already Connected")
	}
	return db
}

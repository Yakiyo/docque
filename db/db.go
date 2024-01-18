package db

import "github.com/ostafen/clover/v2"

var cdb *clover.DB

// initialize the database
func Init() error {
	var err error
	cdb, err = clover.Open("data.db")
	return err
}

func Close() error {
	return cdb.Close()
}

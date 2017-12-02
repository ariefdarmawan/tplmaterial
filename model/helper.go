package model

import (
	"github.com/eaciit/dbflex"
	. "github.com/eaciit/dbflex/drivers/mongodb"
	"github.com/eaciit/toolkit"
)

var connectionString string

func SetConnectinString(txt string) error {
	connectionString = txt
	c, e := Connection()
	if c != nil {
		defer c.Close()
	}

	if e != nil {
		return e
	}
}

func Connection() (dbflex.IConnection, error) {
	if connectionString == "" {
		return toolkit.Errorf("connectionString is not yet setup")
	}

	conn := dbflex.NewConnectionFromUri("mongodb", connectionString)
	if err := conn.Connect(); err != nil {
		return err
	}

	return conn
}

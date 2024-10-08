package es

import "github.com/acronis/perfkit/db"

func (g *esGateway) Exec(format string, args ...interface{}) (db.Result, error) {
	return nil, nil
}

func (g *esGateway) QueryRow(format string, args ...interface{}) db.Row {
	return &db.EmptyRows{}
}

func (g *esGateway) Query(format string, args ...interface{}) (db.Rows, error) {
	return &db.EmptyRows{}, nil
}

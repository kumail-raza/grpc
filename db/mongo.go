package db

import (
	"github.com/globalsign/mgo"
)

type Mongo struct {
	Session *mgo.Session
	DBName  string
}

//NewMongoDB NewMongoDB
func NewMongoDB(connStr string, dbName string) (*Mongo, error) {

	session, err := mgo.Dial(connStr + "/" + dbName)
	if err != nil {
		return nil, err
	}
	return &Mongo{Session: session, DBName: dbName}, nil
}

//GetCollection GetCollection
func (m *Mongo) GetCollection(colName string) *mgo.Collection {
	return m.Session.DB(m.DBName).C(colName)
}

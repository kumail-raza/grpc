package db

import "github.com/mongodb/mongo-go-driver/mongo"

type Mongo struct {
	Client *mongo.Client
}

//NewMongoDB NewMongoDB
func NewMongoDB(connStr string) (*Mongo, error) {

	client, err := mongo.NewClient(connStr)
	if err != nil {
		return nil, err
	}
	return &Mongo{Client: client}, nil
}

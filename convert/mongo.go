package convert

import (
	"github.com/globalsign/mgo/bson"
)

//MongoIDToStringPtr MongoIDToStringPtr
func MongoIDToStringPtr(objID bson.ObjectId) *string {

	str := objID.Hex()
	return &str
}

//StringPtrToMongoID StringPtrToMongoID
func StringPtrToMongoID(str *string) bson.ObjectId {
	return bson.ObjectIdHex(*str)
}

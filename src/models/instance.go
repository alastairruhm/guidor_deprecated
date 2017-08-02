package models

import (
	"time"

	"github.com/alastairruhm/guidor/src/service/mongo"
	"github.com/alastairruhm/guidor/src/util"
	"gopkg.in/mgo.v2/bson"
)

// Instance Model
type Instance struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Token     string        `json:"token" bson:"token"`
	IP        string        `json:"ip" bson:"ip"`
	Hostname  string        `json:"hostname" bson:"hostname"`
	DbType    string        `json:"db_type" bson:"db_type"`
	DbVersion string        `json:"db_version" bson:"db_version"`
	DbName    string        `json:"db_name" bson:"db_name"`
	CreatedAt time.Time     `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at,omitempty" bson:"updated_at"`
}

func newInstanceCollection() *mongo.Collection {
	return mongo.NewCollectionSession("instances")
}

// CreateInstance post instance
func CreateInstance(instance Instance) (Instance, error) {
	// Get post collection connection
	c := newInstanceCollection()
	defer c.Close()

	// set default mongodb ID  and created date
	instance.ID = bson.NewObjectId()
	// instance.Token = bson.NewObjectId().Hex()
	instance.Token = util.GenerateToken()
	instance.CreatedAt = time.Now()
	// Insert post to mongodb
	err := c.Session.Insert(&instance)
	if err != nil {
		return instance, err
	}
	return instance, nil
}

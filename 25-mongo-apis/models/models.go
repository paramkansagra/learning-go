package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
We are giving bson id because mongo db gives us an id on its own
also _id means that it is private most probably

bson is a bit different than json as it adds more functionality over json
*/
type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}

package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ToDo struct {
    ID primitive.ObjectID `bson:"_id,omitempty"`
    Title string `bson:"title,omitempty"`
    Description string `bson:"desc,omitempty"`
    CreatedBy string `bson:"by,omitempty"`
    CreatedOn int64 `bson:"on,emitempty"`
}

type ToDos struct {
    ToDos []ToDo `json: ToDos`
}
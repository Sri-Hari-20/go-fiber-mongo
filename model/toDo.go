package model

import (
	"go.mongodb.org/mongo-driver/bson"
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
    ToDos []ToDo `json: "todos"`
}

func PrepareBsonUpdateTodo (newToDo ToDo) bson.M {
    bsonData := bson.M{}

    if newToDo.Title != "" {
        bsonData["title"] = newToDo.Title
    }

    if newToDo.Description != "" {
        bsonData["desc"] = newToDo.Description
    }

    preparedBson := bson.M{"$set" : bsonData}

    return preparedBson
}
package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Group struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title"`
	TeacherId uint32             `bson:"teacher_id"`
	Students  []uint32           `bson:"students"`
}

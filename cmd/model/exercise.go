package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Exercise struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	Group_id     primitive.ObjectID `bson:"group_id"`
	TeacherId    uint32             `bson:"teacher_id"`
	StudentId    uint32             `bson:"student_id"`
	StatId       string             `bson:"stat_id"`
	ExerciseName string             `bson:"exercise_name"`
	Words        []uint32           `bson:"words"`
}

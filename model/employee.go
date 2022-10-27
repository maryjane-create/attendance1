package model

import "time"

type Employee struct {
	Firstname string `json:"firstname"bson:"firstname"`
	Lastname string  `json:"lastname" bson:"lastname"`
	time time.Time  `json:"time" bson:"time"`
}

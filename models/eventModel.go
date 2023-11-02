package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	Id          primitive.ObjectID `bson:"_id"`
	Event_id    string             `json:"event_id"`
	Event       string             `json:"event" validate:"required"`
	Start       time.Time          `json:"start" validate:"required"`
	End         time.Time          `json:"end" validate:"required"`
	Description string             `json:"description"`
	Trainer     string             `json:"trainer"`
	Places      string                `json:"places" validate:"required"`
	Visitors    int                `json:"visitors"`
	Day         string             `json:"day"`
	Company_id  string             `json:"company_id"`
	Admin_id     string             `json:"admin_id"`
}

package dao_mongo

import "gopkg.in/mgo.v2"

type EventDaoMongo struct {
	mongodb *mgo.Database
}

func (e *EventDaoMongo) AddEvent(eventType string) bool {
	return true
}

func (e *EventDaoMongo) CountByType(eventType string) int {
	return 1
}

func NewEventDao(mongoDb *mgo.Database) *EventDaoMongo {
	return &EventDaoMongo{

	}
}
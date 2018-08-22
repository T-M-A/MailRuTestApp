package dao_mongo

import "gopkg.in/mgo.v2"

type EventDaoMongo struct {
	mongodb *mgo.Database
}

func (e *EventDaoMongo) collection() (*mgo.Collection) {
	return e.mongodb.C("routes")
}

func (e *EventDaoMongo) AddEvent(eventType string) error {
	row := make(map[string]string)
	row["type"] = eventType

	err := e.collection().Insert(row)
	return err
}

func (e *EventDaoMongo) CountByType(eventType string) (int, error) {
	query := make(map[string]string)
	query["type"] = eventType

	count, err := e.collection().Find(query).Count()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func NewEventDao(mongoConnect *mgo.Database) *EventDaoMongo {
	return &EventDaoMongo{
		mongoConnect,
	}
}
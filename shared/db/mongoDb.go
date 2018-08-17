package db

import (
	"gopkg.in/mgo.v2"
	"errors"
	"strconv"
)

func NewMongoDriver(host string, port int, login string, password string, database string) (*mgo.Database, error) {
	session, err := mgo.Dial(host+":"+strconv.Itoa(port))

	if err != nil {
		return nil, errors.New("can't connect")
	}

	if login != "" {
		cred := &mgo.Credential{
			Username:  login,
			Password:  password,
			Mechanism: "SCRAM-SHA-1",
			Source:    database,
		}

		err = session.Login(cred)

		if err != nil {
			return nil, errors.New("auth error")
		}
	}

	return session.DB(database), nil
}


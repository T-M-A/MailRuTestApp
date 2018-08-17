package application

import (
	"net/http"
	"strconv"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"gopkg.in/mgo.v2"
	"github.com/ruelephant/MailRuTestApp/shared/db"
)

type config struct {
	Database struct {
		Type string `json:"type"`
	} `json:"db"`
	HTTP struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"http"`
	Mongodb struct {
		Host string `json:"host"`
		Port int    `json:"port"`
		Login string `json:"login"`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:"mongodb"`
}

type application struct {
	conf config
	mongoConnect *mgo.Database
}

func (a *application) Http(handle http.Handler) error {
	return http.ListenAndServe(a.conf.HTTP.Host+":"+strconv.Itoa(a.conf.HTTP.Port), handle)
}

func (a *application) MongoDb() (*mgo.Database, error) {
	if a.mongoConnect != nil {
		return a.mongoConnect, nil
	}

	mongoConnect, err := db.NewMongoDriver(a.conf.Mongodb.Host, a.conf.Mongodb.Port, a.conf.Mongodb.Login, a.conf.Mongodb.Password, a.conf.Mongodb.Database)
	if err != nil {
		return nil, err
	}
	a.mongoConnect = mongoConnect
	return a.mongoConnect, nil
}

func NewApplication(configPath string, logger log.Logger) (*application, error) {
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		//Do something
	}

	var conf config

	err = yaml.Unmarshal(content, &conf)
	if err != nil {
		return nil, err
	}

	return &application{
		conf: conf,
	}, nil
}

package main

import (
	"github.com/ruelephant/MailRuTestApp/shared/application"
	"github.com/julienschmidt/httprouter"
	"github.com/ruelephant/MailRuTestApp/shared/handlers"
	"github.com/ruelephant/MailRuTestApp/shared/dao_mongo"
	"log"
)

func main() {
	app, err := application.NewApplication("config.yaml", log.Logger{})
	if err != nil {
		log.Fatal("app: ", err)
	}

	mongo, err := app.MongoDb()
	if err != nil {
		log.Fatal("mongo: ", err)
	}

	dao := dao_mongo.NewEventDao(mongo)

	r := httprouter.New()
	r.Handler("POST", "/events", handlers.NewEventHandler(dao))
	r.Handler("GET", "/stat", handlers.NewStatHandler(dao))

	err = app.Http(r)

	if err != nil {
		log.Fatal(err)
	}
}

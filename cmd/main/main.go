package main

import (
	"github.com/ruelephant/MailRuTestApp/shared/application"
	"github.com/julienschmidt/httprouter"
	"github.com/ruelephant/MailRuTestApp/shared/handlers"
	"github.com/ruelephant/MailRuTestApp/shared/dao_mongo"
	"log"
	"github.com/ruelephant/MailRuTestApp/shared/handlers/limiter"
	"github.com/ruelephant/MailRuTestApp/shared/handlers/responce"
)

func main() {
	app, err := application.NewApplication("./config.yaml", log.Logger{})
	if err != nil {
		log.Fatal("app: ", err)
	}

	mongo, err := app.MongoDb()
	if err != nil {
		log.Fatal("mongo: ", err)
	}

	dao := dao_mongo.NewEventDao(mongo)
	response := &responce.JsonResponse{}

	middleware, err := limiter.GetLimitMiddleware()
	if err != nil {
		log.Fatal("middleware: ", err)
	}

	r := httprouter.New()
	r.Handler("POST", "/events", middleware.Handler(handlers.NewEventHandler(dao, response)))
	r.Handler("GET", "/stat", middleware.Handler(handlers.NewStatHandler(dao, response)))

	err = app.Http(r)

	if err != nil {
		log.Fatal(err)
	}
}

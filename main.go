package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

	_planetsDeliveryHttp "github.com/jsperandio/b2w-star-wars/planet/delivery/http"
	_planetsDeliveryHttpClient "github.com/jsperandio/b2w-star-wars/planet/delivery/http/client"
	_planetsDeliveryHttpMiddleware "github.com/jsperandio/b2w-star-wars/planet/delivery/http/middleware"
	_planetRepository "github.com/jsperandio/b2w-star-wars/planet/repository/mongodb"
	_planetUsecase "github.com/jsperandio/b2w-star-wars/planet/usecase"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Setup Fiber Middlewares
	_planetsDeliveryHttpMiddleware.InitFiberMiddleware(app)

	// // Init db for app
	mngdbrep, err := _planetRepository.NewMongoAppRepository(os.Getenv("MONGODB_CONNECTION_URL"))
	if err != nil {
		log.Fatalf("error creating Mongo Client: %v", err)
	}

	// Init new Http Client
	http_client := _planetsDeliveryHttpClient.NewRESTClient(
		"https://swapi.dev/api/",
		3,
		5,
		20)

	// Setup Swapi Consumer
	swapi := _planetsDeliveryHttpClient.NewSwapi(http_client)

	// // Init Planet Repository Mongo
	prepo := _planetRepository.NewMongoDbPlanetRepository(mngdbrep.DB)

	// // Init Planet Usecase
	pucase := _planetUsecase.NewPlanetUsecase(prepo)

	// // Routes
	_planetsDeliveryHttp.NewPlanetHandler(app, pucase, swapi)

	// // Start server
	log.Fatal(app.Listen(":3000"))
}

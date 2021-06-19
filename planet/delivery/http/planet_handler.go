package http

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/jsperandio/b2w-star-wars/domain"
	_client "github.com/jsperandio/b2w-star-wars/planet/delivery/http/client"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// PlanetHandler  represent the httphandler for Planet
type PlanetHandler struct {
	PUsecase domain.PlanetUsecase
	SWapi    _client.SwapiClient
}

// NewPlanetHandler will initialize the Planets/ resources endpoint
func NewPlanetHandler(a *fiber.App, us domain.PlanetUsecase, swapi _client.SwapiClient) {
	handler := &PlanetHandler{
		PUsecase: us,
		SWapi:    swapi,
	}

	route := a.Group("/api/v1")

	route.Post("/planet", handler.Store)
	route.Get("/planets", handler.FindAll)
	route.Get("/planet", handler.GetByName)
	route.Get("/planet/:id", handler.GetByID)
	route.Delete("/planet/:id", handler.Delete)
}

// FindAll will get all Planets stored in database
func (p *PlanetHandler) FindAll(c *fiber.Ctx) error {
	// var plnt *domain.Planet
	// var eg errgroup.Group

	// eg.Go(func() error {
	// 	// Fetch from Database.
	// 	p, err := p.PUsecase.GetByID(idP)
	// 	if err == nil {
	// 		plnt = p
	// 	}
	// 	return err
	// })

	// eg.Go(func() error {
	// 	// Fetch from Swapi.
	// 	swapi_plnt, err := p.SWapi.GetPlanetByName(plnt.Name)
	// 	if err == nil {
	// 		plnt.Appearances = len(swapi_plnt.Films)
	// 	}
	// 	return err
	// })

	// // Wait for all HTTP fetches to complete.
	// if err := eg.Wait(); err == nil {
	// 	fmt.Println("Successfully fetched all URLs.")
	// }

	plnts, err := p.PUsecase.FindAll()

	if err != nil {
		return c.Status(getStatusCode(err)).JSON(ResponseError{Message: err.Error()})
	}

	// swapi_plnt, err := p.SWapi.GetPlanetByName(gname)

	// if err != nil {
	// 	return c.Status(getStatusCode(err)).JSON(ResponseError{Message: err.Error()})
	// }

	// plnt.Appearances = len(swapi_plnt.Films)

	return c.Status(fiber.StatusOK).JSON(&plnts)
}

// GetByName will get Planet by given param name
func (p *PlanetHandler) GetByName(c *fiber.Ctx) error {

	gname := c.Query("name")

	if gname == "" {
		return c.Status(fiber.StatusNotFound).JSON(domain.ErrNotFound.Error())
	}

	// var plnt *domain.Planet
	// var eg errgroup.Group

	// eg.Go(func() error {
	// 	// Fetch from Database.
	// 	p, err := p.PUsecase.GetByID(idP)
	// 	if err == nil {
	// 		plnt = p
	// 	}
	// 	return err
	// })

	// eg.Go(func() error {
	// 	// Fetch from Swapi.
	// 	swapi_plnt, err := p.SWapi.GetPlanetByName(plnt.Name)
	// 	if err == nil {
	// 		plnt.Appearances = len(swapi_plnt.Films)
	// 	}
	// 	return err
	// })

	// // Wait for all HTTP fetches to complete.
	// if err := eg.Wait(); err == nil {
	// 	fmt.Println("Successfully fetched all URLs.")
	// }

	plnt, err := p.PUsecase.GetByName(gname)

	if err != nil {
		return c.Status(getStatusCode(err)).JSON(ResponseError{Message: err.Error()})
	}

	swapi_plnt, err := p.SWapi.GetPlanetByName(gname)

	if err != nil {
		return c.Status(getStatusCode(err)).JSON(ResponseError{Message: err.Error()})
	}

	plnt.Appearances = len(swapi_plnt.Films)

	return c.Status(fiber.StatusOK).JSON(&plnt)
}

// GetByID will get Planet by given id
func (p *PlanetHandler) GetByID(c *fiber.Ctx) error {

	idP := c.Params("id")

	if idP == "" {
		return c.Status(fiber.StatusNotFound).JSON(domain.ErrNotFound.Error())
	}

	// var plnt *domain.Planet
	// var eg errgroup.Group

	// eg.Go(func() error {
	// 	// Fetch from Database.
	// 	p, err := p.PUsecase.GetByID(idP)
	// 	if err == nil {
	// 		plnt = p
	// 	}
	// 	return err
	// })

	// eg.Go(func() error {
	// 	// Fetch from Swapi.
	// 	swapi_plnt, err := p.SWapi.GetPlanetByName(plnt.Name)
	// 	if err == nil {
	// 		plnt.Appearances = len(swapi_plnt.Films)
	// 	}
	// 	return err
	// })

	// // Wait for all HTTP fetches to complete.
	// if err := eg.Wait(); err == nil {
	// 	fmt.Println("Successfully fetched all URLs.")
	// }

	plnt, err := p.PUsecase.GetByID(idP)

	if err != nil {
		return c.Status(getStatusCode(err)).JSON(ResponseError{Message: err.Error()})
	}

	swapi_plnt, err := p.SWapi.GetPlanetByName(plnt.Name)

	if err != nil {
		return c.Status(getStatusCode(err)).JSON(ResponseError{Message: err.Error()})
	}

	if swapi_plnt != nil {
		plnt.Appearances = len(swapi_plnt.Films)
	}

	return c.Status(fiber.StatusOK).JSON(&plnt)
}

func isRequestValid(m *domain.Planet) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Store will store the Planet by given request body
func (p *PlanetHandler) Store(c *fiber.Ctx) (err error) {
	plnt := new(domain.Planet)

	if err := c.BodyParser(plnt); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseError{Message: err.Error()})
	}

	var ok bool
	if ok, err = isRequestValid(plnt); !ok {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseError{Message: err.Error()})
	}

	fmt.Println("entrada ")
	fmt.Println(plnt)

	err = p.PUsecase.Store(plnt)
	if err != nil {
		return c.Status(getStatusCode(err)).JSON(ResponseError{Message: err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON("")
}

// Delete will delete Planet by given param
func (p *PlanetHandler) Delete(c *fiber.Ctx) error {
	idP := c.Params("id")

	if idP == "" {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrBadParamInput.Error())
	}

	err := p.PUsecase.Delete(idP)

	if err != nil {
		return c.Status(getStatusCode(err)).JSON(ResponseError{Message: err.Error()})
	}

	return c.Status(fiber.StatusNoContent).JSON("")
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	case domain.ErrBadBodyInput:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

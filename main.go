package main

import (
	"errors"
	"flyer/sol"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func overall(c echo.Context) (err error) {
	var journey sol.Journey
	if err := c.Bind(&journey); err != nil {
		return ErrInvalidRequest
	}
	sd, err := sol.Overall(journey)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, sd)
	return
}

func itinerary(c echo.Context) (err error) {
	var journey sol.Journey
	if err := c.Bind(&journey); err != nil {
		return ErrInvalidRequest
	}
	it, err := sol.Itinerary(journey)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, it)
	return
}

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Use(errHandler)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.POST("/calculate", overall)
	e.POST("/itinerary", itinerary)
	e.Static("/api", "api")
	e.Logger.Fatal(e.Start(":8080"))
}

var ErrInvalidRequest = errors.New("invalid request")

func errHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err == nil {
			return nil
		}
		switch {
		case errors.Is(err, sol.ErrInvalidFlightData), errors.Is(err, ErrInvalidRequest):
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return err
	}
}

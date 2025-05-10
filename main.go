package main

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type FlightTicketsPayload [][]string

func main() {
	// start new echo instance
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// routes
	e.POST("/tickets/order", orderFlights)

	// start server
	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}

}

func orderFlights(ctx echo.Context) error {

	var tickets FlightTicketsPayload
	if err := ctx.Bind(&tickets); err != nil || len(tickets) == 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid payload"})
	}

	itinerary := orderItinerary(tickets)

	return ctx.JSON(http.StatusOK, itinerary)
}

func orderItinerary(tickets [][]string) []string {

	//check for empty tickets payload
	if len(tickets) == 0 {
		return []string{}
	}

	//create maps to store routes & destinations
	routeMap := make(map[string]string)
	destinationMap := make(map[string]bool)

	//loop through the tickets payload to populate routes & destinations maps
	for _, ticket := range tickets {
		from := ticket[0]
		to := ticket[1]
		routeMap[from] = to
		destinationMap[to] = true
	}

	//check fo start of itinerary based on the source that doesn't appear as a destination
	var start string
	for from := range routeMap {
		if !destinationMap[from] {
			start = from
			break
		}
	}

	// instantiate itinerary beginning with the starting airport
	itinerary := []string{start}
	// loop through the routes while checking if the next route is the current starting
	// route while appending the new route to the itinerary slice, after every iteration
	// update the start airport to the airport just added to the itinerary
	for {
		next, ok := routeMap[start]
		if !ok {
			break
		}
		itinerary = append(itinerary, next)
		start = next
	}

	return itinerary
}

package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestOrderFlights(t *testing.T) {

	flightsPayload := `[["LAX", "DXB"], ["JFK","LAX"],["SFO","SJC"],["DXB","SFO"]]`
	itinerary := `["JFK","LAX","DXB","SFO","SJC"]`

	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(flightsPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, orderFlights(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, itinerary, rec.Body.String())
	}
}

func TestOrderFlights_NoPayload(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, orderFlights(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestOrderItinerary(t *testing.T) {
	testsArray := []struct {
		name     string
		tickets  [][]string
		expected []string
	}{
		{
			name: "Valid Input",
			tickets: [][]string{
				{"LAX", "DXB"},
				{"JFK", "LAX"},
				{"SFO", "SJC"},
				{"DXB", "SFO"},
			},
			expected: []string{"JFK", "LAX", "DXB", "SFO", "SJC"},
		},
		{
			name: "One Input",
			tickets: [][]string{
				{"NBO", "LAX"},
			},
			expected: []string{"NBO", "LAX"},
		},
		{
			name:     "No Input",
			tickets:  [][]string{},
			expected: []string{},
		},
		{
			name: "Longer Input",
			tickets: [][]string{
				{"A", "B"},
				{"C", "D"},
				{"B", "C"},
				{"F", "G"},
				{"E", "F"},
				{"D", "E"},
			},
			expected: []string{"A", "B", "C", "D", "E", "F", "G"},
		},
	}

	for _, test := range testsArray {
		t.Run(test.name, func(t *testing.T) {
			result := orderItinerary(test.tickets)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("got %v, want %v", result, test.expected)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

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
			name: "One Ticket",
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
			},
			expected: []string{"A", "B", "C", "D"},
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

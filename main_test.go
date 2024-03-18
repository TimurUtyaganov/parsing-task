package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser(t *testing.T) {
	assert.New(t)

	testCases := []struct {
		Value    string
		Expected *Cond
	}{
		{
			Value: "@id:22",
			Expected: &Cond{
				Type:      "@",
				Field:     "id",
				Condition: "=",
				Value:     "22",
			},
		},
		{
			Value: "@id:>30",
			Expected: &Cond{
				Type:      "@",
				Field:     "id",
				Condition: ">",
				Value:     "30",
			},
		},
		{
			Value: "@id:>=100",
			Expected: &Cond{
				Type:      "@",
				Field:     "id",
				Condition: ">=",
				Value:     "100",
			},
		},
		{
			Value: "@id:<30",
			Expected: &Cond{
				Type:      "@",
				Field:     "id",
				Condition: "<",
				Value:     "30",
			},
		},
		{
			Value: "@id:<=100",
			Expected: &Cond{
				Type:      "@",
				Field:     "id",
				Condition: "<=",
				Value:     "100",
			},
		},
		{
			Value: "@id:!=100",
			Expected: &Cond{
				Type:      "@",
				Field:     "id",
				Condition: "!=",
				Value:     "100",
			},
		},
		{
			Value: "#office",
			Expected: &Cond{
				Type:      "#",
				Field:     "office",
				Condition: "=",
				Value:     nil,
			},
		},
		{
			Value: "#office:1010",
			Expected: &Cond{
				Type:      "#",
				Field:     "office",
				Condition: "=",
				Value:     "1010",
			},
		},
		{
			Value: "#office:>101010",
			Expected: &Cond{
				Type:      "#",
				Field:     "office",
				Condition: ">",
				Value:     "101010",
			},
		},
		{
			Value: "#office:>=101011",
			Expected: &Cond{
				Type:      "#",
				Field:     "office",
				Condition: ">=",
				Value:     "101011",
			},
		},
		{
			Value: "#office:<101010",
			Expected: &Cond{
				Type:      "#",
				Field:     "office",
				Condition: "<",
				Value:     "101010",
			},
		},
		{
			Value: "#office:<=101011",
			Expected: &Cond{
				Type:      "#",
				Field:     "office",
				Condition: "<=",
				Value:     "101011",
			},
		},
		{
			Value: "#office:!=101010",
			Expected: &Cond{
				Type:      "#",
				Field:     "office",
				Condition: "!=",
				Value:     "101010",
			},
		},
	}

	for _, tc := range testCases {
		actual, err := Parser(tc.Value)
		if !assert.NoError(t, err) {
			continue
		}

		assert.Equal(t, tc.Expected, actual)
	}
}

package main

import (
	"testing"
)

var parseValueTests = []struct {
	argument        string
	expectedResult  interface{}
	testDescription string
}{
	{"true", true, "boolean"},
	{"\"true\"", "true", "boolean as string"},
	{"3.4", 3.4, "number"},
	{"\"3.4\"", "3.4", "number as string"},
}

func TestParseValue(t *testing.T) {
	for _, tt := range parseValueTests {
		assertResultWithContext(t, tt.expectedResult, parseValue(tt.argument), tt.testDescription)
	}
}

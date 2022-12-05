package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCron_ConvertStringCron(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{input: "1 * * * *", expected: SEC},
		{input: "1 1 * * *", expected: SEC + MIN},
		{input: "1 1 1 * *", expected: SEC + MIN + HOUR},
		{input: "1 1 1 1 *", expected: SEC + MIN + HOUR + DAY},
		{input: "1 1 1 1 1", expected: SEC + MIN + HOUR + DAY + YEAR},
		{input: "* * * * *", expected: 0},
		{input: "* 1 * * *", expected: MIN},
		{input: "* * 1 * *", expected: HOUR},
		{input: "* * * 1 *", expected: DAY},
		{input: "* * * * 1", expected: YEAR},
		{input: "* 1", expected: MIN},
		{input: "* * 1", expected: HOUR},
		{input: "* * * 1", expected: DAY},
		{input: "* * * * 1", expected: YEAR},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			cron, err := Cron{}.ConvertStringCron(test.input)
			assert.Nil(t, err)
			assert.Equal(t, test.expected, cron)
		})
	}
}

func TestCron_ConvertStringCronErrorSymbol(t *testing.T) {
	cron, err := Cron{}.ConvertStringCron("c * * * *")
	assert.Error(t, err)
	assert.Equal(t, -1, cron)
}

func TestCron_ConvertStringCronErrorToManySymbols(t *testing.T) {
	cron, err := Cron{}.ConvertStringCron("1 * * * * *")
	assert.Error(t, err)
	assert.Equal(t, -1, cron)
}

const (
	SEC  int = 1
	MIN      = 60
	HOUR     = 3600
	DAY      = 86400
	YEAR     = 31536000
)

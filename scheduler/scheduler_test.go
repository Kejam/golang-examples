package scheduler

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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

type DoUp struct {
}

func (DoUp) Do() {
	fmt.Println("Up at time ", time.Now())
}

type DoDown struct {
}

func (DoDown) Do() {
	fmt.Println("Down at time ", time.Now())
}

func TestAddCron(t *testing.T) {
	isSchedulerStart, errorScheduler := InitTaskScheduler()
	isDoDown, errorDoDown := AddCron("DoDown", 10, DoDown{})
	isDoUp, errorDoUp := AddCron("DoUp", 20, DoUp{})

	assert.True(t, isSchedulerStart)
	assert.Nil(t, errorScheduler)
	assert.True(t, isDoDown)
	assert.Nil(t, errorDoDown)
	assert.True(t, isDoUp)
	assert.Nil(t, errorDoUp)
	time.Sleep(time.Second * time.Duration(61))
	var cronDown Cron
	var cronUp Cron
	for cron, _ := range taskStorage {
		if cron.Name == "DoDown" {
			cronDown = *cron
		}
		if cron.Name == "DoUp" {
			cronUp = *cron
		}
	}
	assert.Equal(t, uint64(6), cronDown.count)
	assert.Equal(t, uint64(3), cronUp.count)
}

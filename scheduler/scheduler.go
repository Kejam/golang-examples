package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Cron struct {
	Name        string
	TimeToStart int // in second
	nextStart   time.Time
	lastStart   time.Time
	count       uint64 // for testing count times, reset after 18446744073709551615 to 0
}

// ConvertStringCron
// Convert simple cron format '*(second) *(minute) *(hour) *(day) *(year)'
// Every argument adds in one sum in second and return second time for scheduler
// Cron can short:
// "1" 		- convert to 1 second
// "2 1 1"  - convert to 3662 second
// "* 1"    - convert to 60 second
func (c Cron) ConvertStringCron(cron string) (int, error) {
	crones := strings.Split(cron, " ")
	var totalTime int
	if len(crones) >= 6 {
		return -1, fmt.Errorf("too many argument %v", cron)
	}
	for i, cronStep := range crones {
		var atoi int
		// Skip * and months
		if cronStep != "*" {
			rawAtoi, err := strconv.Atoi(cronStep)
			if err != nil {
				return -1, fmt.Errorf("Error parse to int %v\n: %v\n", cronStep, err)
			}
			atoi = rawAtoi
		} else {
			continue
		}
		switch i {
		case 0:
			totalTime += atoi
			continue
		case 1:
			totalTime += atoi * 60
			continue
		case 2:
			totalTime += atoi * 60 * 60
			continue
		case 3:
			totalTime += atoi * 60 * 60 * 24
			continue
		case 4:
			totalTime += atoi * 60 * 60 * 24 * 365
			continue
		default:
			return -1, fmt.Errorf("Unknown argument %v with index %v\n", cronStep, i)
		}
	}
	return totalTime, nil
}

type Task interface {
	Do()
}

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

func main() {
	currentStageTime := time.Now()
	var tasks = make(map[*Cron]Task)
	up := Cron{
		Name:        "Up",
		TimeToStart: 10,
		nextStart:   time.Now(),
	}
	tasks[&up] = DoUp{}
	down := Cron{
		Name:        "Down",
		TimeToStart: 20,
		nextStart:   time.Now(),
	}
	tasks[&down] = DoDown{}
	for true {
		for cron, task := range tasks {
			if cron.nextStart.Before(currentStageTime) {
				go task.Do()
				duration := time.Second * time.Duration(cron.TimeToStart)
				cron.nextStart = time.Now().Add(duration)
			}
		}
		currentStageTime = time.Now()
	}
}

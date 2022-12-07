package scheduler

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

var taskStorage = make(map[*Cron]Task)
var initStage bool = false

type Task interface {
	Do()
}

func AddCron(name string, timeToStart int, task Task) (bool, error) {
	for cron, _ := range taskStorage {
		if cron.Name == name {
			return false, fmt.Errorf("duplicate cron with name %v", name)
		}
	}
	duration := time.Second * time.Duration(timeToStart)
	cron := Cron{
		Name:        name,
		TimeToStart: timeToStart,
		nextStart:   time.Now().Add(duration),
		count:       0,
	}
	taskStorage[&cron] = task
	return true, nil
}

// InitTaskScheduler
// Init scheduler. Can use only once in runtime.
func InitTaskScheduler() (bool, error) {
	fmt.Println("Tried to start scheduler")
	if initStage {
		return false, fmt.Errorf("scheduler has benn started already")
	}
	go func() {
		currentStageTime := time.Now()
		for true {
			for cron, task := range taskStorage {
				if cron.nextStart.Before(currentStageTime) {
					go task.Do()
					duration := time.Second * time.Duration(cron.TimeToStart)
					cron.lastStart = cron.nextStart
					cron.nextStart = time.Now().Add(duration)
					if cron.count == 18446744073709551615 {
						cron.count = 0
						// reset test counter
					} else {
						cron.count = cron.count + 1
					}
				}
			}
			currentStageTime = time.Now()
		}
	}()
	fmt.Println("Success start scheduler")
	initStage = true
	return true, nil
}

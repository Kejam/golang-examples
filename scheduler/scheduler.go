package main

import (
	"fmt"
	"time"
)

type Cron struct {
	Name            string
	TimeToStart     int // in second
	TimeToNextStart time.Time
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
		Name:            "Up",
		TimeToStart:     10,
		TimeToNextStart: time.Now(),
	}
	tasks[&up] = DoUp{}
	down := Cron{
		Name:            "Down",
		TimeToStart:     20,
		TimeToNextStart: time.Now(),
	}
	tasks[&down] = DoDown{}
	for true {
		fmt.Println("Start range at ", currentStageTime)
		for cron, task := range tasks {
			if cron.TimeToNextStart.Before(currentStageTime) {
				go task.Do()
				duration := time.Second * time.Duration(cron.TimeToStart)
				cron.TimeToNextStart = time.Now().Add(duration)
				fmt.Println("Set new start time to cron: ", cron)
			}
		}
		time.Sleep(time.Second * time.Duration(1))
		currentStageTime = time.Now()
	}
}

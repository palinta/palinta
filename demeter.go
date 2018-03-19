package main

import (
	"strconv"
	"time"

	"github.com/matisszilard/palinta/palinta/model"
)

func saveTemperature(temperature string) {
	var temp model.Temperature
	lt := time.Now()
	temp.Time = lt.String()

	t, err := strconv.ParseFloat(temperature, 64)
	if err != nil {
		t = 0.0
	}
	temp.Temperature = t

	dbStore.Temperatures().Save(temp)
}

package main

import (
	"math/rand"
	"time"
	"fmt"
)

func SolarRadiation() string {
	
	quantification := rand.Intn(1281)

	frequency_rate := 60.0 * 1e9

	time.Sleep(time.Duration(frequency_rate))

	text := fmt.Sprintf("Medição Radiação Solar: %d W/m2", quantification)

	return text
}

func Sensor(sensor string) string {
	if sensor == "SolarRadiation" {
		return SolarRadiation()
	} else {
		return "Sensor not found"
	}
}
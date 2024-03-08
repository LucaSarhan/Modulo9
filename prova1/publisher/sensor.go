package main

import (
	"math/rand"
	"time"
	"fmt"
)

func Fridge() string {
	const id = "lj01f01"
	
	temp := rand.Intn(31) - 10

	tempalert := ""
	if temp > 10 {
		tempalert = "[ALERTA: Temperatura alta]"
	} else if temp < 2 { 
		tempalert = "[ALERTA: Temperatura baixa]"
	}

	answer := fmt.Sprintf(`
	{
		"id": "%s",
		"tipo": "geladeira",
		"temperatura": "%d",
		"timestamp": "%s"
		"alerta": "%s",
	  }
	`, id, temp, time.Now().Format(time.RFC3339), tempalert)

	return answer
}

func Freezer() string {
	const id = "lj01f01"
	
	temp := rand.Intn(51) - 50

	tempalert := ""
	if temp > -15 {
		tempalert = "[ALERTA: Temperatura alta]"
	} else if temp < -25 { 
		tempalert = "[ALERTA: Temperatura baixa]"
	}

	answer := fmt.Sprintf(`
	{
		"id": "%s",
		"tipo": "freezer",
		"temperatura": "%d",
		"timestamp": "%s"
		"alerta": "%s",
	  }
	`, id, temp, time.Now().Format(time.RFC3339), tempalert)

	return answer
}

func Choice(location string) string {
	switch location {
	case "freezer":
		return Freezer()
	case "fridge":
		return Fridge()
	default:
		return "Setor inválido!"
	}
}

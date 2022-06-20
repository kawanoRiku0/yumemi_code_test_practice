package main

import (
	"esports/handler"
	"esports/inits"
	"esports/registry"
	"fmt"
)

func main() {

	w, err := inits.InitCSVFileWriter("ranking.csv")
	if err != nil {
		fmt.Printf("init csv error : %s\n", err.Error())
	}

	reg := registry.NewRegistry(w)
	controller := handler.NewController(reg)

	controller.CalculateRanking()
}

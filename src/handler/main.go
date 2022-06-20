package handler

import (
	"esports/usecase"
	"fmt"
	"os"
)

func (c *Controller) CalculateRanking() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Print("need argument")
		os.Exit(1)
	}
	fName := os.Args[1]

	repo := c.reg.NewPlayer()
	_, err := usecase.GetPlayersInOrderScore(repo, fName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

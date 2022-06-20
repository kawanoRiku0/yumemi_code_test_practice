package usecase

import (
	"esports/domain/entity"
	"esports/domain/repository"
	"esports/domain/service"
	"fmt"
)

func GetPlayersInOrderScore(repo repository.Player, fileName string) (entity.Players, error) {
	logs, err := repo.GetPlayLogs(fileName)
	if err != nil {
		return nil, fmt.Errorf("[usecase error]error occured in getting play logs : %s\n", err.Error())
	}

	tallyer := service.NewTallyer(logs)
	playerIDToAverage := tallyer.CalcAverageScorePerPlayer()
	players := entity.NewPlayers(playerIDToAverage)
	players.SortByScore()

	rankedPlayers := tallyer.Ranking(players, 10)
	err = repo.SaveRanking(rankedPlayers)
	if err != nil {
		return nil, fmt.Errorf("[usecase error] error occured in saving ranking : %s \n", err.Error())
	}

	return rankedPlayers, nil
}

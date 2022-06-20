package service

import (
	"esports/domain/entity"
	"math"
)

// ログの集計オブジェクト(ちょっと広義過ぎるかも？)
type Tallyer struct {
	PlayLogs entity.PlayLogs
}

func NewTallyer(playlogs entity.PlayLogs) *Tallyer {
	tallyer := new(Tallyer)
	tallyer.PlayLogs = playlogs

	return tallyer
}

func (s *Tallyer) Ranking(players entity.Players, minNumOfBeRankedIn uint) entity.Players {
	firstPlayer := players[0]
	firstPlayer.Rank = 1
	rankedPlayers := entity.Players{firstPlayer}

	// ランクインした人数
	numOfBeRankedIn := 1
	for _, p := range players[1:] {
		prevPlayer := rankedPlayers[len(rankedPlayers)-1]

		// 直前のプレイヤーと同点なら同率
		if prevPlayer.AvarageScore == p.AvarageScore {
			p.Rank = prevPlayer.Rank
			rankedPlayers = append(rankedPlayers, p)
			numOfBeRankedIn++

			// ランクイン可能人数を上回ったら離脱
		} else if numOfBeRankedIn >= int(minNumOfBeRankedIn) {
			break
		} else {
			p.Rank = prevPlayer.Rank + 1
			rankedPlayers = append(rankedPlayers, p)
			numOfBeRankedIn++
		}
	}

	return rankedPlayers

}

func (s *Tallyer) CalcAverageScorePerPlayer() map[string]int {
	playerIDToTotalScore := map[string]int{}
	playerIDToCount := map[string]int{}

	for _, playLog := range s.PlayLogs {
		playerIDToTotalScore[playLog.PlayerID] += playLog.Score
		playerIDToCount[playLog.PlayerID]++
	}

	playerIDToAvarage := map[string]int{}
	for ID, score := range playerIDToTotalScore {
		count := playerIDToCount[ID]

		avarage := float64(score) / float64(count)
		roundedAvarage := math.Round(avarage)
		playerIDToAvarage[ID] = int(roundedAvarage)
	}

	return playerIDToAvarage
}

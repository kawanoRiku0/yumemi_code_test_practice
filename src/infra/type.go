package infra

import (
	"esports/domain/entity"
	"fmt"
	"strconv"
)

type PlayLogCSVRow struct {
	CreatedTime string `csv:"created_time"`
	PlayerID    string `csv:"player_id"`
	Score       string `csv:"score"`
}

func (c PlayLogCSVRow) Validate() (*entity.PlayLog, error) {

	if c.PlayerID == "" {
		return nil, fmt.Errorf("playerID must not be empty")
	}

	score, err := strconv.Atoi(c.Score)
	if err != nil {
		return nil, fmt.Errorf("score can't be converted to number")
	}

	if score < 0 {
		return nil, fmt.Errorf("score must be more than zero")
	}

	playLog := entity.NewPlayLog(c.PlayerID, score)

	return playLog, nil
}

type RankingCSVRow []string

func RankingCSVRowMapper(player *entity.Player) RankingCSVRow {
	rank := strconv.Itoa(player.Rank)
	score := strconv.Itoa(player.AvarageScore)

	row := RankingCSVRow{rank, player.ID, score}
	return row
}

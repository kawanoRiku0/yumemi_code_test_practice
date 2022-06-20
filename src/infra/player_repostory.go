package infra

import (
	"encoding/csv"
	"esports/domain/entity"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type PlayerRepository struct {
	csvWriter *csv.Writer
}

func NewPlayer(w *csv.Writer) *PlayerRepository {
	return &PlayerRepository{w}
}

func (repo *PlayerRepository) GetPlayLogs(fileName string) (entity.PlayLogs, error) {

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("[infra error] file open faild : %s \n", err.Error())
	}
	defer file.Close()

	var rows []*PlayLogCSVRow

	if err := gocsv.UnmarshalFile(file, &rows); err != nil {
		return nil, fmt.Errorf("[infra error] unmarshal csv file faild : %s\n", err.Error())
	}

	playLogs := entity.PlayLogs{}
	for _, row := range rows {
		playLog, err := row.Validate()
		if err != nil {
			continue
		}

		playLogs = append(playLogs, playLog)
	}
	return playLogs, nil
}

func (repo *PlayerRepository) SaveRanking(players entity.Players) error {

	RankingLabel := []string{"rank", "player_id", "mean_score"}
	RankingCSVRows := [][]string{RankingLabel}

	for _, player := range players {
		RankingCSVRow := RankingCSVRowMapper(player)
		RankingCSVRows = append(RankingCSVRows, RankingCSVRow)
	}

	err := repo.csvWriter.WriteAll(RankingCSVRows)
	if err != nil {
		return fmt.Errorf("[infra error] error occured in writing csv : %s\n", err)
	}

	return nil
}

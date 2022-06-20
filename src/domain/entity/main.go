package entity

import (
	"fmt"
	"sort"
)

type PlayLog struct {
	PlayerID string
	Score    int
}

func NewPlayLog(ID string, score int) *PlayLog {
	playLog := new(PlayLog)

	playLog.PlayerID = ID
	playLog.Score = score

	return playLog

}

type PlayLogs []*PlayLog

type Player struct {
	ID           string
	AvarageScore int
	Rank         int
}

// Playerのconstructor
func NewPlayer(id string, avarage int) (*Player, error) {
	player := new(Player)

	if avarage <= 0 {
		return nil, fmt.Errorf("avarage must be more than one")
	}

	player.ID = id
	player.AvarageScore = avarage

	return player, nil
}

type Players []*Player

// Playersのconstructor
func NewPlayers(playerIDToAvarage map[string]int) Players {
	players := Players{}
	for ID, avarage := range playerIDToAvarage {
		p, _ := NewPlayer(ID, avarage)

		players = append(players, p)
	}

	return players
}

// scoreによって降順にソート
func (p Players) SortByScore() {
	sort.Slice(p, func(i, j int) bool {
		return p[i].AvarageScore > p[j].AvarageScore
	})
}

package repository

import (
	"esports/domain/entity"
)

type Player interface {
	GetPlayLogs(fileName string) (entity.PlayLogs, error)
	SaveRanking(players entity.Players) error
}

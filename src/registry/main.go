package registry

import (
	"encoding/csv"
	"esports/domain/repository"
	"esports/infra"
)

type Registry interface {
	NewPlayer() repository.Player
}

func NewRegistry(w *csv.Writer) Registry {
	return &RegistryImpl{w}
}

type RegistryImpl struct {
	csvWriter *csv.Writer
}

func (reg *RegistryImpl) NewPlayer() repository.Player {
	return infra.NewPlayer(reg.csvWriter)
}

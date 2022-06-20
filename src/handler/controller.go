package handler

import (
	"esports/registry"
)

type Controller struct {
	reg registry.Registry
}

func NewController(reg registry.Registry) *Controller {
	return &Controller{reg: reg}
}

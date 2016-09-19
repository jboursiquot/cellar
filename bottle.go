package main

import (
	"fmt"

	"github.com/goadesign/goa"
	"github.com/jboursiquot/cellar/app"
)

// BottleController implements the bottle resource.
type BottleController struct {
	*goa.Controller
}

// NewBottleController creates a bottle controller.
func NewBottleController(service *goa.Service) *BottleController {
	return &BottleController{Controller: service.NewController("BottleController")}
}

// Show runs the show action.
func (c *BottleController) Show(ctx *app.ShowBottleContext) error {
	// BottleController_Show: start_implement

	if ctx.BottleID == 0 {
		return ctx.NotFound()
	}

	res := &app.GoaExampleBottle{
		ID:   ctx.BottleID,
		Name: fmt.Sprintf("Bottle #%d", ctx.BottleID),
		Href: app.BottleHref(ctx.BottleID),
	}

	// BottleController_Show: end_implement
	return ctx.OK(res)
}

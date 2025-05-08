package overlay

import (
	"mask_of_the_tomb/internal/game/core/events"

	"github.com/hajimehoshi/ebiten/v2"
)

type Overlay struct {
	state         deathEffectState
	image         *ebiten.Image
	alpha         float64
	OnFinishEnter *events.Event
	OnFinishExit  *events.Event
}

// type Overlay interface {
// 	StartEnter()
// 	StartExit()
// 	Update()
// 	Draw() // can be to draw a shader, a
// }

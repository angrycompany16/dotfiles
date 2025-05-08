package overlay

import (
	"mask_of_the_tomb/internal/game/core/events"

	"github.com/hajimehoshi/ebiten/v2"
)

type OverlayImage struct {
	state         deathEffectState
	image         *ebiten.Image
	OnFinishEnter *events.Event
	OnFinishExit  *events.Event
}

func (oi *OverlayImage) StartEnter() {
	oi.state = enter
}

func (oi *OverlayImage) StartExit() {
	oi.state = exit
}

type Overlay interface {
	FadeIn(t float64)
	FadeOut(t float64)
	Update()
	Draw() // can be to draw a shader, a
}

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

type Overlay interface {
	FadeIn()
	FadeOut()
	StartExit()
	Update()
	Draw() // can be to draw a shader, a
}

package overlay

import (
	"mask_of_the_tomb/internal/game/core/events"
	"mask_of_the_tomb/internal/game/core/rendering"

	"github.com/hajimehoshi/ebiten/v2"
)

type overlayState int

const (
	enter overlayState = iota
	exit
	idle
)

type Overlay struct {
	OverlayContent
	state         overlayState
	image         *ebiten.Image
	OnFinishEnter *events.Event
	OnFinishExit  *events.Event
}

func (oi *Overlay) StartEnter() {
	oi.state = enter
}

func (oi *Overlay) StartExit() {
	oi.state = exit
}

func NewOverlayImage() *Overlay {
	return &Overlay{
		state:         idle,
		image:         ebiten.NewImage(rendering.GameWidth, rendering.GameHeight),
		OnFinishEnter: events.NewEvent(),
		OnFinishExit:  events.NewEvent(),
	}
}

type OverlayContent interface {
	// FadeIn()
	// FadeOut()
	// Update()
	Draw(t float64) // can be to draw a shader, a
}

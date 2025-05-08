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

type OverlayImage struct {
	state         overlayState
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

func NewOverlayImage() *OverlayImage {
	return &OverlayImage{
		state: idle,
		image: ebiten.NewImage(rendering.GameWidth, rendering.GameHeight)
	}
}

// func (oi *OverlayImage) Update() {

// }

type Overlay interface {
	FadeIn(t float64)
	FadeOut(t float64)
	Update()
	Draw() // can be to draw a shader, a
}

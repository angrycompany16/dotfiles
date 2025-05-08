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

type Overlayer struct {
	OverlayContent
	state         overlayState
	image         *ebiten.Image
	OnFinishEnter *events.Event
	OnFinishExit  *events.Event
}

func (oi *Overlayer) StartEnter() {
	oi.state = enter
}

func (oi *Overlayer) StartExit() {
	oi.state = exit
}

func NewOverlayImage() *Overlayer {
	return &Overlayer{
		state:         idle,
		image:         ebiten.NewImage(rendering.GameWidth, rendering.GameHeight),
		OnFinishEnter: events.NewEvent(),
		OnFinishExit:  events.NewEvent(),
	}
}

type OverlayContent interface {
	FadeIn()
	FadeOut()
	Update()
	Draw() // can be to draw a shader, a
}

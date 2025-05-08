package overlay

import (
	"mask_of_the_tomb/internal/game/core/events"
	"mask_of_the_tomb/internal/game/core/rendering"
	"mask_of_the_tomb/internal/maths"

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
	t             float64
	state         overlayState
	image         *ebiten.Image
	OnFinishEnter *events.Event
	OnFinishExit  *events.Event
}

func (oi *Overlay) FadeIn() {
	oi.state = enter
	oi.t = maths.Lerp(oi.t, 3, 0.01)
	if 1-oi.t <= 0.01 {
		oi.t = 1
		oi.OnFinishEnter.Raise(events.EventInfo{})
		oi.state = exit
	}
}

func (oi *Overlay) FadeOut() {
	oi.state = exit
	d.t = maths.Lerp(d.t, -2, 0.01)
	if d.alpha <= 0.01 {
		d.alpha = 0
		d.OnFinishExit.Raise(events.EventInfo{})
		d.state = idle
	}
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

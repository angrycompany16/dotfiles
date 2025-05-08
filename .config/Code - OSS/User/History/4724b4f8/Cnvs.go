package overlay

import (
	"image/color"
	ebitenrenderutil "mask_of_the_tomb/internal/ebitenrenderutil"
	"mask_of_the_tomb/internal/game/core/events"
	"mask_of_the_tomb/internal/game/core/rendering"
	"mask_of_the_tomb/internal/maths"
)

// TODO: Generalize into an overlay type?
// For suuure we can make an interface for this, or something

var (
	OverlayColor = []uint8{20, 16, 19}
)

type DeathEffect struct {
	*Overlayer
	alpha float64
}

func (d *DeathEffect) FadeIn() {
	d.alpha = maths.Lerp(d.alpha, 3, 0.01)
	if 1-d.alpha <= 0.01 {
		d.alpha = 1
		d.OnFinishEnter.Raise(events.EventInfo{})
		d.state = exit
	}
}

func (d *DeathEffect) FadeOut() {
	d.alpha = maths.Lerp(d.alpha, -2, 0.01)
	if d.alpha <= 0.01 {
		d.alpha = 0
		d.OnFinishExit.Raise(events.EventInfo{})
		d.state = idle
	}
}

func (d *DeathEffect) Update() {
	switch d.state {
	case enter:
		d.FadeIn()
	case exit:
		d.FadeOut()
	case idle:
	}
}

func (d *DeathEffect) Draw() {
	alpha := uint8(d.alpha * 255)
	d.image.Fill(color.RGBA{OverlayColor[0], OverlayColor[1], OverlayColor[2], alpha})
	ebitenrenderutil.DrawAt(d.image, rendering.RenderLayers.Overlay, 0, 0)
}

func NewDeathEffect() *DeathEffect {
	return &DeathEffect{
		Overlayer: NewOverlayImage(),
		alpha:     0.0,
	}
}

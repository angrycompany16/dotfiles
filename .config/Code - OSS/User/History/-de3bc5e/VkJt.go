package particles

import (
	"image/color"
	"mask_of_the_tomb/internal/ebitenrenderutil"
	"mask_of_the_tomb/internal/maths"

	"github.com/hajimehoshi/ebiten/v2"
)

type Particle struct {
	posX, posY                  float64
	velX, velY                  float64
	angle, angleVel             float64
	startScale, scale, endScale float64
	lifetime, t                 float64 // seconds
	startColor, color, endColor color.Color
	sprite                      *ebiten.Image
	colorOverlay                *ebiten.Image
	particleImage               *ebiten.Image
}

func (p *Particle) update() bool {
	dt := 0.016666667
	p.t += dt
	if p.t > p.lifetime {
		return true
	}

	p.scale = maths.Lerp(p.startScale, p.endScale, p.t/p.lifetime)
	if p.scale <= 0.0001 {
		return true
	}

	p.angle += p.angleVel * dt
	p.posX += p.velX * dt
	p.posY += p.velY * dt
	p.color = maths.Mix(p.startColor, p.endColor, p.t/p.lifetime)

	return false
}

// I've got it now
func (p *Particle) draw(layer *ebiten.Image, offsetX, offsetY float64) {
	ebitenrenderutil.DrawAtRotatedScaled(p.sprite, layer, p.posX-offsetX, p.posY-offsetY, p.angle, p.scale, p.scale)

	p.colorOverlay.Fill(p.color)
	overlayOp := ebitenrenderutil.RotatedScaledOp(p.colorOverlay, p.posX-offsetX, p.posY-offsetY, p.angle, p.scale, p.scale)
	overlayOp.Blend = ebiten.BlendSourceOver
	layer.DrawImage(p.colorOverlay, overlayOp)
}

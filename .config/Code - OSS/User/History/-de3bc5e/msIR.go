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

func (p *Particle) draw(layer *ebiten.Image, offsetX, offsetY float64) {
	p.particleImage.Clear()
	// ebitenrenderutil.DrawAt(p.sprite, p.particleImage, 0, 0)

	image := ebiten.NewImage()

	p.colorOverlay.Fill(color.Black)
	op := ebiten.DrawImageOptions{}
	op.Blend = ebiten.BlendSourceAtop
	p.particleImage.DrawImage(p.colorOverlay, &op)
	// ebitenrenderutil.DrawAt(p.colorOverlay, p.particleImage, 0, 0)
	// layer.DrawImage(p.colorOverlay, overlayOp)

	ebitenrenderutil.DrawAtRotatedScaled(p.colorOverlay, layer, p.posX-offsetX, p.posY-offsetY, p.angle, p.scale, p.scale)
	// ebitenrenderutil.DrawAtRotatedScaled(p.colorOverlay, layer, p.posX-offsetX, p.posY-offsetY, p.angle, p.scale, p.scale)
}

// // I've got it now (?)))
// func (p *Particle) draw(layer *ebiten.Image, offsetX, offsetY float64) {
// 	p.particleImage.Clear()
// 	ebitenrenderutil.DrawAt(p.sprite, p.particleImage, 0, 0)

// 	p.colorOverlay.Fill(p.color)
// 	op := ebiten.DrawImageOptions{}
// 	op.Blend = ebiten.BlendSourceAtop
// 	p.particleImage.DrawImage(p.colorOverlay, &op)
// 	// ebitenrenderutil.DrawAt(p.colorOverlay, p.particleImage, 0, 0)
// 	// layer.DrawImage(p.colorOverlay, overlayOp)

// 	ebitenrenderutil.DrawAtRotatedScaled(p.particleImage, layer, p.posX-offsetX, p.posY-offsetY, p.angle, p.scale, p.scale)
// 	// ebitenrenderutil.DrawAtRotatedScaled(p.colorOverlay, layer, p.posX-offsetX, p.posY-offsetY, p.angle, p.scale, p.scale)
// }

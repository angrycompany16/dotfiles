package camera

import (
	"mask_of_the_tomb/internal/game/core/rendering"
	"mask_of_the_tomb/internal/maths"
)

// TODO: Solve the single-screen height bug better

var (
	_camera = &Camera{}
)

type Camera struct {
	posX, posY                 float64
	width, height              float64
	offsetX, offsetY           float64
	shakeOffsetX, shakeOffsetY float64
	shaking                    bool
	shakeDuration              float64
	shakeStrength              float64
	shakeTime                  float64
}

func Init(width, height, offsetX, offsetY float64) {
	SetBorders(width, height)
	_camera.offsetX, _camera.offsetY = offsetX, offsetY
}

func Update() {
	if !_camera.shaking {
		return
	}

	_camera.shakeTime += 1.0 / 60.0
	if _camera.shakeTime > _camera.shakeDuration {
		_camera.shaking = false
		_camera.shakeTime = 0
		return
	}

	_camera.shakeOffsetX = 
}

func GetPos() (float64, float64) {
	if _camera.height == 272 {
		return _camera.posX, _camera.posY + 1
	}
	return _camera.posX, _camera.posY
}

func SetPos(x, y float64) {
	if _camera.height == 272 {
		return
	}
	_camera.posX = maths.Clamp(x-_camera.offsetX, 0, _camera.width-rendering.GameWidth)
	_camera.posY = maths.Clamp(y-_camera.offsetY, 0, _camera.height-rendering.GameHeight)
}

func SetBorders(width, height float64) {
	_camera.width = width
	_camera.height = height
}

func Shake(duration, strength float64) {
	// Shake the camera
}

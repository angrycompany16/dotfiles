package audiocontext

import "github.com/hajimehoshi/ebiten/v2/audio"

var (
	_globalAudioContext *GlobalAudioContext
)

type GlobalAudioContext struct {
	audio.Context
}

func Get() *GlobalAudioContext {
	if _globalAudioContext 
}

package levelmemory

import "mask_of_the_tomb/internal/game/world/entities"

type SlamboxPosition struct {
	X, Y float64
}

type LevelMemory struct {
	Slamboxes []*entities.Slambox
}

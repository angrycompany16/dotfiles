package world

import (
	"fmt"
	"mask_of_the_tomb/internal/errs"
	"mask_of_the_tomb/internal/game/core/assetloader/assettypes"
	"mask_of_the_tomb/internal/game/world/levelmemory"

	ebitenLDTK "github.com/angrycompany16/ebiten-LDTK"
)

type World struct {
	worldLDTK        *ebitenLDTK.World
	ActiveLevel      *Level
	worldStateMemory map[string]levelmemory.LevelMemory
}

func NewWorld() *World {
	return &World{
		worldStateMemory: make(map[string]levelmemory.LevelMemory),
	}
}

func (w *World) Load() {
	w.worldLDTK = assettypes.NewLDTKAsset(LDTKMapPath)
}

func (w *World) Init(initLevelName string, worldStateMemory map[string]levelmemory.LevelMemory) {
	if initLevelName == "" {
		initLevelName = w.worldLDTK.Levels[0].Name
	}
	ChangeActiveLevel(w, initLevelName, "")
	for id, levelmemory := range worldStateMemory {
		w.worldStateMemory[id] = levelmemory
	}
}

func (w *World) LoadMemory(memory map[string]levelmemory.LevelMemory) {
}

func (w *World) Update() {
	w.ActiveLevel.Update()
}

func (w *World) GetWorldStateMemory() map[string]levelmemory.LevelMemory {
	return w.worldStateMemory
}

func ChangeActiveLevel[T string | int](world *World, id T, doorIid string) error {
	var newLevelLDTK ebitenLDTK.Level
	var err error

	switch levelId := any(id).(type) {
	case string:
		newLevelLDTK, err = world.worldLDTK.GetLevelByName(levelId)
		if err != nil {
			fmt.Println("Couldn't switch levels by name (id string), trying Iid...")
			var Ierr error
			newLevelLDTK, Ierr = world.worldLDTK.GetLevelByIid(levelId)
			if Ierr != nil {
				fmt.Println("Error when switching levels by Iid (id string)")
				return Ierr
			}
		}
	case int:
		newLevelLDTK, err = world.worldLDTK.GetLevelByUid(levelId)
		if err != nil {
			fmt.Println("Error when switching levels (id int)")
			return err
		}
	}

	newLevel, err := newLevel(&newLevelLDTK, &world.worldLDTK.Defs)
	if err != nil {
		return err
	}

	if memory, ok := world.worldStateMemory[newLevelLDTK.Iid]; ok {
		newLevel.restoreFromMemory(&memory)
	}

	if world.ActiveLevel != nil {
		world.SaveLevel(world.ActiveLevel)
	}
	world.SaveLevel(newLevel)

	newLevel.resetX, newLevel.resetY = newLevel.GetDefaultSpawnPoint()
	if doorIid != "" {
		doorEntity := errs.Must(newLevel.levelLDTK.GetEntityByIid(doorIid))
		newLevel.resetX, newLevel.resetY = doorEntity.Px[0], doorEntity.Px[1]
	}
	world.ActiveLevel = newLevel
	return nil
}

func (w *World) SaveLevel(level *Level) {
	w.worldStateMemory[level.levelLDTK.Iid] = levelmemory.LevelMemory{level.GetSlamboxPositions()}
}

func (w *World) ResetActiveLevel() (float64, float64) {
	_resetX, _resetY := w.ActiveLevel.GetResetPoint()

	// reset slambox positions
	level := w.worldStateMemory[w.ActiveLevel.levelLDTK.Iid]
	w.ActiveLevel.restoreFromMemory(&level)
	// reset player position
	w.ActiveLevel.resetX = _resetX
	w.ActiveLevel.resetY = _resetY
	return _resetX, _resetY
}

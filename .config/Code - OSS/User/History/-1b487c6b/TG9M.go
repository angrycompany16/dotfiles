package main

import save "mask_of_the_tomb/internal/game/core/savesystem"

var saveProfile int

func main() {

	save.SaveGame(save.GameData{}, 1)
	save.SaveGame(save.GameData{}, 99)
}

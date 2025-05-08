package save

import (
	"encoding/json"
	"fmt"
	"mask_of_the_tomb/internal/errs"
	"mask_of_the_tomb/internal/files"
	"mask_of_the_tomb/internal/game/world/levelmemory"
	"os"
	"path/filepath"
)

var (
	savePath = filepath.Join("save", "savedata.json")
)

type GameData struct {
	WorldStateMemory map[string]levelmemory.LevelMemory
}

func SaveGame(data GameData) {
	// fmt.Println("Saving game.....")
	// defer fmt.Println("Done!")
	// TODO: Check if save file directory exists
	exists := errs.Must(files.Exists(savePath))
	if !exists {
		os.MkdirAll(filepath.Dir(savePath), os.ModePerm)
	}
	file := errs.Must(os.Create(savePath))
	// if err != nil {
	// 	fmt.Println("Could not create file ", s.savePath)
	// 	fmt.Println(err)
	// 	return
	// }
	defer file.Close()
	errs.MustSingle(json.NewEncoder(file).Encode(&data))
	// if err != nil {
	// 	fmt.Println("Could not write save data to ", s.savePath)
	// 	fmt.Println(err)
	// 	return
	// }
}

func LoadGame() GameData {
	fmt.Println("Loading game.....")
	defer fmt.Println("Done!")
	exists := errs.Must(files.Exists(savePath))
	if !exists {
		SaveGame(GameData{})
		return GameData{}
	}

	gameData := GameData{}
	file := errs.Must(os.Open(savePath))
	// if err != nil {
	// 	fmt.Println("Could not open file")
	// 	fmt.Println(err)
	// 	return
	// }
	defer file.Close()

	errs.MustSingle(json.NewDecoder(file).Decode(&gameData))
	// if err != nil {
	// 	fmt.Println("Could not decode JSON")
	// 	fmt.Println(err)
	// 	return
	// }
	return gameData
}

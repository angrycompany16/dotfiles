package save

import (
	"encoding/json"
	"fmt"
	"mask_of_the_tomb/internal/errs"
	"mask_of_the_tomb/internal/files"
	"os"
	"path/filepath"
)

var (
	// GlobalSave = Save{
	// 	GameData: NewGameData(),
	// 	savePath: savePath,
	// }
	savePath = filepath.Join("save", "savedata.json")
)

type gameData struct {
}

func NewGameData() gameData {
	return gameData{}
}

// type Save struct {
// 	GameData gameData
// 	savePath string
// }

func SaveGame() {
	fmt.Println("Saving game.....")
	defer fmt.Println("Done!")
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
	errs.MustSingle(json.NewEncoder(file).Encode(&s.GameData))
	// if err != nil {
	// 	fmt.Println("Could not write save data to ", s.savePath)
	// 	fmt.Println(err)
	// 	return
	// }
}

func LoadGame() {
	fmt.Println("Loading game.....")
	defer fmt.Println("Done!")
	gameData := NewGameData()
	file, err := os.Open(s.savePath)
	if err != nil {
		fmt.Println("Could not open file")
		fmt.Println(err)
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&gameData)
	if err != nil {
		fmt.Println("Could not decode JSON")
		fmt.Println(err)
		return
	}
	s.GameData = gameData
}

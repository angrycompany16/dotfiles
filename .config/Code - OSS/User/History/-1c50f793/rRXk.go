package resources

var (
	ResourceMan ResourceManager
)

type ResourceManager struct {
	PlayerData PlayerData
}

type PlayerData struct {
	// Contains all fields required for interactions with player entity
}

type LevelData struct {
	// Contains all fields required for interactions with level entity
}

type MenuData struct {
	// Contains all fields required for interactions with menu entity
}

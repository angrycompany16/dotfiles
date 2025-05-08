package resources

var (
	ResourceMan ResourceManager
)

type ResourceManager struct {
	PlayerData PlayerData
}

func (r *ResourceManager) Read()

type PlayerData struct {
	// Contains all fields required for interactions with player entity
}

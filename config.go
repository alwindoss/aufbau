package aufbau

// Configuration represents the configuration that the client wants to store with Aufbau
type Configuration struct {
	OrgID    string
	EntityID string
	ConfigID string
	Config   string
}

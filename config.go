package aufbau

// Configuration represents the configuration that the client wants to store with Aufbau
type Configuration struct {
	OrgID    string `json:"org_id,omitempty"`
	EntityID string `json:"entity_id,omitempty"`
	ConfigID string `json:"config_id,omitempty"`
	Config   string `json:"config,omitempty"`
}

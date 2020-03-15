package aufbau

// Repository is the interface that has to be implemented by storage
type Repository interface {
	Create(*Configuration) (*Configuration, error)
	Fetch(orgID, entityID, configID string) (*Configuration, error)
	FetchAll(orgID, entityID string) ([]*Configuration, error)
	Update(*Configuration) (*Configuration, error)
	Delete(*Configuration) (*Configuration, error)
}

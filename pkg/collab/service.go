package collab

// Service implements Servicer interface.
type Service struct {
	Repo Repository
}

// NewuserService is a factory func to create a new user.Service.
func NewcollabService(repo Repository) *Service {
	var ps Service
	ps.Repo = repo

	return &ps
}

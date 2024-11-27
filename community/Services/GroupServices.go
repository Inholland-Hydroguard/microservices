package Services

import (
	"microservices/community/Domain"
)

type GroupService interface {
	GetAllGroup() ([]Domain.Group, error)
	FindGroupById(id string) (*Domain.Group, *Domain.AppError)
}

// "constructor" like function
// whereby we pass in the repo (interface) as a dependency
type DefaultGroupService struct {
	repo Domain.GroupRepository
}

// receiver function -attaches it as a method to a class
func (s DefaultGroupService) GetAllGroup() ([]Domain.Group, error) {

	//Once again we talk to the interface
	return s.repo.FindAll()
}

func (s DefaultGroupService) FindGroupById(id string) (*Domain.Group, *Domain.AppError) {

	//Once again we talk to the interface
	return s.repo.FindById(id)
}

// Helper function to instantiate Group service
func NewGroupService(repo Domain.GroupRepository) DefaultGroupService {
	return DefaultGroupService{repo}
}

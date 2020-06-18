package user

import "github.com/thearyanahmed/wallet/schema"

type Service struct {
	userRepository
}

func (service *Service) FindUserById(id uint) ( *schema.User, []error )  {
	return service.userRepository.findUserById(id)
}

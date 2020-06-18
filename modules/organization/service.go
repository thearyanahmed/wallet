package organization

import "github.com/thearyanahmed/wallet/schema"

type Service struct {
	organizationRepository
}

func (s *Service) CreateOrganization(userId uint,name string) (*schema.Organization,error) {
	return s.organizationRepository.createOrganization(userId,name)
}
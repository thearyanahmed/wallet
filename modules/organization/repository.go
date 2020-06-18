package organization

import (
	"github.com/thearyanahmed/wallet/database"
	"github.com/thearyanahmed/wallet/schema"
)

type organizationRepository struct {
	schema.Organization
}

func (repo *organizationRepository) createOrganization(userId uint,name string) (*schema.Organization,error) {
	org := schema.Organization{
		Name:      name,
		UserID:    userId,
	}
	err := database.DB().Create(&org).Error

	if err != nil {
		return nil, err
	}
	return &org, nil
}

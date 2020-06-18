package organization

import (
	"github.com/thearyanahmed/wallet/database"
	"github.com/thearyanahmed/wallet/schema"
)

type organizationRepository struct {
	schema.Organization
}

func createOrganization(userId uint,name string) (*schema.Organization,[]error) {
	org := schema.Organization{
		Name:      name,
		UserID:    userId,
	}
	errs := database.DB().Create(&org).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return &org, nil
}

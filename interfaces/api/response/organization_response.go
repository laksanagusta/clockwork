package response

import (
	"clockwork-server/domain/model"
	"time"
)

type Organization struct {
	ID        uint      `json:"id"`
	Name      string    `json:"title"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FormatOrganization(organization model.Organization) Organization {
	var dataOrganization Organization

	dataOrganization.ID = organization.ID
	dataOrganization.Name = organization.Name
	dataOrganization.Address = organization.Address
	dataOrganization.CreatedAt = organization.CreatedAt
	dataOrganization.UpdatedAt = organization.UpdatedAt

	return dataOrganization
}

func FormatOrganizations(order []model.Organization) []Organization {
	var dataOrganizations []Organization

	for _, value := range order {
		singleDataOrganization := FormatOrganization(value)
		dataOrganizations = append(dataOrganizations, singleDataOrganization)
	}

	return dataOrganizations
}

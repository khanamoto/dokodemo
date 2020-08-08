package model

type Department struct {
	ID             uint64 `db:"id"`
	OrganizationID uint64 `db:"organization_id"`
	Name           string `db:"name" validate:"required,max=255"`
	URL            string `db:"url" validate:"required,alphanum,max=255"`
}

package model

type Access struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	SupplierID string  `json:"supplierID"`
	URL        *string `json:"url"`
	User       *string `json:"user"`
	CreatedAt  string  `json:"createdAt"`
	UpdatedAt  string  `json:"updatedAt"`
}

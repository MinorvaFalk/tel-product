package model

type Products struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Price       float64 `json:"price"`
	Variety     *string `json:"variety"`
	Rating      float32 `json:"rating"`
	Stock       int     `json:"stock"`
}

type ProductsCreateRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
	Price       float64 `json:"price" validate:"required"`
	Variety     *string `json:"variety"`
	Stock       int     `json:"stock" validate:"required"`
}

type ProductsUpdateRequest struct {
	ID          string  `json:"id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
	Price       float64 `json:"price" validate:"required"`
	Variety     *string `json:"variety"`
	Rating      float32 `json:"rating"`
	Stock       int     `json:"stock" validate:"required"`
}

type ProductsPatchRequest struct {
	ID          string   `json:"id" validate:"required,uuid4_rfc4122" mapstructure:"-"`
	Name        *string  `json:"name" mapstructure:"name,omitempty"`
	Description *string  `json:"description" mapstructure:"description,omitempty"`
	Price       *float64 `json:"price" mapstructure:"price,omitempty"`
	Variety     *string  `json:"variety" mapstructure:"variety,omitempty"`
	Rating      *float32 `json:"rating" mapstructure:"rating,omitempty"`
	Stock       *int     `json:"stock" mapstructure:"stock,omitempty"`
}

type ProductIDRequest struct {
	ID string `json:"id" query:"id" param:"id" form:"id" validate:"required,uuid4_rfc4122"`
}

package dto

type CategoryRequestDto struct {
	Name        string `json:"name" validate:"required,max=100"`
	Description string `json:"description,omitempty" validate:"max=500"`
	CreatedBy   string `json:"created_by,omitempty" validate:"max=50"`
	UpdatedBy   string `json:"updated_by,omitempty" validate:"max=50"`
}

type CategoryResponseDto struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	CreatedBy   string `json:"created_by,omitempty" validate:"max=50"`
	UpdatedBy   string `json:"updated_by,omitempty" validate:"max=50"`
}

type SizeRequestDto struct {
	Name       string `json:"name" validate:"required,max=50"`
	CategoryID uint   `json:"category_id" validate:"required"`
}

type SizeResponseDto struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	CategoryID uint   `json:"category_id"`
}

type ProductSizeRequestDto struct {
	ProductID uint    `json:"product_id" validate:"required"`
	SizeID    uint    `json:"size_id" validate:"required"`
	Price     float64 `json:"price" validate:"required"`
}

type ProductSizeResponseDto struct {
	ProductID uint            `json:"product_id"`
	SizeID    uint            `json:"size_id"`
	Price     float64         `json:"price"`
	Size      SizeResponseDto `json:"size"`
}

type ProductRequestDto struct {
	Name        string                  `json:"name" validate:"required,max=100"`
	Description string                  `json:"description,omitempty" validate:"max=500"`
	CategoryID  uint                    `json:"category_id" validate:"required"`
	ImageURL    string                  `json:"image_url,omitempty" validate:"max=255"`
	Prices      []ProductSizeRequestDto `json:"prices,omitempty"`
}

type ProductResponseDto struct {
	Id          uint                     `json:"id"`
	Name        string                   `json:"name"`
	Description string                   `json:"description,omitempty"`
	Category    CategoryResponseDto      `json:"category"`
	CategoryID  uint                     `json:"category_id"`
	ImageURL    string                   `json:"image_url,omitempty"`
	Prices      []ProductSizeResponseDto `json:"prices,omitempty"`
}

package dto

type CategoryRequestDto struct {
	Name        string `json:"name" binding:"required,max=100"`
	Description string `json:"description,omitempty" binding:"max=500,non_special_char"`
	CreatedBy   string `json:"created_by,omitempty" binding:"max=50"`
	UpdatedBy   string `json:"updated_by,omitempty" binding:"max=50"`
}

type CategoryResponseDto struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	CreatedBy   string `json:"created_by,omitempty" binding:"max=50"`
	UpdatedBy   string `json:"updated_by,omitempty" binding:"max=50"`
}

type SizeRequestDto struct {
	Name       string `json:"name" binding:"required,max=50,non_special_char"`
	CategoryID uint   `json:"category_id" binding:"required"`
}

type SizeResponseDto struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	CategoryID uint   `json:"category_id"`
}

type ProductSizeRequestDto struct {
	ProductID uint    `json:"product_id" binding:"required"`
	SizeID    uint    `json:"size_id" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
}

type ProductSizeResponseDto struct {
	ProductID uint            `json:"product_id"`
	SizeID    uint            `json:"size_id"`
	Price     float64         `json:"price"`
	Size      SizeResponseDto `json:"size"`
}

type ProductPriceRequestDto struct {
	SizeID uint    `json:"size_id" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
}

type ProductRequestDto struct {
	Name        string                   `json:"name" binding:"required,max=100"`
	Description string                   `json:"description,omitempty" binding:"max=500"`
	CategoryID  uint                     `json:"category_id" binding:"required"`
	ImageURL    string                   `json:"image_url,omitempty" binding:"max=255"`
	Prices      []ProductPriceRequestDto `json:"prices,omitempty"`
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

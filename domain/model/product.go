package model

type Category struct {
	BaseModel
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	Description string `gorm:"type:text" json:"description,omitempty"`
}

type Size struct {
	BaseModel
	Name       string `gorm:"type:varchar(50);not null" json:"name"`
	CategoryID uint   `gorm:"not null" json:"category_id"`
}

type Price struct {
	BaseModel
	ProductID uint    `gorm:"not null" json:"product_id"`
	SizeID    uint    `gorm:"not null" json:"size_id"`
	Price     float64 `gorm:"not null" json:"price"`
	Size      Size    `gorm:"foreignKey:SizeID" json:"size"`
}

type Product struct {
	BaseModel
	Name        string   `gorm:"type:varchar(100);not null" json:"name"`
	Description string   `gorm:"type:text" json:"description,omitempty"`
	Category    Category `gorm:"foreignKey:CategoryID" json:"category"`
	CategoryID  uint     `gorm:"not null" json:"category_id"`
	Prices      []Price  `gorm:"foreignKey:ProductID"`
}

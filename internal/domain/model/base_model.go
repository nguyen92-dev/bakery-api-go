package model

type BaseModel struct {
	ID uint `gorm:"primaryKey" json:"id"`

	CreatedBy string `gorm:"type:varchar(50);not null" json:"created_by"`
	UpdatedBy string `gorm:"type:varchar(50);not null" json:"updated_by"`
}

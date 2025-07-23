package model

type BaseModel struct {
	Id uint `gorm:"primaryKey" json:"id"`

	CreatedBy string `gorm:"type:varchar(50);not null" json:"created_by"`
	UpdatedBy string `gorm:"type:varchar(50);not null" json:"updated_by"`
}

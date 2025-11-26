package entity

type Example struct {
	ID          uint   `gorm:"primaryKey" json:"id" example:"1"`
	Description string `gorm:"type:text" json:"description" binding:"required" example:"This is an example"`
}

type Examples []Example

func (e *Example) TableName() string {
	return "examples"
}

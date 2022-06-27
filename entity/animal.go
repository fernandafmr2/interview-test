package entity

type Animal struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name       string `gorm:"type:varchar(255) not null" json:"name"`
	Class string `gorm:"type:text not null" json:"class"`
	Legs      uint64 `gorm:"not null" json:"legs"`
}

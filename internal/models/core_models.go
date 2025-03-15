package models

type Category struct {
	Name string `gorm:"unique;size:20"`
	Icon string `gorm:"type:text"`
}

type Budget struct {
	Name     string   `gorm:"unique;size:20"`
	Icon     string   `gorm:"type:text"`
	Category Category `gorm:"foreignKey:CategoryID"`
}

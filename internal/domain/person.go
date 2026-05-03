package domain

type Person struct {
	BaseEntity
	FullName    string `json:"full_name" gorm:"size:255;not null"`
	PhoneNumber string `json:"phone_number" gorm:"size:20;unique;not null"`
}

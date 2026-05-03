package domain

type ShippingAddress struct {
	BaseEntity
	UserID         string `json:"user_id" gorm:"type:uuid;index"`
	AddressName    string `json:"address_name" gorm:"size:100"`
	PhoneNumber    string `json:"phone_number" gorm:"size:20"`
	City           string `json:"city" gorm:"size:100"`
	Area           string `json:"area" gorm:"size:100"`
	StreetName     string `json:"street_name" gorm:"size:255"`
	BuildingNo     string `json:"building_no" gorm:"size:50"`
	Floor          string `json:"floor" gorm:"size:50"`
	AdditionalInfo string `json:"additional_info" gorm:"type:text"`
}

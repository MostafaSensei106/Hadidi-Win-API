package domain

type ShippingAddress struct {
	AddressName    string `json:"address_name"`
	PhoneNumber    string `json:"phone_number"`
	City           string `json:"city"`
	Area           string `json:"area"`
	StreetName     string `json:"street_name"`
	BuildingNo     string `json:"building_no"`
	Floor          string `json:"floor"`
	AdditionalInfo string `json:"additional_info"`
}

type ShippingAddresses struct {
	Addresses []ShippingAddress `json:"addresses"`
}

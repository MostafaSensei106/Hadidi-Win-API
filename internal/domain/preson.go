package domain

type Person struct {
	ID          string `json:"id"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}

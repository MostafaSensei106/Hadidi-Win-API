package domain

type Admin struct {
	Person
	Role string `json:"role" gorm:"default:'admin'"`
}

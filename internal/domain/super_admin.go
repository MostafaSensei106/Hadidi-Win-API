package domain

type SuperAdmin struct {
	Person
	Role string `json:"role" gorm:"default:'super_admin'"`
}

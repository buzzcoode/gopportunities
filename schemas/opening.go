package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role        string
	Company     string
	Description string
	Location    string
	Remote      bool
	Link        string
	Salary      int64
}

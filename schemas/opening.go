package schemas

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string `json:"name"`
	LastName      string `json:"lastname"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	F_inscripcion string `json:"fecha_inscripcion"`
	Titular       *bool  `json:"titular"`
	CI            string `json:"ci"`
}

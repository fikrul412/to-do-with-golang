
package models

import (
	"time"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
	CategoryID  uint       `json:"category_id"`
	Category    Category   `json:"category" gorm:"foreignKey:CategoryID"`
	Priority    string     `json:"priority"`       
	DueDate     *time.Time `json:"due_date"`      
}

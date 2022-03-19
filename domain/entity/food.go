package entity

import (
	"html"
	"strings"
	"time"
)

type Food struct {
	ID          uint64     `json:"primary_key;auto_increment" json:"id"`
	UserID      uint64     `json:"size:100;not null;" json:"user_id"`
	Title       string     `json:"size:100;not null;unique" json:"title"`
	Description string     `json:"text;not null;" json:"description"`
	FoodImage   string     `json:"size:255;null;" json:"food_image"`
	CreatedAt   time.Time  `json:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `json:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func (f *Food) BeforeSave() {
	f.Title = html.EscapeString(strings.TrimSpace(f.Title))
}

func (f *Food) Prepare() {
	f.Title = html.EscapeString(strings.TrimSpace(f.Title))
	f.CreatedAt = time.Now()
	f.UpdatedAt = time.Now()
}

func (f *Food) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)

	switch strings.ToLower(action) {
	case "update":
		if f.Title == "" || f.Title == "null" {
			errorMessages["title_required"] = "title is required"
		}
		if f.Description == "" || f.Description == "null" {
			errorMessages["desc_required"] = "description is required"
		}
	default:
		if f.Title == "" || f.Title == "null" {
			errorMessages["title_required"] = "title is required"
		}
		if f.Description == "" || f.Description == "null" {
			errorMessages["desc_required"] = "description is required"
		}
	}
	return errorMessages
}

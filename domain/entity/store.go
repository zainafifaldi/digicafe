package entity

import (
	"html"
	"strings"
	"time"
)

type StoreType int32

const (
	main StoreType = iota + 1
	branch
)

type Store struct {
	ID                uint64     `json:"id" db:"id"`
	Type              StoreType  `json:"type" db:"type"`
	Name              string     `json:"name" db:"name"`
	Description       string     `json:"description" db:"description"`
	Address           string     `json:"address" db:"address"`
	Active            bool       `json:"active" db:"active"`
	LocationLatitude  float64    `json:"location_latitude" db:"location_latitude"`
	LocationLongitude float64    `json:"location_longitude" db:"location_longitude"`
	CreatedAt         time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at" db:"deleted_at"`
}

func getAvailableStoreTypes() []StoreType {
	return []StoreType{main, branch}
}

func (s *Store) BeforeSave() {
	s.Description = html.EscapeString(strings.TrimSpace(s.Description))
	s.Address = html.EscapeString(strings.TrimSpace(s.Address))
}

func (s *Store) Prepare() {
	s.Description = html.EscapeString(strings.TrimSpace(s.Description))
	s.Address = html.EscapeString(strings.TrimSpace(s.Address))
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
}

func (s *Store) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)

	switch strings.ToLower(action) {
	case "create":
		if s.Name == "" || s.Name == "null" {
			errorMessages["name_required"] = "Name is required"
		}
		if !memberOfType(s.Type) {
			errorMessages["type_invalid"] = "Type is invalid"
		}
	case "update":
		if s.Name == "" || s.Name == "null" {
			errorMessages["name_required"] = "Name is required"
		}
		if !memberOfType(s.Type) {
			errorMessages["type_invalid"] = "Type is invalid"
		}
	}
	return errorMessages
}

func memberOfType(storeType StoreType) bool {
	for _, availableType := range getAvailableStoreTypes() {
		if storeType == availableType {
			return true
		}
	}

	return false
}

package model

import (
	"time"

	"github.com/google/uuid"
)

type Provider string

const (
	ProviderGoogle Provider = "google"
	ProviderEmail Provider = "email"
)

type User struct {
	ID 					uuid.UUID 	`json:"id"`
	Email 				string 		`json:"email"`
	PasswordHash 		string 		`json:"password_hash"`
	Name 				string 		`json:"name"`
	Provider 			Provider 	`json:"provider"`
	ProviderID 			string 		`json:"provider_id"`
	CreatedAt 			time.Time 	`json:"created_at"`
	UpdatedAt 			time.Time 	`json:"updated_at"`
	// DeletedAt 			time.Time `json:"deleted_at"`
}
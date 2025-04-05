package domain

import "time"

type Schedule struct {
}

type User struct {
	ID         int64
	Provider   string    // Social login provider
	ProviderID string    // User ID provided by the social login provider
	Email      string    // Email address (may not be available from some providers)
	Name       string    // name or nickname
	CreatedAt  time.Time // Timestamp of account creation
	UpdatedAt  time.Time // Timestamp of the last update
}

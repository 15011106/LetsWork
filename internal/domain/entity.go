package domain

import "time"

type Schedule struct {
	Id        int64
	StartTime time.Time
	EndTime   time.Time
	Memo      string
	UserId    int64
}

type User struct {
	Id         int64
	Provider   string    // Social login provider
	ProviderId string    // User ID provided by the social login provider
	Email      string    // Email address (may not be available from some providers)
	Name       string    // name or nickname
	CreatedAt  time.Time // Timestamp of account creation
	UpdatedAt  time.Time // Timestamp of the last update
}

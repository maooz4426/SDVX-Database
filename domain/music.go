package domain

import "time"

type Music struct {
	ID        int
	MusicName string
	Composer  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

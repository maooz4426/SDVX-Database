package domain

import "time"

type Level struct {
	ID         int
	LevelName  string
	LevelValue int
	SongID     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

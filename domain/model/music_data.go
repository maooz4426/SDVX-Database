package model

import "time"

//扱うモデルを定義

type MusicData struct {
	MusicName  string
	Composer   string
	LevelName  string
	LevelValue int
}

type Music struct {
	ID        int
	MusicName string
	Composer  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Level struct {
	ID         int
	LevelName  string
	LevelValue int
	SongID     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

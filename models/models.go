package models

import "time"

type Game struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	CoverImage   string   `json:"coverImage"`
	Tags         []string `json:"tags"`
	Difficulties []string `json:"difficulties"` // 简单、中等、复杂
}

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"createdAt"`
}

type GameRecord struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	GameID    string    `json:"gameId"`
	Score     int       `json:"score"`
	Duration  int       `json:"duration"` // seconds
	PlayedAt  time.Time `json:"playedAt"`
	Result    string    `json:"result"` // "win" | "lose" | "complete"
}

type UserStats struct {
	UserID     string     `json:"userId"`
	UserName   string     `json:"userName"`
	TotalGames int        `json:"totalGames"`
	TotalTime  int        `json:"totalTime"`
	TotalScore int        `json:"totalScore"`
	GameStats  []GameStat `json:"gameStats"`
}

type GameStat struct {
	GameID    string `json:"gameId"`
	GameName  string `json:"gameName"`
	PlayCount int    `json:"playCount"`
	BestScore int    `json:"bestScore"`
	TotalTime int    `json:"totalTime"`
}

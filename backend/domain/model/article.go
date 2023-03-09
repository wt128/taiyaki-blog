package model

import (
	"time"

	"github.com/uptrace/bun"
)


type Article struct {
	bun.BaseModel `bun:"table:article,alias:ar"`
	ID int `json:"id" bun:"id,pk,autoincrement"`
	UserId int `json:"userId" bun:"user_id"`
	Title string `json:"title" bun:"title"`
	Content string `json:"content" bun:"content"`
	Explain string `json:"explain" bun:"explain"`
	CreatedAt time.Time `json:"createdAt" bun:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" bun:"updated_at"`
}



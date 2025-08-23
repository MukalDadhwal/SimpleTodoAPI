package api

import "time"

type Todo struct{
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Completed bool      `json:"completed" default:"false"`
    CreatedAt time.Time `json:"created_at"`
}
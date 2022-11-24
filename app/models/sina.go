package models

import (
	"gorm.io/gorm"
	"time"
)

type Sina struct {
	Id          int64          `json:"id"`
	Symbol      string         `json:"symbol"`
	Bofu        float64        `json:"bofu"`
	Company     string         `json:"company"`
	Date        time.Time      `json:"date"`
	Fullcompany string         `json:"fullcompany"`
	Higner      string         `json:"higner"`
	Lower       string         `json:"lower"`
	Name        string         `json:"name"`
	Open        string         `json:"open"`
	Percent     string         `json:"percent"`
	Price       float64        `json:"price"`
	Spj         string         `json:"spj"`
	Updown      string         `json:"updown"`
	Updowncount float64        `json:"updownCount"`
	Zf          string         `json:"zf"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

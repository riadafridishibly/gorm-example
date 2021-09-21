package models

import (
	"time"

	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	Name     string
	Released time.Time
	Artists  []*Artist
}

type Album struct {
	gorm.Model
	Name    string
	Band    *Band
	Artists *[]Artist
}

type Artist struct {
	gorm.Model
	Name  string
	Songs []*Song
}

type Band struct {
	gorm.Model
	Artists []*Artist
	Songs   []*Song
}

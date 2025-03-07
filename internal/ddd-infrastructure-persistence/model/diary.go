package model

import (
	"time"

	"github.com/hedon-go-road/ddd-archetype-go/pkg/ddd"
)

type Diary struct {
	ddd.DomainModel

	ID      string    `gorm:"column:id;primaryKey"`
	UID     string    `gorm:"column:uid;not null"`
	Date    time.Time `gorm:"column:date;not null"`
	Content string    `gorm:"column:content"`
}

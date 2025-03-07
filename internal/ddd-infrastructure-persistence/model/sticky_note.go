package model

import (
	"time"

	"github.com/hedon-go-road/ddd-archetype-go/pkg/ddd"
)

type StickyNote struct {
	ddd.DomainModel

	ID             string    `gorm:"column:id;primaryKey"`
	UID            string    `gorm:"column:uid;not null"`
	DiaryID        string    `gorm:"column:diary_id;not null"`
	Participants   []string  `gorm:"column:participants"`
	OccurrenceTime time.Time `gorm:"column:occurrence_time"`
	Content        string    `gorm:"column:content"`
}

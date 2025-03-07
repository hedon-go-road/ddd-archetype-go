package model

import (
	"time"

	"github.com/hedon-go-road/ddd-archetype-go/pkg/ddd"
)

type StickyNote struct {
	ddd.DomainModel

	EntityID       string    `gorm:"column:entity_id;type:varchar(36);unique;not null"`
	UID            string    `gorm:"column:uid;type:varchar(36);not null"`
	DiaryID        string    `gorm:"column:diary_id;type:varchar(36);not null"`
	Participants   string    `gorm:"column:participants;type:varchar(256)"`
	OccurrenceTime time.Time `gorm:"column:occurrence_time;not null"`
	Content        string    `gorm:"column:content;type:varchar(1024)"`
}

func (s *StickyNote) TableName() string {
	return "t_sticky_note"
}

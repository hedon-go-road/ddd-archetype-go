package model

import (
	"time"

	"github.com/hedon-go-road/ddd-archetype-go/pkg/ddd"
)

type Diary struct {
	ddd.DomainModel

	EntityID string    `gorm:"column:entity_id;type:varchar(36);unique;not null"`
	UID      string    `gorm:"column:uid;type:varchar(32);not null"`
	Date     time.Time `gorm:"column:date;not null"`
	Content  string    `gorm:"column:content;type:varchar(1024)"`
}

func (d *Diary) TableName() string {
	return "t_diary"
}

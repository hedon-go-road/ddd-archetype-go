package domain

import (
	"time"

	"github.com/hedon-go-road/ddd-archetype-go/pkg/ddd"
)

type DiaryID string

func (d DiaryID) Value() string {
	return string(d)
}

type Diary struct {
	ddd.DomainMask

	// EntityID is the ID of the diary.
	EntityID DiaryID
	// UID is the UID of the diary.
	UID string
	// Date is the date of the diary.
	Date time.Time
	// Content is the content of the diary.
	Content string
}

func (d *Diary) Create() error {
	// TODO: implement
	return nil
}

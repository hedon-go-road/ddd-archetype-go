package domain

import (
	"context"
	"time"

	"github.com/hedon-go-road/ddd-archetype-go/pkg/ddd"
	"github.com/hedon-go-road/ddd-archetype-go/pkg/log"
)

type DiaryID string

func (d DiaryID) Value() string {
	return string(d)
}

type Diary struct {
	ddd.DomainMask

	// EntityID is the ID of the diary.
	EntityID DiaryID
	// UID is the diary's owner's unique id.
	UID string
	// Date is the date of the diary.
	Date time.Time
	// Content is the content of the diary.
	Content string
}

// Create does the diary creation.
func (d *Diary) Create(ctx context.Context) error {
	log.Ctx(ctx).Info().Msg("create diary")
	return nil
}

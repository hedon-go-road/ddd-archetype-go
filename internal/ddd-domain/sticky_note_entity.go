package domain

import (
	"context"
	"time"

	"github.com/hedon-go-road/ddd-archetype-go/pkg/ddd"
	"github.com/hedon-go-road/ddd-archetype-go/pkg/log"
)

type StickyNoteID string

func (s StickyNoteID) Value() string {
	return string(s)
}

type StickyNote struct {
	ddd.DomainMask

	// EntityID is the ID of the sticky note.
	EntityID StickyNoteID
	// DiaryEntityID is the ID of the diary entity.
	DiaryEntityID DiaryID
	// UID is the sticky note's owner's unique id.
	UID string
	// Participants is the participants of the sticky note.
	Participants []string
	// OccurrenceTime is the occurrence time of the sticky note.
	OccurrenceTime time.Time
	// Content is the content of the sticky note.
	Content string
}

// Modify modifies the content and participants of the sticky note.
func (s *StickyNote) Modify(content string, participants []string) {
	s.Content = content
	s.Participants = participants
}

// CompleteCreate completes the sticky note creation.
func (s *StickyNote) CompleteCreate(ctx context.Context) error {
	log.Ctx(ctx).Info().Msg("complete sticky note creation")
	return nil
}

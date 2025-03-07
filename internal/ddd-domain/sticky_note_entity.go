package domain

import (
	"time"

	"github.com/hedon-go-road/ddd-archetype-go/pkg/ddd"
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
	// UID is the UID of the sticky note.
	UID string
	// Participants is the participants of the sticky note.
	Participants []string
	// OccurrenceTime is the occurrence time of the sticky note.
	OccurrenceTime time.Time
	// Content is the content of the sticky note.
	Content string
}

func (s *StickyNote) Modify(content string, participants []string) {
	s.Content = content
	s.Participants = participants
}

func (s *StickyNote) CompleteCreate() {
	// TODO: implement
}

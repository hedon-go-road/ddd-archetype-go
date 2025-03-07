package domain

// StickyNoteEntityFactory is a factory for creating sticky note entities.
type StickyNoteEntityFactory interface {
	// NewInstance creates a new sticky note entity.
	// uid is the UID of the sticky note.
	// diaryEntityID is the ID of the diary entity.
	// content is the content of the sticky note.
	// occurrenceTimeStr is the occurrence time of the sticky note.
	// participants is the participants of the sticky note.
	NewInstance(uid, diaryEntityID, content, occurrenceTimeStr string, participants []string) (StickyNote, error)
}

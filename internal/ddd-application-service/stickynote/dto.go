package stickynote

type CreateCommand struct {
	UID               string   `json:"uid" binding:"required"`
	DiaryID           string   `json:"diary_id" binding:"required"`
	OccurrenceTimeStr string   `json:"occurrence_time_str" binding:"required"`
	Content           string   `json:"content" binding:"required"`
	Participants      []string `json:"participants" binding:"required"`
}

type CreateView struct {
	StickyNoteID string `json:"sticky_note_id"`
}

type GenerateContentCommand struct {
	DiaryID string `json:"diary_id" binding:"required"`
}

type GenerateContentView struct {
	GeneratedContent string `json:"generated_content"`
}

type ModifyCommand struct {
	StickyNoteID string   `json:"sticky_note_id" binding:"required"`
	Content      string   `json:"content" binding:"required"`
	Participants []string `json:"participants" binding:"required"`
}

type Query struct{}

type View struct {
	UID               string   `json:"uid"`
	DiaryID           string   `json:"diary_id"`
	Content           string   `json:"content"`
	Participants      []string `json:"participants"`
	OccurrenceTimeStr string   `json:"occurrence_time_str"`
}

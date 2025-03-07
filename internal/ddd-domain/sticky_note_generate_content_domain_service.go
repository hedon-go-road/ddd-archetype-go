package domain

import "context"

// StickyNoteGenerateContentDomainService is a domain service for generating content
// according to sticky notes which are below the diary.
type StickyNoteGenerateContentDomainService interface {
	// GenerateContent generates a content for a diary.
	GenerateContent(ctx context.Context, diaryEntityID DiaryID) (string, error)
}

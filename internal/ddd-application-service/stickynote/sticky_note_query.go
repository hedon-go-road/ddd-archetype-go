package stickynote

import (
	"context"

	domain "github.com/hedon-go-road/ddd-archetype-go/internal/ddd-domain"
)

type StickyNoteQueryService struct {
	genSvc domain.StickyNoteGenerateContentDomainService
}

func NewStickyNoteQueryService(genSvc domain.StickyNoteGenerateContentDomainService) *StickyNoteQueryService {
	return &StickyNoteQueryService{genSvc: genSvc}
}

func (s *StickyNoteQueryService) GenerateDiaryContent(ctx context.Context, cmd GenerateContentCommand) (GenerateContentView, error) {
	content, err := s.genSvc.GenerateContent(ctx, domain.DiaryID(cmd.DiaryID))
	if err != nil {
		return GenerateContentView{}, err
	}
	return GenerateContentView{GeneratedContent: content}, nil
}

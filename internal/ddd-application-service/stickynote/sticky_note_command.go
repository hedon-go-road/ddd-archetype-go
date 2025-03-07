package stickynote

import (
	"context"

	domain "github.com/hedon-go-road/ddd-archetype-go/internal/ddd-domain"
)

type StickyNoteCommandService struct {
	factory domain.StickyNoteEntityFactory
	repo    domain.StickyNoteEntityRepo
}

func NewStickyNoteCommandService(factory domain.StickyNoteEntityFactory, repo domain.StickyNoteEntityRepo) *StickyNoteCommandService {
	return &StickyNoteCommandService{factory: factory, repo: repo}
}

func (s *StickyNoteCommandService) CreateStickyNote(ctx context.Context, cmd CreateCommand) (CreateView, error) {
	stickyNote, err := s.factory.NewInstance(
		cmd.UID,
		cmd.DiaryID,
		cmd.Content,
		cmd.OccurrenceTimeStr,
		cmd.Participants,
	)
	if err != nil {
		return CreateView{}, err
	}
	err = s.repo.Save(ctx, stickyNote)
	if err != nil {
		return CreateView{}, err
	}
	return CreateView{StickyNoteID: stickyNote.EntityID.Value()}, nil
}

func (s *StickyNoteCommandService) ModifyStickyNote(ctx context.Context, cmd ModifyCommand) error {
	stickyNote, err := s.repo.Load(ctx, domain.StickyNoteID(cmd.StickyNoteID))
	if err != nil {
		return err
	}
	stickyNote.Modify(cmd.Content, cmd.Participants)
	return s.repo.Save(ctx, stickyNote)
}

package service

import (
	"context"
	"fmt"

	"github.com/golang-module/carbon"

	domain "github.com/hedon-go-road/ddd-archetype-go/internal/ddd-domain"
	"github.com/hedon-go-road/ddd-archetype-go/internal/ddd-infrastructure-persistence/db"
)

type GenerateContentDomainServiceImpl struct {
	openAIGateway domain.OpenAIGateway
	stickyNoteDB  *db.StickyNoteDB
}

func NewGenerateContentDomainServiceImpl(
	openAIGateway domain.OpenAIGateway,
	stickyNoteDB *db.StickyNoteDB,
) *GenerateContentDomainServiceImpl {
	return &GenerateContentDomainServiceImpl{
		openAIGateway: openAIGateway,
		stickyNoteDB:  stickyNoteDB,
	}
}

func (s *GenerateContentDomainServiceImpl) GenerateContent(ctx context.Context, diaryEntityID domain.DiaryID) (string, error) {
	noteList, err := s.stickyNoteDB.GetListByDiaryID(ctx, diaryEntityID)
	if err != nil {
		return "", err
	}
	if len(noteList) == 0 {
		return "empty", nil
	}

	uid := noteList[0].UID
	content := ""
	for _, note := range noteList {
		content += fmt.Sprintf("time: %s, participants: %v, content: %s; ",
			note.OccurrenceTime.Format(carbon.DateLayout),
			note.Participants,
			note.Content,
		)
	}

	return s.openAIGateway.GenerateContent(ctx, uid, content)
}

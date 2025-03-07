package stickynote

import (
	"fmt"
	"time"

	"github.com/golang-module/carbon"
	"github.com/google/uuid"

	domain "github.com/hedon-go-road/ddd-archetype-go/internal/ddd-domain"
)

type StickyNoteEntityFactoryImpl struct{}

// NewInstance creates a new sticky note entity.
// occurrenceTimeStr is the occurrence time of the sticky note,
// it should be in the format of "2021-01-01 00:00:00".
func (s *StickyNoteEntityFactoryImpl) NewInstance(uid string, diaryEntityID string, content string, occurrenceTimeStr string, participants []string) (domain.StickyNote, error) {
	parseResult := carbon.Parse(occurrenceTimeStr)
	if parseResult.Error != nil {
		return domain.StickyNote{}, fmt.Errorf("invalid occurrence time: %w", parseResult.Error)
	}

	res := domain.StickyNote{
		EntityID:       domain.StickyNoteID(uuid.New().String()),
		DiaryEntityID:  domain.DiaryID(diaryEntityID),
		UID:            uid,
		Content:        content,
		OccurrenceTime: parseResult.ToStdTime(),
		Participants:   participants,
	}
	(&res).Init(uid, time.Now())
	return res, nil
}

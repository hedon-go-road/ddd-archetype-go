package db

import (
	"context"
	"encoding/json"

	"github.com/samber/lo"
	"gorm.io/gorm"

	domain "github.com/hedon-go-road/ddd-archetype-go/internal/ddd-domain"
	"github.com/hedon-go-road/ddd-archetype-go/internal/ddd-infrastructure-persistence/model"
)

// StickyNoteDB is a database for sticky note.
// TODO: it can be designed as an interface.
type StickyNoteDB struct {
	db *gorm.DB
}

func NewStickyNoteDB(db *gorm.DB) *StickyNoteDB {
	if !db.Migrator().HasTable(&model.StickyNote{}) {
		if err := db.Migrator().CreateTable(&model.StickyNote{}); err != nil {
			panic("failed to create sticky note table: " + err.Error())
		}
	}
	return &StickyNoteDB{
		db: db,
	}
}

func (d *StickyNoteDB) GetByEntityID(ctx context.Context, id domain.StickyNoteID) (domain.StickyNote, error) {
	var model model.StickyNote
	if err := d.db.WithContext(ctx).Where("entity_id = ?", id).First(&model).Error; err != nil {
		return domain.StickyNote{}, err
	}
	return d.toEntity(&model), nil
}

func (d *StickyNoteDB) GetListByDiaryID(ctx context.Context, diaryID domain.DiaryID) ([]domain.StickyNote, error) {
	var models []model.StickyNote
	if err := d.db.WithContext(ctx).Where("diary_id = ?", diaryID).Find(&models).Error; err != nil {
		return nil, err
	}
	return lo.Map(models, func(model model.StickyNote, _ int) domain.StickyNote {
		return d.toEntity(&model)
	}), nil
}

func (d *StickyNoteDB) Save(ctx context.Context, entity domain.StickyNote) error {
	return d.db.WithContext(ctx).Save(d.toModel(&entity)).Error
}

func (d *StickyNoteDB) toModel(entity *domain.StickyNote) *model.StickyNote {
	participants, _ := json.Marshal(entity.Participants)
	return &model.StickyNote{
		DomainModel:    entity.ToModel(),
		EntityID:       entity.EntityID.Value(),
		UID:            entity.UID,
		DiaryID:        entity.DiaryEntityID.Value(),
		Participants:   string(participants),
		OccurrenceTime: entity.OccurrenceTime,
		Content:        entity.Content,
	}
}

func (d *StickyNoteDB) toEntity(model *model.StickyNote) domain.StickyNote {
	var participants []string
	_ = json.Unmarshal([]byte(model.Participants), &participants)
	return domain.StickyNote{
		DomainMask:     model.ToMask(),
		EntityID:       domain.StickyNoteID(model.EntityID),
		UID:            model.UID,
		DiaryEntityID:  domain.DiaryID(model.DiaryID),
		Participants:   participants,
		OccurrenceTime: model.OccurrenceTime,
		Content:        model.Content,
	}
}

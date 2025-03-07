package db

import (
	"context"

	"github.com/samber/lo"
	"gorm.io/gorm"

	domain "github.com/hedon-go-road/ddd-archetype-go/internal/ddd-domain"
	model "github.com/hedon-go-road/ddd-archetype-go/internal/ddd-infrastructure-persistence/model"
)

// DiaryDB is a database for diary.
// TODO: it can be designed as an interface.
type DiaryDB struct {
	db *gorm.DB
}

func NewDiaryDB(db *gorm.DB) *DiaryDB {
	return &DiaryDB{db: db}
}

func (d *DiaryDB) PageList(ctx context.Context, lastID string, pageSize int) ([]domain.Diary, error) {
	var models []model.Diary
	if err := d.db.WithContext(ctx).
		Where("id > ?", lastID).
		Limit(pageSize).
		Find(&models).Error; err != nil {
		return nil, err
	}
	return lo.Map(models, func(model model.Diary, _ int) domain.Diary {
		return d.toEntity(&model)
	}), nil
}

func (d *DiaryDB) CountForPageList(ctx context.Context) (int64, error) {
	var count int64
	if err := d.db.WithContext(ctx).Model(&model.Diary{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (d *DiaryDB) GetByID(ctx context.Context, id domain.DiaryID) (domain.Diary, error) {
	var model model.Diary
	if err := d.db.WithContext(ctx).Where("id = ?", id).First(&model).Error; err != nil {
		return domain.Diary{}, err
	}
	return d.toEntity(&model), nil
}

func (d *DiaryDB) Save(ctx context.Context, entity domain.Diary) error {
	return d.db.WithContext(ctx).Save(d.toModel(&entity)).Error
}

func (d *DiaryDB) toEntity(model *model.Diary) domain.Diary {
	return domain.Diary{
		DomainMask: model.ToMask(),
		EntityID:   domain.DiaryID(model.ID),
		UID:        model.UID,
		Date:       model.Date,
		Content:    model.Content,
	}
}

func (d *DiaryDB) toModel(entity *domain.Diary) *model.Diary {
	return &model.Diary{
		DomainModel: entity.ToModel(),
		ID:          entity.EntityID.Value(),
		UID:         entity.UID,
		Date:        entity.Date,
		Content:     entity.Content,
	}
}

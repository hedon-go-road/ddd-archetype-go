package diary

import (
	"fmt"
	"time"

	"github.com/golang-module/carbon"
	"github.com/google/uuid"

	domain "github.com/hedon-go-road/ddd-archetype-go/internal/ddd-domain"
)

type DiaryEntityFactoryImpl struct{}

func (d *DiaryEntityFactoryImpl) NewInstance(uid string, date string) (domain.Diary, error) {
	// do a lot of check
	if uid == "" {
		return domain.Diary{}, fmt.Errorf("uid is empty")
	}
	res := carbon.ParseByFormat(date, carbon.DateLayout)
	if res.Error != nil {
		return domain.Diary{}, fmt.Errorf("invalid date format: %w", res.Error)
	}

	// create a new diary entity
	diary := domain.Diary{
		EntityID: domain.DiaryID(uuid.NewString()),
		UID:      uid,
		Date:     res.ToStdTime(),
		Content:  "",
	}
	(&diary).Init(uid, time.Now())
	return diary, nil
}

package diary

import (
	"context"

	"github.com/golang-module/carbon"
	"github.com/samber/lo"

	domain "github.com/hedon-go-road/ddd-archetype-go/internal/ddd-domain"
	"github.com/hedon-go-road/ddd-archetype-go/internal/ddd-infrastructure-persistence/db"
	"github.com/hedon-go-road/ddd-archetype-go/pkg/ddd"
)

type DiaryQueryService struct {
	db *db.DiaryDB
}

func NewDiaryQueryService(db *db.DiaryDB) *DiaryQueryService {
	return &DiaryQueryService{db: db}
}

func (s *DiaryQueryService) PageList(ctx context.Context, query Query) (ddd.PageList[QueryView], error) {
	count, err := s.db.CountForPageList(ctx)
	if err != nil {
		return ddd.PageList[QueryView]{}, err
	}
	diaryList, err := s.db.PageList(ctx, query.LastID, query.PageSize)
	if err != nil {
		return ddd.PageList[QueryView]{}, err
	}
	res := ddd.PageList[QueryView]{
		Total:    int(count),
		PageSize: len(diaryList),
		List: lo.Map(diaryList, func(diary domain.Diary, _ int) QueryView {
			return QueryView{
				ID:           diary.ID,
				DiaryID:      diary.EntityID.Value(),
				Content:      diary.Content,
				DiaryDateStr: diary.Date.Format(carbon.DateLayout),
				UID:          diary.UID,
			}
		}),
	}
	return res, nil
}

package diary

import (
	"context"

	domain "github.com/hedon-go-road/ddd-archetype-go/internal/ddd-domain"
)

type DiaryCommandService struct {
	factory domain.DiaryEntityFactory
	repo    domain.DiaryEntityRepo
}

func NewDiaryCommandService(factory domain.DiaryEntityFactory, repo domain.DiaryEntityRepo) *DiaryCommandService {
	return &DiaryCommandService{factory: factory, repo: repo}
}

func (s *DiaryCommandService) CreateDiary(ctx context.Context, cmd CreateCommand) (CreateView, error) {
	diary, err := s.factory.NewInstance(cmd.UID, cmd.DiaryDate)
	if err != nil {
		return CreateView{}, err
	}

	if err := diary.Create(ctx); err != nil {
		return CreateView{}, err
	}

	if err := s.repo.Save(ctx, diary); err != nil {
		return CreateView{}, err
	}
	return CreateView{DiaryID: diary.EntityID.Value()}, nil
}

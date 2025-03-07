package domain

import "github.com/hedon-go-road/ddd-archetype-go/pkg/ddd"

type DiaryEntityRepo interface {
	ddd.DomainRepo[DiaryID, string, Diary]
}

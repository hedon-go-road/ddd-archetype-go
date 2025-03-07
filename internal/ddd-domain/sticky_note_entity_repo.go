package domain

import "github.com/hedon-go-road/ddd-archetype-go/pkg/ddd"

type StickyNoteEntityRepo interface {
	ddd.DomainRepo[StickyNoteID, string, StickyNote]
}

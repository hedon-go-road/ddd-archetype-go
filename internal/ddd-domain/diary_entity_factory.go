package domain

// DiaryEntityFactory is a factory for creating diary entities.
type DiaryEntityFactory interface {
	// NewInstance creates a new diary entity.
	NewInstance(uid string, date string) (Diary, error)
}

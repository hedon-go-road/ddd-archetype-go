package diary

type CreateCommand struct {
	DiaryDate string `json:"diary_date" binding:"required"`
	UID       string `json:"uid" binding:"required"`
}

type CreateView struct {
	DiaryID string `json:"diary_id"`
}

type Query struct {
	LastID   int64 `json:"last_id"`
	PageSize int   `json:"page_size" binding:"required,gte=1,lte=100"`
}

type QueryView struct {
	ID           int64  `json:"id"`
	DiaryID      string `json:"diary_id"`
	Content      string `json:"content"`
	DiaryDateStr string `json:"diary_date_str"`
	UID          string `json:"uid"`
}

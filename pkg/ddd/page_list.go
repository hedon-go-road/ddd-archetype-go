package ddd

// PageList is a struct that contains the metadata for a page list.
type PageList[T any] struct {
	// PageSize is the number of items per page.
	PageSize int
	// Total is the total number of items.
	Total int
	// List is the list of items.
	List []T
}

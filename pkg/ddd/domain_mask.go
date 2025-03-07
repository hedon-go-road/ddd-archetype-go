package ddd

import "time"

const (
	DefaultVersion = 1
)

// DomainMask is a struct that contains the metadata for a domain object.
type DomainMask struct {
	// ID is the auto-incremented ID from db.
	ID int64
	// Deleted is a flag indicating if the domain object has been deleted.
	Deleted bool
	// CreateBy is the user who created the domain object.
	CreateBy string
	// CreateAt is the timestamp when the domain object was created.
	CreateAt time.Time
	// UpdateBy is the user who last updated the domain object.
	UpdateBy string
	// UpdateAt is the timestamp when the domain object was last updated.
	UpdateAt time.Time
	// Version is the version of the domain object, used for optimistic locking.
	Version int64
}

func (d *DomainMask) Init(by string, now time.Time) {
	d.Deleted = false
	d.CreateBy = by
	d.CreateAt = now
	d.UpdateBy = by
	d.UpdateAt = now
	d.Version = DefaultVersion
}

func (d *DomainMask) ToModel() DomainModel {
	return DomainModel{
		ID:       d.ID,
		Deleted:  d.Deleted,
		CreateBy: d.CreateBy,
		CreateAt: d.CreateAt,
		UpdateBy: d.UpdateBy,
		UpdateAt: d.UpdateAt,
		Version:  d.Version,
	}
}

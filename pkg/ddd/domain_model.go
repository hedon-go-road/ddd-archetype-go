package ddd

import "time"

// DomainModel is a common model for all domain objects used for gorm.
type DomainModel struct {
	// ID is the ID of the domain object.
	ID int64 `gorm:"column:id;primaryKey;autoIncrement:true"`
	// Deleted is the deleted flag of the domain object.
	Deleted bool `gorm:"column:deleted"`
	// CreateBy is the user who created the domain object.
	CreateBy string `gorm:"column:create_by;type:varchar(36);not null"`
	// CreateAt is the timestamp when the domain object was created.
	CreateAt time.Time `gorm:"column:create_at;autoCreateTime"`
	// UpdateBy is the user who last updated the domain object.
	UpdateBy string `gorm:"column:update_by;type:varchar(36);not null"`
	// UpdateAt is the timestamp when the domain object was last updated.
	UpdateAt time.Time `gorm:"column:update_at;autoUpdateTime"`
	// Version is the version of the domain object, used for optimistic locking.
	Version int64 `gorm:"column:version"`
}

func (d *DomainModel) ToMask() DomainMask {
	return DomainMask{
		ID:       d.ID,
		Deleted:  d.Deleted,
		CreateBy: d.CreateBy,
		CreateAt: d.CreateAt,
		UpdateBy: d.UpdateBy,
		UpdateAt: d.UpdateAt,
		Version:  d.Version,
	}
}

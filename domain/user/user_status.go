package domain

type UserStatus struct {
	ID   uint64 `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);not null"`
}

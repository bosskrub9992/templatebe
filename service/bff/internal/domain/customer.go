package domain

type Customer struct {
	ID   int64  `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

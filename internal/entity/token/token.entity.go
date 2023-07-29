package token

type Token struct {
	Id      string `gorm:"column:id;unique;not null;primaryKey"`
	EventId string `gorm:"column:event_id;not null"`
	EndAt   int64  `gorm:"column:end_at;not null"`
}

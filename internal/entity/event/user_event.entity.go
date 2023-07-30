package event

type UserEvent struct {
	UserId  string `gorm:"column:user_id;not null"`
	EventId string `gorm:"column:event_id;not null"`
	TakenAt int64  `gorm:"column:taken_at;not null"`
}

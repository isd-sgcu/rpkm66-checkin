package event

type UserEvent struct {
	UserId  string `gorm:"column:user_id;index:user;not null"`
	EventId string `gorm:"column:event_id;index:event;not null"`
	Event   Event  `gorm:"references:event_id"`
	TakenAt int64  `gorm:"column:taken_at;not null"`
}

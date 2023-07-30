package event

type UserEvent struct {
	UserId  string `gorm:"column:user_id;not null"`
	EventId string `gorm:"column:event_id;not null"`
	When    int64  `gorm:"column:when;not null"`
}

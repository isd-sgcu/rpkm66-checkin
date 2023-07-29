package staff

type Staff struct {
	UserId string `gorm:"column:user_id;unique;not null;primaryKey"`
}

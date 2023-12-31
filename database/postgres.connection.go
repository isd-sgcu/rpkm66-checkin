package database

import (
	"fmt"
	"strconv"

	"github.com/isd-sgcu/rpkm66-checkin/cfgldr"
	"github.com/isd-sgcu/rpkm66-checkin/internal/entity/event"
	"github.com/isd-sgcu/rpkm66-checkin/internal/entity/namespace"
	"github.com/isd-sgcu/rpkm66-checkin/internal/entity/staff"
	"github.com/isd-sgcu/rpkm66-checkin/internal/entity/token"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(conf *cfgldr.Database) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", conf.Host, conf.User, conf.Password, conf.Name, strconv.Itoa(conf.Port))

	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	/// Migrate here
	db.AutoMigrate(namespace.Namespace{}, event.Event{}, event.UserEvent{}, staff.Staff{}, token.Token{})

	if err != nil {
		return nil, err
	}

	return
}

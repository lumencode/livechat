package util

import (
	"fmt"
	"github.com/lumencode/livechat/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Pg () (*gorm.DB, error) {

	config := config.Config().PostgreSQL

	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s port=%d dbname=%s sslmode=%s TimeZone=%s",
		config.Host,
		config.User,
		config.Password,
		config.Port,
		config.Database,
		config.SSLMode,
		config.Timezone,
	)

	return gorm.Open(postgres.Open(connectionString), &gorm.Config{})
}

package databases

import (
	"fmt"

	"github.com/sunshibao/connection"
)

var connector *connection.ExternalProcedure

func Init(host string, port uint16, username, password, database string) {
	connector = connection.NewExternalProcedure(
		connection.NewMySQLConfig(
			connection.MySQLHost(host),
			connection.MySQLPort(fmt.Sprintf("%d", port)),
			connection.MySQLUsername(username),
			connection.MySQLPassword(password),
			connection.MySQLDatabase(database),
		),
	)
}

func Migrate(object interface{}) error {
	return connection.GetMySQL().AutoMigrate(object).Error
}

func Close() {
	connector.Close()
}

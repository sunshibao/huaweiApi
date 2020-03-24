package test

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/sunshibao/connection"
)

func Release() {

	var err error
	err = connection.GetMySQL().Exec(fmt.Sprintf("DROP DATABASE %s", databaseName)).Error
	if err != nil {
		logrus.WithField("error", err).Error("drop database error")
	}

	err = connection.GetMySQL().Close()
	if err != nil {
		logrus.WithField("error", err).Error("close database error")
	}
}

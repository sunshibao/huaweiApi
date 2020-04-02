package application

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"huaweiApi/pkg/databases"
	"huaweiApi/pkg/models/huawei"
	"huaweiApi/pkg/models/user"
)

func (a *app) migrateDatabases() {

	logrus.Info("starting to migrate database")
	a.migrateDatabaseAndLogError(user.Users{})
	a.migrateDatabaseAndLogError(huawei.PaymentRecord{})
	logrus.Info("migrate database succeed")
}

func (a *app) migrateDatabaseAndLogError(object interface{}) {

	if err := databases.Migrate(object); err != nil {

		logrus.WithField("object", fmt.Sprintf("%T", object)).
			WithField("action", "migrate database").
			Error(err)
	}
}

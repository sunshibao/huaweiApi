package test

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"

	configUtils "huaweiApi/pkg/utils/config"
	_ "huaweiApi/pkg/utils/log"

	"huaweiApi/pkg/config"
	"huaweiApi/pkg/databases"
	"huaweiApi/pkg/models/user"
)

var databaseFilePath string
var databaseName string

func Init() {

	rand.Seed(time.Now().UnixNano())

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("dir = %s\n", dir)
	for {
		dir = path.Dir(dir)
		if strings.LastIndex(dir, "/pkg") <= 0 {
			break
		}
	}

	fmt.Printf("processed dir = %s\n", dir)
	configUtils.MustLoadConf(config.Config, fmt.Sprintf("%s/config/test.json", dir))

	databaseFilePath = fmt.Sprintf("%s/.test-result/%d.db", dir, rand.Uint64())
	databaseName = fmt.Sprintf("%s_%d", config.Config.MySQL.DatabaseName, rand.Uint64())

	logrus.Infof("database name: %s", databaseName)

	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8mb4&loc=Local&parseTime=true",
			config.Config.MySQL.Username,
			config.Config.MySQL.Password,
			config.Config.MySQL.Host,
			config.Config.MySQL.Port,
		))
	if err != nil {
		panic(err)
	}

	err = db.Exec(fmt.Sprintf("CREATE DATABASE `%s` character set UTF8mb4 collate utf8mb4_general_ci", databaseName)).Error
	if err != nil {
		panic(err)
	}
	_ = db.Close()

	databases.Init(
		config.Config.MySQL.Host,
		config.Config.MySQL.Port,
		config.Config.MySQL.Username,
		config.Config.MySQL.Password,
		databaseName,
	)

	objects := []interface{}{
		user.Users{},
	}
	for _, object := range objects {
		err = databases.Migrate(object)
		if err != nil {
			logrus.WithField("action", "migrate").Error(err)
		}
	}
}

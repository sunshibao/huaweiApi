package config

import (
	"time"

	"github.com/gin-gonic/gin"
)

const (
	DefaultListenAddress           = "127.0.0.1"
	DefaultListenPort              = uint16(8080)
	DefaultLogLevel                = "info"
	DefaultReadWriteTimeout        = time.Minute
	DefaultUserAuthTokenKey        = "user_token_key"
	DefaultServiceId               = 0
	DefaultEtcdAddress             = "http://127.0.0.1:2379"
	DefaultMySQLHost               = "127.0.0.1"
	DefaultMySQLPort        uint16 = 3306
	DefaultMySQLUsername           = "root"
	DefaultMySQLDatabase           = "services"
	DefaultRPCTTL                  = 30 * time.Second
	DefaultRPCInterval             = 10 * time.Second
)

// Gin Service Mode
type WebServiceMode string

const (
	WebServiceModeDebug   WebServiceMode = "debug"
	WebServiceModeRelease WebServiceMode = "release"
)

type configuration struct {
	ListenAddress  string                `json:"address"`
	RESTfulService restfulServiceSetting `json:"restful"`
	RPCService     rpcServiceSetting     `json:"rpc"`
	MySQL          mysqlSetting          `json:"mysql"`
	Etcd           etcdSetting           `json:"etcd"`
	Log            logSetting            `json:"log"`
	ServiceId      uint16                // used to distinguish between different services when highly available. no parse from configuration file, because services will use the same configuration file.
}

func (config *configuration) GetServiceId() uint16 {
	if config.ServiceId == 0 {
		return DefaultServiceId
	}
	return config.ServiceId
}

func (config *configuration) GetListenAddress() string {
	if config.ListenAddress == "" {
		return DefaultListenAddress
	}
	return config.ListenAddress
}

type restfulServiceSetting struct {
	Port             uint16             `json:"port"`
	Mode             WebServiceMode     `json:"mode"`
	ReadWriteTimeout Duration           `json:"readWriteTimeout"`
	Auth             restfulAuthSetting `json:"auth"`
}

func (service *restfulServiceSetting) GetPort() uint16 {
	if service.Port <= 0 {
		return DefaultListenPort
	}
	return service.Port
}

func (service *restfulServiceSetting) GetMode() WebServiceMode {
	if service.Mode == "" {
		return gin.ReleaseMode
	}
	return service.Mode
}

func (service *restfulServiceSetting) GetReadWriteTimeout() time.Duration {

	if service.ReadWriteTimeout.Duration == 0 {
		return DefaultReadWriteTimeout
	}
	return service.ReadWriteTimeout.Duration
}

type restfulAuthSetting struct {
	UserTokenKey string `json:"userTokenKey"`
}

func (setting *restfulAuthSetting) GetUserTokenKey() string {
	if len(setting.UserTokenKey) == 0 {
		return DefaultUserAuthTokenKey
	}
	return setting.UserTokenKey
}

type rpcServiceSetting struct {
	Port             uint16   `json:"port"`
	ReadWriteTimeout Duration `json:"readWriteTimeout"`
	TTL              Duration `json:"ttl"`
	Interval         Duration `json:"interval"`
}

func (service *rpcServiceSetting) GetPort() uint16 {
	if service.Port <= 0 {
		return DefaultListenPort
	}
	return service.Port
}

func (service *rpcServiceSetting) GetReadWriteTimeout() time.Duration {
	if service.ReadWriteTimeout.Duration == 0 {
		return DefaultReadWriteTimeout
	}
	return service.ReadWriteTimeout.Duration
}

func (service *rpcServiceSetting) GetTTL() time.Duration {
	if service.TTL.Duration == 0 {
		return DefaultRPCTTL
	}
	return service.TTL.Duration
}

func (service *rpcServiceSetting) GetInterval() time.Duration {
	if service.Interval.Duration == 0 {
		return DefaultRPCInterval
	}
	return service.Interval.Duration
}

type logSetting struct {
	Level string `json:"level"` // level: trace, debug, info, warn|warning, error, fatal, panic
}

func (log *logSetting) GetLevel() string {
	if log.Level == "" {
		return DefaultLogLevel
	}
	return log.Level
}

type mysqlSetting struct {
	Host         string `json:"host"`
	Port         uint16 `json:"port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	DatabaseName string `json:"database"`
}

func (mysql *mysqlSetting) GetHost() string {

	if len(mysql.Host) == 0 {
		return DefaultMySQLHost
	}

	return mysql.Host
}

func (mysql *mysqlSetting) GetPort() uint16 {

	if mysql.Port == 0 {
		return DefaultMySQLPort
	}

	return mysql.Port
}

func (mysql *mysqlSetting) GetUsername() string {

	if len(mysql.Username) == 0 {
		return DefaultMySQLUsername
	}

	return mysql.Username
}

func (mysql *mysqlSetting) GetPassword() string {

	return mysql.Password
}

func (mysql *mysqlSetting) GetDatabase() string {

	if len(mysql.DatabaseName) == 0 {
		return DefaultMySQLDatabase
	}

	return mysql.DatabaseName
}

type etcdSetting struct {
	Addresses []string `json:"addresses"`
}

func (etcd *etcdSetting) GetAddresses() []string {

	if len(etcd.Addresses) == 0 {
		return []string{DefaultEtcdAddress}
	}

	return etcd.Addresses
}

var (
	Config = new(configuration)
)

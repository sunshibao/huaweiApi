package application

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/transport"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"huaweiApi/pkg/config"
	"huaweiApi/pkg/constants"
	"huaweiApi/pkg/databases"
	configUtils "huaweiApi/pkg/utils/config"
	"huaweiApi/pkg/utils/idcreator"
	"huaweiApi/pkg/utils/log"
)

type app struct {
	cmd *cobra.Command

	httpHandler *gin.Engine
	httpServer  *http.Server

	etcdRegistry     registry.Registry
	rpcService       micro.Service
	rpcCancelContext context.Context
	rpcStopFunc      context.CancelFunc
	rpcStopped       chan bool

	isClosing bool
}

func NewApp() *app {
	return new(app)
}

func (a *app) run(cmd *cobra.Command, args []string) {

	a.cmd = cmd
	a.parseFlags()
	a.initService()
	a.migrateDatabases()
	a.startService()
	a.waitSignal()
}

func (a *app) startService() {
	a.startRESTfulAPIListener()
	a.startRPCListener()
}

func (a *app) initService() {

	a.initLogLevel()
	a.initRandomSeed()
	a.initSnowFlake()
	a.initMemoriesData()
	a.initRESTfulAPIHandler()
	a.initRequestLogger()
	a.setRESTfulRoutes()
	a.initRESTfulListener()
	a.initMySQL()
	a.initEtcd()
	a.initRPCStopped()
	a.initRPContext()
	a.initRPCService()
	a.setRegisterRPCControllers()
}

func (a *app) parseFlags() {

	a.loadConfig()
	a.parseParameters()
}

func (a *app) initRandomSeed() {

	logrus.Debug("init random seed")
	rand.Seed(time.Now().UnixNano())
	logrus.Debug("random seed init succeed")
}

func (a *app) initLogLevel() {

	logLevel, err := logrus.ParseLevel(config.Config.Log.GetLevel())
	if err != nil {
		logrus.Errorf("Parse log level error")
	} else {
		logrus.SetLevel(logLevel)
		logrus.Debugf("log level set: %s", config.Config.Log.GetLevel())
	}
}

func (a *app) waitSignal() {
	chanSignal := make(chan os.Signal, 1)
	signal.Notify(chanSignal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)

	for {
		select {
		case sig := <-chanSignal:
			logrus.Infof("Received signal: %d", sig)
			a.close()
			goto exit
		}
	}
exit:
	logrus.Infof("loop exited")
	return
}

func (a *app) close() {

	a.markClosing()
	a.ClearMemoryData()
	a.closeHTTPServer()
	a.closeRPCServer()
	a.closeMySQL()
}

func (a *app) loadConfig() {
	configFile, _ := a.cmd.Flags().GetString("config-file")
	if configFile == "" {
		return
	}
	configUtils.MustLoadConf(config.Config, configFile)
	logrus.Infof("ConfigFile: %s\n%+v", configFile, config.Config)
}

func (a *app) parseParameters() {

	a.parseParameterLogLevel()
	a.parseParameterServiceId()
}

func (a *app) parseParameterLogLevel() {

	var err error
	var logLevel string
	logLevel, err = a.cmd.Flags().GetString(FlagLogLevel)
	if err == nil && logLevel != "" {
		config.Config.Log.Level = logLevel
	}
	logrus.Infof("log level: %v", config.Config.Log.GetLevel())
}

func (a *app) parseParameterServiceId() {

	var err error
	var serviceId uint16
	serviceId, err = a.cmd.Flags().GetUint16(FlagServiceId)
	if err == nil && serviceId > 0 {
		config.Config.ServiceId = serviceId
	}
	logrus.Infof("serviceId: %d", config.Config.GetServiceId())
}

func (a *app) initRESTfulAPIHandler() {

	logrus.Debug("start to init restful api services handler")
	a.httpHandler = gin.Default()
	switch config.Config.RESTfulService.GetMode() {
	case config.WebServiceModeDebug:
		gin.SetMode(gin.DebugMode)
	default: // Default Release Mode
		gin.SetMode(gin.ReleaseMode)
	}
	logrus.Debug("init restful api handler services succeed")
}

func (a *app) initRequestLogger() {
	logrus.Debug("start to register log middleware")
	a.httpHandler.Use(log.ReqLoggerMiddleware())
	logrus.Debug("register log middleware succeed")
}

func (a *app) initRESTfulListener() {
	logrus.Debug("start to init restful api listener")
	// start the server，For services exposed on the public network, timeout must be set
	a.httpServer = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.Config.GetListenAddress(), config.Config.RESTfulService.GetPort()),
		Handler:      a.httpHandler,
		ReadTimeout:  config.Config.RESTfulService.GetReadWriteTimeout(),
		WriteTimeout: config.Config.RESTfulService.GetReadWriteTimeout(),
	}
	logrus.Debug("init restful api listener succeed")
}

func (a *app) startRESTfulAPIListener() {
	logrus.Infof("start server listening")
	go func() {
		err := a.httpServer.ListenAndServe()
		if err != nil && !a.isClosing {
			logrus.Errorf("listen error: %v", err)
		}
	}()
}

func (a *app) initSnowFlake() {

	logrus.Debug("start init id creator")
	idcreator.InitCreator(config.Config.GetServiceId())
	logrus.Debug("id creator init succeed")
}

func (a *app) initMemoriesData() {

}

func (a *app) markClosing() {
	a.isClosing = true
}

func (a *app) ClearMemoryData() {

}

func (a *app) closeHTTPServer() {

	logrus.Infof("closing http server")
	var err error
	err = a.httpServer.Close()
	if err != nil {
		logrus.Errorf("happened error at close http server: %v", err)
	}
	logrus.Infof("http server closed")
}

func (a *app) initMySQL() {

	databases.Init(
		config.Config.MySQL.GetHost(),
		config.Config.MySQL.GetPort(),
		config.Config.MySQL.GetUsername(),
		config.Config.MySQL.GetPassword(),
		config.Config.MySQL.GetDatabase(),
	)
}

func (a *app) closeMySQL() {
	databases.Close()
}

func (a *app) initEtcd() {

	a.etcdRegistry = etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = config.Config.Etcd.GetAddresses()
	})
}

func (a *app) initRPCService() {

	logrus.Info("starting to create rpc services")
	a.rpcService = micro.NewService(
		micro.Address(fmt.Sprintf("%s:%d", config.Config.GetListenAddress(), config.Config.RPCService.GetPort())),
		micro.Name(constants.RPCServiceName),
		micro.Registry(a.etcdRegistry),
		micro.Client(client.NewClient(
			client.RequestTimeout(config.Config.RPCService.GetReadWriteTimeout()),
			client.DialTimeout(config.Config.RPCService.GetReadWriteTimeout()),
		)),
		micro.Transport(
			transport.NewTransport(
				transport.Timeout(config.Config.RPCService.GetReadWriteTimeout()),
			),
		),
		micro.RegisterTTL(config.Config.RPCService.GetTTL()),
		micro.RegisterInterval(config.Config.RPCService.GetInterval()),
		micro.Context(a.rpcCancelContext),
		micro.Flags(a.getRPCServiceFlags()...),
	)

	logrus.Info("create rpc services succeed")

	logrus.Info("starting to init rpc services")

	// Initialise services
	a.rpcService.Init()

	logrus.Info("init rpc services succeed")
}

func (a *app) initRPContext() {

	logrus.Info("starting init rpc cancel context")
	a.rpcCancelContext = context.Background()
	a.rpcCancelContext, a.rpcStopFunc = context.WithCancel(a.rpcCancelContext)
	logrus.Info("init rpc cancel context succeed")
}

func (a *app) startRPCListener() {

	logrus.Info("starting rpc services")
	go func() {
		err := a.rpcService.Run()
		if err != nil {
			logrus.WithField("error", err).Errorf("rpc services error")
		}
		a.rpcStopped <- true
	}()
}

func (a *app) closeRPCServer() {

	logrus.Info("Stopping rpc services")
	a.rpcStopFunc()

	<-a.rpcStopped
	logrus.Info("Stopped rpc services")
}

func (a *app) initRPCStopped() {

	logrus.Info("start to init rpc stop signal channel")
	a.rpcStopped = make(chan bool)
	logrus.Info("init rpc stop signal channel succeed")
}

func (a *app) getRPCServiceFlags() []cli.Flag {

	var flags = make([]cli.Flag, 0, len(a.cmd.Flags().Args()))
	a.cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if flag.Name == "help" {
			return
		}
		flags = append(flags, &cli.StringFlag{Name: flag.Name, Usage: flag.Usage, Value: flag.DefValue})
	})
	return flags
}

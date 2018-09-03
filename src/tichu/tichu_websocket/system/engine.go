package system

import (
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"os/signal"
	"syscall"
	"github.com/Sirupsen/logrus"
)

var initializedEngine = false
var configuredEngine = false
var engineDebugMode = false

func InitEngine() {
	config()
}

func IsDebugging() bool {
	return engineDebugMode
}

func config() {
	InitConfig("./conf")
}

func InitConfig(path string) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(path)
	viper.SetConfigType("json")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		panic(err.Error())
	}

	configName := "config"

	if viper.InConfig("include") {
		oldSetting := viper.AllSettings()
		configName = viper.GetString("include")
		viper.SetConfigName(configName)
		if err = viper.ReadInConfig(); err != nil {
			panic(err.Error())
		}
		// resave old settings
		for oldKey, oldVal := range oldSetting {
			if !viper.InConfig(oldKey) {
				viper.Set(oldKey, oldVal)
			}
		}
	}

	l := &lumberjack.Logger{
		Filename:   viper.GetString("log_file"), //  /var/log/quest_server/server.log
		MaxSize:    50,	// megabytes
		MaxBackups: 5,
		MaxAge:     3, 	// days
	}

	if viper.InConfig("mode") {
		modeString := viper.GetString("mode")
		if modeString == gin.DebugMode {
			engineDebugMode = true
		}
		gin.SetMode(modeString)
	}

	if IsDebugging() {
		l.MaxAge = 1
		l.MaxSize = 10
		l.MaxBackups = 3
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetOutput(io.MultiWriter(l, os.Stdout))

	} else {
		logrus.SetLevel(logrus.InfoLevel)
		logrus.SetOutput(l)
	}
	if viper.InConfig("log_level") {
		switch viper.GetString("log_level") {
		case "debug" : logrus.SetLevel(logrus.DebugLevel)
		case "info" : logrus.SetLevel(logrus.InfoLevel)
		case "warn" : logrus.SetLevel(logrus.WarnLevel)
		case "error": logrus.SetLevel(logrus.ErrorLevel)
		}
	}

	l.Rotate()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)
	go func() {
		for {
			<-c
			l.Rotate()
		}
	}()

	timeStampFormat := new(logrus.TextFormatter)
	timeStampFormat.TimestampFormat = "2006-01-02 15:04:05"
	timeStampFormat.FullTimestamp = true
	timeStampFormat.DisableColors = false
	timeStampFormat.ForceColors = true
	logrus.SetFormatter(timeStampFormat)

	configuredEngine = true
	logrus.Info("-----------------------------------------------------------")
	logrus.Infof("Load Config: %s (%s) log_level:%v ", configName, viper.GetString("mode"), viper.GetString("log_level") )
}
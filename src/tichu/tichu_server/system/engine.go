package system

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	log "github.com/Sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"os/signal"
	"syscall"
)

var initializedEngine = false
var configuredEngine = false
var engineDebugMode = false

func IsDebugging() bool {
	return engineDebugMode
}

func InitEngine() {
	config()
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
		log.SetLevel(log.DebugLevel)
		log.SetOutput(io.MultiWriter(l, os.Stdout))

	} else {
		log.SetLevel(log.InfoLevel)
		log.SetOutput(l)
	}
	if viper.InConfig("log_level") {
		switch viper.GetString("log_level") {
		case "debug" : log.SetLevel(log.DebugLevel)
		case "info" : log.SetLevel(log.InfoLevel)
		case "warn" : log.SetLevel(log.WarnLevel)
		case "error": log.SetLevel(log.ErrorLevel)
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

	timeStampFormat := new(log.TextFormatter)
	timeStampFormat.TimestampFormat = "2006-01-02 15:04:05"
	timeStampFormat.FullTimestamp = true
	timeStampFormat.DisableColors = false
	timeStampFormat.ForceColors = true
	log.SetFormatter(timeStampFormat)

	configuredEngine = true
	log.Info("-----------------------------------------------------------")
	log.Infof("Load Config: %s (%s) log_level:%v ", configName, viper.GetString("mode"), viper.GetString("log_level") )
}

func StartRouter(fn func(*gin.Engine)) {
	r := gin.New()

	fn(r)

	r.Run(viper.GetString("http_addr")) // listen and server on 0.0.0.0:8080
	/*
		srv := &http.Server{
			Addr:           viper.GetString("http_addr"),
			Handler:        r,
			ReadTimeout:    30 * time.Second,
			WriteTimeout:   30 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		http2.ConfigureServer(srv, &http2.Server{})
		err = srv.ListenAndServe()
	*/
}

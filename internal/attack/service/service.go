package service

import (
	"database/sql"
	"fmt"
	"net"
	"runtime/debug"
	"strings"
	"time"

	"github.com/RioRizkyRainey/pokedex/internal/attack/config"
	"github.com/RioRizkyRainey/pokedex/internal/attack/service/attack"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mattn/go-colorable"
	"github.com/snowzach/rotatefilehook"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	log "github.com/sirupsen/logrus"
)

var (
	// DB variable global db
	DB          *sql.DB
	initialized = false
)

func configureLogging() {
	lLevel := config.Get("server.log.level")
	fmt.Println("Setting log level to ", lLevel)
	switch strings.ToUpper(lLevel) {
	default:
		fmt.Println("Unknown level [", lLevel, "]. Log level set to ERROR")
		log.SetLevel(log.ErrorLevel)
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "FATAL":
		log.SetLevel(log.FatalLevel)
	}

	currentTime := time.Now()

	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   "logs/pokemon-" + currentTime.Format("2006-01-02") + ".log",
		MaxSize:    50, // megabytes
		MaxBackups: 3,
		MaxAge:     7, //days
		Level:      log.GetLevel(),
		Formatter: &log.JSONFormatter{
			TimestampFormat: time.RFC822,
		},
	})

	if err != nil {
		log.Fatalf("Failed to initialize file rotate hook: %v", err)
	}

	log.SetLevel(log.GetLevel())
	log.SetOutput(colorable.NewColorableStdout())
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})
	log.AddHook(rotateFileHook)
}

func recoverPanic() {
	if r := recover(); r != nil {
		log.WithField("panic", r).WithField("stack trace", string(debug.Stack())).Error("we panicked!")
	}
}

// Start this server
func Start() {
	defer recoverPanic()
	configureLogging()

	dbHost := config.Get("db.mysql.host")
	dbPort := config.Get("db.mysql.port")
	dbUser := config.Get("db.mysql.user")
	dbPass := config.Get("db.mysql.password")
	dbName := config.Get("db.mysql.name")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4,utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	var err error

	DB, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(config.GetInt("db.mysql.maxopen"))
	DB.SetMaxIdleConns(config.GetInt("db.mysql.maxidle"))

	list, err := net.Listen("tcp", config.Get("server.port"))
	if err != nil {
		fmt.Println("SOMETHING HAPPEN")
	}

	server := grpc.NewServer(grpc.MaxRecvMsgSize(1024*1024*32), grpc.MaxSendMsgSize(1024*1024*32))
	cOpts := []grpc.DialOption{grpc.WithInsecure()}
	fmt.Println("address " + config.Get("pokemon.connect.host") + config.Get("pokemon.connect.port"))
	pokeConn, err := grpc.Dial(config.Get("pokemon.connect.host")+config.Get("pokemon.connect.port"), cOpts...)
	if err != nil {
		panic(err)
	}
	moveConn, err := grpc.Dial(config.Get("move.connect.host")+config.Get("move.connect.port"), cOpts...)
	if err != nil {
		panic(err)
	}

	attack.Server(server, pokeConn, moveConn)

	reflection.Register(server)

	fmt.Println("Server Run at ", config.Get("server.host"), config.Get("server.port"))

	err = server.Serve(list)
	if err != nil {
		fmt.Println("Unexpected Error", err)
	}

}

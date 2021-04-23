package service

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"strings"
	"time"

	"github.com/RioRizkyRainey/pokedex/internal/gateway/config"
	"github.com/RioRizkyRainey/pokedex/internal/gateway/service/attack"
	"github.com/RioRizkyRainey/pokedex/internal/gateway/service/moves"
	"github.com/RioRizkyRainey/pokedex/internal/gateway/service/pokemon"
	"github.com/gorilla/mux"
	"github.com/hyperjumptech/jiffy"
	"github.com/mattn/go-colorable"
	"github.com/snowzach/rotatefilehook"
	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
)

var (
	Router *mux.Router
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
		Filename:   "logs/gateway-" + currentTime.Format("2006-01-02") + ".log",
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

func InitializeRouter() {
	fmt.Println("Initialize Router")

	cOpts := []grpc.DialOption{grpc.WithInsecure()}
	fmt.Println("address " + config.Get("pokemon.connect.host") + config.Get("pokemon.connect.port"))
	pokeConn, err := grpc.Dial(config.Get("pokemon.connect.host")+config.Get("pokemon.connect.port"), cOpts...)
	if err != nil {
		panic(err)
	}

	fmt.Println("address " + config.Get("move.connect.host") + config.Get("move.connect.port"))
	moveConn, err := grpc.Dial(config.Get("move.connect.host")+config.Get("move.connect.port"), cOpts...)
	if err != nil {
		panic(err)
	}

	fmt.Println("address " + config.Get("attack.connect.host") + config.Get("attack.connect.port"))
	attackConn, err := grpc.Dial(config.Get("attack.connect.host")+config.Get("attack.connect.port"), cOpts...)
	if err != nil {
		panic(err)
	}

	pokemon.Server(Router, pokeConn)
	moves.Server(Router, moveConn)
	attack.Server(Router, attackConn)
}

// Start this server
func Start() {
	defer recoverPanic()
	startTime := time.Now()

	Router = mux.NewRouter()
	configureLogging()

	var wait time.Duration

	graceShut, err := jiffy.DurationOf(config.Get("server.timeout.graceshut"))
	if err != nil {
		panic(err)
	}
	wait = graceShut

	WriteTimeout, err := jiffy.DurationOf(config.Get("server.timeout.write"))
	if err != nil {
		panic(err)
	}
	ReadTimeout, err := jiffy.DurationOf(config.Get("server.timeout.read"))
	if err != nil {
		panic(err)
	}
	IdleTimeout, err := jiffy.DurationOf(config.Get("server.timeout.idle"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Server binding to %s", config.Get("server.host")+config.Get("server.port"))

	srv := &http.Server{
		Addr: config.Get("server.host") + config.Get("server.port"),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: WriteTimeout,
		ReadTimeout:  ReadTimeout,
		IdleTimeout:  IdleTimeout,
		Handler:      Router, // Pass our instance of gorilla/mux in.
	}

	InitializeRouter()

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	dur := time.Now().Sub(startTime)
	durDesc := jiffy.DescribeDuration(dur, jiffy.NewWant())
	log.Infof("Shutting down. This Hansip been protecting the world for %s", durDesc)
	os.Exit(0)

}

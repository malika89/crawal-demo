package main

import (
	"crawl/handler"
	"flag"
	"fmt"
	"github.com/judwhite/go-svc"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var (
	flagSet = flag.NewFlagSet("tasks", flag.ExitOnError)
	cfgPath = flagSet.String("config", "./conf/", "Path of Config Files")
	version = flagSet.Bool("version", false, "show relate version info")
)
var (
	Version  string
	CommitId string
	Built    string
)

type program struct {
	LogFile *os.File
	svr     *handler.Server
}

func (p *program) Init(env svc.Environment) error {

	if env.IsWindowsService() {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			return err
		}

		logPath := filepath.Join(dir, "example.log")

		f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}

		p.LogFile = f

		log.SetOutput(f)
	}
	flagSet.Parse(os.Args[1:])

	if *version {
		fmt.Printf("commit id:%s\n", CommitId)
		fmt.Printf("built by %s %s/%s at %s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH, Built)
		os.Exit(2)
	}

	daemon := &handler.Server{
		Config: *cfgPath,
	}
	if err := daemon.Init(); err != nil {
		return err
	}

	p.svr = daemon
	return nil
}

func (p *program) Start() error {
	fmt.Printf("Starting...\n")
	go p.svr.Start()
	return nil
}

func (p *program) Stop() error {
	fmt.Printf("Stopping...\n")
	if err := p.svr.Stop(); err != nil {
		return err
	}
	fmt.Printf("Stopped.\n")
	return nil
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	prg := program{
		svr: &handler.Server{},
	}
	defer func() {
		if prg.LogFile != nil {
			if closeErr := prg.LogFile.Close(); closeErr != nil {
				log.Printf("error closing '%s': %v\n", prg.LogFile.Name(), closeErr)
			}
		}
	}()
	if err := svc.Run(&prg); err != nil {
		log.Println(err)
	}
}

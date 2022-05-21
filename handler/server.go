package handler

import (
	"context"
	"crawl/conf"
	"crawl/handler/models"
	"crawl/scheduler"
	"crawl/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type Server struct {
	server *http.Server
	hasPid bool
	Config string
}

func (s *Server) Init() error {
	// load config
	if err := conf.Init(s.Config); err != nil {
		return err
	}
	// init db
	models.Init()

	if err := s.initHttp(); err != nil {
		return err
	}

	if !s.hasPid {
		dir := filepath.Dir(conf.Conf.Server.Pid)
		if err := os.MkdirAll(dir, 755); err != nil {
			return err
		}

		pid := strconv.Itoa(os.Getpid())
		if err := ioutil.WriteFile(conf.Conf.Server.Pid, []byte(pid), 644); err != nil {
			return err
		}

	}
	go func() {
		log.Println("launch scheduler...")
		scheduler.Schedule()
	}() //启动调度器
	return nil

}

func (s *Server) initHttp() error {
	if utils.CheckLink(conf.Conf.Server.Host, conf.Conf.Server.Port) {
		return fmt.Errorf("address: %s:%d was binded", conf.Conf.Server.Host, conf.Conf.Server.Port)
	}
	gin.SetMode(gin.DebugMode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	//set route
	setupRoute(engine)

	addr := fmt.Sprintf("%s:%d", conf.Conf.Server.Host, conf.Conf.Server.Port)
	s.server = &http.Server{
		Addr:    addr,
		Handler: engine,
	}
	return nil

}

func (s *Server) Stop() error {
	if s.server == nil {
		return nil
	}
	timeout := time.Duration(conf.Conf.Server.ShutdownTimeout) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}
	select {
	case <-ctx.Done():
		fmt.Printf("server shutdown timeout of %ds\n", conf.Conf.Server.ShutdownTimeout)
	}
	log.Println("server exiting")
	if _, err := os.Stat(conf.Conf.Server.Pid); err == nil {
		if err := os.Remove(conf.Conf.Server.Pid); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) Start() {
	if err := s.server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

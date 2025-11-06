package prometheus_test

import (
	"github.com/FANddaaa/test/prometheus/config"
	"github.com/FANddaaa/test/prometheus/monitor"
	"github.com/FANddaaa/test/prometheus/service"
)

type Server struct {
	PromService *service.PromService
}

func NewServer(configPath string) *Server {
	cfg := &config.Config{}
	if err := cfg.Init(configPath); err != nil {
		panic(err)
	}

	return &Server{
		PromService: &service.PromService{
			Port:     cfg.Prometheus.Port,
			Interval: cfg.Prometheus.Interval,
		},
	}
}

func (s *Server) Init() {
	monitor.Init()
	s.PromService.Init()
}

func (s *Server) Start() {
	s.PromService.Start()
}

func (s *Server) Stop() {
	s.PromService.Stop()
}

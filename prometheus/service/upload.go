package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/FANddaaa/test/prometheus/monitor"
)

type PromService struct {
	ctx      context.Context
	cancel   context.CancelFunc
	server   *http.Server
	Port     int
	Interval int
}

func (s *PromService) Init() {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: mux,
	}
}

func (s *PromService) Start() {
	s.ctx, s.cancel = context.WithCancel(context.Background())
	go func() {
		fmt.Printf("数据暴露在 localhost:%d/metrics\n", s.Port)
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("服务器启动失败: %v", err)
		}
	}()
	go s.upload()
}

func (s *PromService) Stop() {
	fmt.Printf("收到关闭信号，开始优雅关闭 localhost:%d...\n", s.Port)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 优雅关闭服务器
	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatalf("服务器关闭失败: %v", err)
	}
	s.cancel()
	fmt.Printf("localhost:%d 已优雅关闭\n", s.Port)
}

func (s *PromService) upload() {
	tick := time.NewTicker(time.Second * time.Duration(s.Interval))
	defer tick.Stop()
	busID := "nlb-test"

	for {
		select {
		case <-s.ctx.Done():
			fmt.Println("收到关闭信号，停止upload循环")
			return
		case <-tick.C:
			InPkts := rand.Int() % 1000
			OutPkts := rand.Int() % 1000
			InBits := rand.Int() % 1000
			OutBits := rand.Int() % 1000
			InBps := InBits / s.Interval
			OutBps := OutBits / s.Interval
			monitor.NlbMessage.WithLabelValues(busID, monitor.TypePkt, monitor.UnitPkt, monitor.DirectionIn).Add(float64(InPkts))
			monitor.NlbMessage.WithLabelValues(busID, monitor.TypePkt, monitor.UnitPkt, monitor.DirectionOut).Add(float64(OutPkts))
			monitor.NlbMessage.WithLabelValues(busID, monitor.TypePkt, monitor.UnitBit, monitor.DirectionIn).Add(float64(InBits))
			monitor.NlbMessage.WithLabelValues(busID, monitor.TypePkt, monitor.UnitBit, monitor.DirectionOut).Add(float64(OutBits))
			monitor.NlbMessage.WithLabelValues(busID, monitor.TypePkt, monitor.UnitBps, monitor.DirectionIn).Set(float64(InBps))
			monitor.NlbMessage.WithLabelValues(busID, monitor.TypePkt, monitor.UnitBps, monitor.DirectionOut).Set(float64(OutBps))
		}
	}
}

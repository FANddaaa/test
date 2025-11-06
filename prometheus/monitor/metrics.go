package monitor

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Prometheus相关
const (
	Namespace = "g3Network"
	SubSystem = "nlb_monitor"
)
const (
	LabelBusID = "busID"
)
const (
	LabelType = "type"
	TypeConn  = "conn"
	TypePkt   = "pkt"
	TypeHc    = "hc"
)
const (
	LabelUnit      = "unit"
	UnitNewConn    = "new_conn"    //
	UnitCurrConn   = "curr_conn"   //
	UnitFailedConn = "failed_conn" //
	UnitBit        = "bit"         // 流量
	UnitBps        = "bps"         // 带宽
	UnitPkt        = "pkt"         // 包量
	UnitPps        = "pps"         // 每秒网络数据包数量
	UnitUnhealthy  = "unhealthy"   // 异常数
	UnitHealthy    = "healthy"     // 正常数
)
const (
	LabelDirection = "direction"
	DirectionIn    = "in"
	DirectionOut   = "out"
)

var (
	NlbMessage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: SubSystem,
			Name:      "nlb_message",
		},
		[]string{LabelBusID, LabelType, LabelUnit, LabelDirection},
	)
)

func Init() {
	prometheus.MustRegister(NlbMessage)
}

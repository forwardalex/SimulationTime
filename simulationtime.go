package SimulationTime

import (
	"context"
	"time"
)

const (
	ns = "ns"
	ms = "ms"
	s  = "s"
)

type SimTime struct {
	Tick     float64 // Tick is the current tick of the simulation
	Time     int64
	Unit     time.Duration
	TimeRate float64 //timeRate is the rate of the simulation time
	TimeBus  *Bus
}

func (s *SimTime) GetNow() int64 {
	return s.Time
}

func (s *SimTime) SetNow(t int64) {
	s.Time = t
}

func (s *SimTime) WatchTime(ctx context.Context) TimeChannel {
	return s.TimeBus.Subscribe(ctx)
}

func (s *SimTime) PublishTime(timestamp int64) {
	s.TimeBus.Publish(timestamp)
}

func (s *SimTime) UpdateTimeRate(rate float64) {
	s.TimeRate = rate
}

func NewSimTime(opts ...OptionFunc) *SimTime {
	opt := defaultOption.clone()
	opt.init()
	for _, o := range opts {
		o(opt)
	}
	return &opt.SimTime
}

func (s *SimTime) Star(ctx context.Context) {
	tick := time.Tick(time.Duration(s.Tick) * s.Unit)
	for {
		select {
		case <-tick:
			s.Time += int64(s.Tick * s.TimeRate)
			s.PublishTime(s.Time)
		case <-ctx.Done():
			return
		}
	}
}

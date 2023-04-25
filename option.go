package SimulationTime

import (
	"time"
)

type OptionFunc func(*Option)

type Option struct {
	SimTime SimTime
}

var defaultOption = Option{
	SimTime: SimTime{
		Tick:     100,
		Unit:     time.Millisecond,
		TimeRate: 1,
		Time:     time.Now().UnixMilli(),
		TimeBus:  NewTimeBus(),
	},
}

func (opt *Option) clone() *Option {
	clone := *opt
	return &clone
}

func (opt *Option) init() {

}

func WithUnit(unit time.Duration) OptionFunc {
	return func(o *Option) {
		o.SimTime.Unit = unit
	}
}

func WithStarTime(star int64) OptionFunc {
	return func(option *Option) {
		option.SimTime.Time = star
	}
}

func WithTick(tick float64) OptionFunc {
	return func(option *Option) {
		option.SimTime.Tick = tick
	}
}

func WithTimeRate(rate float64) OptionFunc {
	return func(option *Option) {
		option.SimTime.TimeRate = rate
	}
}

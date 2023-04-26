package main

import (
	"context"
	"fmt"
	"github.com/forwardalex/SimulationTime"
	"time"
)

func main() {
	star := time.Now().UnixMilli()
	sim := SimulationTime.NewSimTime(
		SimulationTime.WithTick(100),
		SimulationTime.WithStarTime(star),
		SimulationTime.WithTimeRate(10),
		SimulationTime.WithUnit(time.Millisecond))
	go sim.Star(context.Background())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := sim.WatchTime(ctx)
	go func() {
		for {
			v, ok := <-ch
			if ok {
				fmt.Println(v)
				fmt.Println(time.Now().UnixMilli())
			}
		}
	}()
	sim.UpdateTimeRate(1)
}

Simulation是一款为用户实现仿真的应用程序，可以帮助用户构建仿真系统的时间环境，以便更好地模拟和研究各种情况下的行为和效果。

该应用程序的主要特点是具有可配置的仿真速率，这使得用户可以调整模拟世界的时间流速，以便更好地模拟不同场景下的情况。同时，Simulation还提供了简单易用的用户界面，用户可以轻松地进行时间流速的修改和控制。

Simulation的设计初衷是提供一种灵活性和易用性较高的仿真方案，不要求高精度测量和精确仿真，而是注重灵活性和可定制性，以满足用户的不同需求。

总之，Simulation是一款非常有用的仿真应用程序，用户可以通过它来构建各种仿真系统，并对不同场景进行模拟和研究，从而更好地理解系统行为和效应，并做出更好的决策。

example

```go
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
```

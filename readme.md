[中文介绍](./readmezh.md)

Additional information and introduction:

Simulation is an application designed to simulate various scenarios for users, helping them to construct time environments in simulation systems for better analysis and understanding of behavior and effects.

The main feature of the program is its configurable simulation rate, which allows users to adjust the time flow speed of the simulated world to better simulate different scenarios

The design intention of Simulation is to provide a flexible and easy-to-use simulation solution, without requiring high-precision measurements and accurate simulations, but rather focusing on flexibility and customizability to meet users' varying needs.

In summary, Simulation is a very useful simulation application enabling users to construct various simulation systems, simulate and analyze different scenarios, and better understand system behavior and effects to make more informed decisions.

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
    // sim.GetNow() getsim time not by chan
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

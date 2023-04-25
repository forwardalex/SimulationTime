package SimulationTime

import (
	"context"
	"sync"
)

type TimeChannel chan int64
type TimeChannelSlice []TimeChannel
type Bus struct {
	Subscribers    []TimeChannel
	SubscribersIds []int
	rm             sync.RWMutex
	pool           *Set //chan pool chan池
	use            *Set
	count          int
}

func (b *Bus) Subscribe(ctx context.Context) TimeChannel {
	ch := make(TimeChannel, 1024)
	var id int
	b.rm.Lock()
	if b.pool.IsEmpty() {
		b.count += 1
		id = b.count

	} else {
		id = b.pool.Pop().(int)
	}
	b.Subscribers = append(b.Subscribers, ch)
	b.SubscribersIds = append(b.SubscribersIds, id)
	b.use.Add(id)
	go b.watchDone(ctx, id)
	b.rm.Unlock()
	return ch
}

func (b *Bus) Publish(timestamp int64) {
	b.rm.RLock()
	go func() {
		for _, ch := range b.Subscribers {
			ch <- timestamp
		}
	}()
	b.rm.RUnlock()
}
func (b *Bus) watchDone(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			b.remove(id)
			return
		}
	}
}

func (b *Bus) remove(id int) {
	b.rm.Lock()
	defer b.rm.Unlock()
	for i, v := range b.SubscribersIds {
		if v == id {
			b.SubscribersIds = append(b.SubscribersIds[:i], b.SubscribersIds[i+1:]...)
			break
		}
	}
	b.pool.Add(id)
	b.use.Remove(id)
}
func NewTimeBus() *Bus {
	return &Bus{
		pool: NewSet(),
		use:  NewSet(),
	}
}

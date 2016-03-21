package traffic

import (
	"math/rand"
	"sync"
	"time"
)

// Generator accepts a bunch of Patterns and starts generating traffic.
type Generator struct {
	concurrency uint32
	inputChan   chan bool
	lock        *sync.Mutex
	maxItems    uint32
	patterns    []Pattern
}

// NewGenerator returns a new Generator instance.
func NewGenerator(iterations int, concurrency int) *Generator {
	return &Generator{
		concurrency: uint32(concurrency),
		inputChan:   make(chan bool),
		lock:        &sync.Mutex{},
		maxItems:    uint32(iterations),
		patterns:    []Pattern{},
	}
}

// AddPattern adds a traffic pattern to the list of patterns contained
// in this generator.
func (g *Generator) AddPattern(p *Pattern) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.patterns = append(g.patterns, *p)
}

// Execute fast forwards through all the iterations.
func (g *Generator) Execute() {
	g.lock.Lock()
	defer g.lock.Unlock()

	// Check if any patterns have been added to the Generator
	if len(g.patterns) == 0 {
		return
	}

	// pSum is the sum of the scores of all patterns
	pSum := 0
	for _, p := range g.patterns {
		pSum += p.Probability
	}
	var wg = sync.WaitGroup{}
	for i := uint32(0); i < g.concurrency; i++ {
		wg.Add(1)
		go g.execute(&wg, pSum)
	}

	for i := uint32(0); i < g.maxItems; i++ {
		g.inputChan <- true
	}

	wg.Wait()
}

func (g *Generator) execute(wg *sync.WaitGroup, pSum int) {
	randGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	defer wg.Done()
	for {
		select {
		case <-g.inputChan:
			// Generate a random number
			prob := randGenerator.Intn(pSum) + 1
			c := 0
			for _, p := range g.patterns {
				c += p.Probability
				if c >= prob {
					p.Fn()
					break
				}
			}
		case <-time.After(100 * time.Millisecond):
			return
		}
	}
}

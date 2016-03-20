package traffic

import (
	"math/rand"
	"sync"
	"time"
)

// Generator accepts a bunch of Patterns and starts generating traffic.
type Generator struct {
	itemsDelivered int
	lock           *sync.Mutex
	maxItems       int
	patterns       []Pattern
	randGenerator  *rand.Rand
	sum            int
}

// NewGenerator returns a new Generator instance.
func NewGenerator(iterations int) *Generator {
	return &Generator{
		itemsDelivered: 0,
		lock:           &sync.Mutex{},
		maxItems:       iterations,
		patterns:       []Pattern{},
		randGenerator:  rand.New(rand.NewSource(time.Now().UnixNano())),
		sum:            0,
	}
}

// AddPattern adds a traffic pattern to the list of patterns contained
// in this generator.
func (g *Generator) AddPattern(p *Pattern) {
	g.lock.Lock()
	g.patterns = append(g.patterns, *p)
	g.lock.Unlock()
	g.recalculate()
}

// recalculate runs a bunch of calculations to set the correct probabilities
// in the generator.
func (g *Generator) recalculate() {
	g.lock.Lock()
	defer g.lock.Unlock()
	newSum := 0
	for _, p := range g.patterns {
		newSum += p.Probability
	}
	g.sum = newSum
}

// Execute fast forwards through all the iterations.
func (g *Generator) Execute() {
	g.lock.Lock()
	defer g.lock.Unlock()

	// Check if any patterns have been added to the Generator
	if len(g.patterns) == 0 {
		//errors.New("No Patterns have been added to the generator.")
		return
	}

	for {
		// Check if we've already delivered all the items
		if g.itemsDelivered >= g.maxItems {
			return
		}

		// Generate a random number
		g.itemsDelivered++
		prob := g.randGenerator.Intn(g.sum) + 1
		c := 0
		for _, p := range g.patterns {
			c += p.Probability
			if c >= prob {
				p.Fn()
				break
			}
		}
	}
}

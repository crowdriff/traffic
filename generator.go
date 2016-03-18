package traffic

import (
	"math/rand"
	"time"
)

// Generator accepts a bunch of Patterns and starts generating traffic.
type Generator struct {
	patterns      []Pattern
	randGenerator *rand.Rand
	sum           int
}

// NewGenerator returns a new Generator instance.
func NewGenerator() *Generator {
	return &Generator{
		[]Pattern{},
		rand.New(rand.NewSource(time.Now().UnixNano())),
		0,
	}
}

// AddPattern adds a traffic pattern to the list of patterns contained
// in this generator.
func (g *Generator) AddPattern(p *Pattern) {
	g.patterns = append(g.patterns, *p)
	g.recalculate()
}

// recalculate runs a bunch of calculations to set the correct probabilities
// in the generator.
func (g *Generator) recalculate() {
	// TODO: Implement recalculation
}

// Next calls a random function according to the traffic patterns contained
// in this generator.
func (g *Generator) Next() interface{} {
	// TODO: Implement this
	return struct{}{}
}

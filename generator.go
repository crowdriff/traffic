package traffic

import (
	"errors"
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
	newSum := 0
	for _, p := range g.patterns {
		newSum += p.Probability
	}
	g.sum = newSum
}

// Next calls a random function according to the traffic patterns contained
// in this generator.
func (g *Generator) Next() (interface{}, error) {
	// Check if any patterns have been added to the Generator
	if len(g.patterns) == 0 {
		return nil, errors.New("No Patterns have been added to the generator.")
	}

	// Generate a random number
	prob := g.randGenerator.Intn(g.sum) + 1
	c := 0
	for _, p := range g.patterns {
		c += p.Probability
		if c >= prob {
			return p.Fn()
		}
	}
	return nil, errors.New("This should never be reached")
}

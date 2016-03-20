package traffic

// Pattern holds a function that gets called with the given probability.
type Pattern struct {
	Probability int
	Fn          func()
}

// NewPattern generates a new Pattern with given parameters.
func NewPattern(probability int, fn func()) *Pattern {
	return &Pattern{probability, fn}
}

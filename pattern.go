package traffic

// Pattern holds a function that gets called with the given probability. The
// function should conform to the Producer interface.
type Pattern struct {
	Probability int
	Fn          Producer
}

// NewPattern generates a new Pattern with given parameters.
func NewPattern(probability int, fn Producer) (*Pattern, error) {
	return &Pattern{probability, fn}
}

func (p *Pattern) Generate() interface{} {

}

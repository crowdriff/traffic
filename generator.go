package traffic

// Generator accepts a bunch of Patterns and starts generating traffic.
type Generator interface {
	AddPattern(p *Pattern)
	Next()
}

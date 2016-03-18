package traffic

// Producer is an interface implemented by all functions that form a traffic
// pattern.
type Producer interface {
	Generate() interface{}
}

// DefaultProducer is the default implementation of the Producer interface.
type DefaultProducer struct{}

// Empty is an empty struct

// Generate allows the DefaultProducer to comply with the Producer interface.
func (p DefaultProducer) Generate() interface{} {
	return struct{}{}
}

// NewDefaultProducer produces a new instance of the DefaultProducer.
func NewDefaultProducer() Producer {
	return DefaultProducer{}
}

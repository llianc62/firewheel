package firewheel

// Option is the way allowed caller offers values they needs.
type Option interface {
	// Apply it-self recursively.
	Apply(Option) Option
}

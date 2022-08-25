package retry

import (
	"errors"
	"time"
)

// Policy determines how Run retries the function.
type Policy interface {
	Next() (time.Duration, bool)
}

type exponentialPolicy struct {
	attempts int
	sleep    time.Duration
}

// ExponentialPolicy retires with an exponential growth in sleep.
func ExponentialPolicy(attempts int, sleep time.Duration) Policy {
	return &exponentialPolicy{
		attempts: attempts,
		sleep:    sleep,
	}
}

func (p *exponentialPolicy) Next() (time.Duration, bool) {
	p.attempts--
	if p.attempts <= 0 {
		return 0, false
	}

	defer func() {
		p.sleep *= 2
	}()

	return p.sleep, true
}

// Run executes the function while the Policy allows
// until it returns nil or Stop.
func Run(p Policy, fn func() error) error {
	if p == nil {
		return errors.New("policy must not be nil")
	}

	if err := fn(); err != nil {
		if s, ok := err.(stop); ok {
			return s.error
		}

		if sleep, ok := p.Next(); ok {
			time.Sleep(sleep)
			Run(p, fn)
		}

		return err
	}

	return nil
}

type stop struct {
	error
}

// Stop wraps an error and stops retrying.
func Stop(err error) error {
	return stop{err}
}

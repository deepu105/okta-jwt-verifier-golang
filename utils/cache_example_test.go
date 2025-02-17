package utils_test

import (
	"fmt"

	jwtverifier "github.com/okta/okta-jwt-verifier-golang"
	"github.com/okta/okta-jwt-verifier-golang/utils"
)

// ForeverCache caches values forever
type ForeverCache struct {
	values map[string]interface{}
	lookup func(string) (interface{}, error)
}

// Get returns the value for the given key
func (c *ForeverCache) Get(key string) (interface{}, error) {
	value, ok := c.values[key]
	if ok {
		return value, nil
	}
	value, err := c.lookup(key)
	if err != nil {
		return nil, err
	}
	c.values[key] = value
	return value, nil
}

// ForeverCache implements the read-only Cacher interface
var _ utils.Cacher = (*ForeverCache)(nil)

// NewForeverCache takes a lookup function and returns a cache
func NewForeverCache(lookup func(string) (interface{}, error)) (utils.Cacher, error) {
	return &ForeverCache{
		values: map[string]interface{}{},
		lookup: lookup,
	}, nil
}

// Example demonstrating how the JwtVerifier can be configured with a custom Cache function.
func Example() {
	jwtVerifierSetup := jwtverifier.JwtVerifier{
		Cache: NewForeverCache,
		// other fields here
	}

	verifier := jwtVerifierSetup.New()
	fmt.Println(verifier)
}

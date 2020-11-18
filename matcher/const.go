package matcher

import (
	"github.com/segmentio/fasthash/fnv1a"
)

// Pattern TODO
type Pattern string

// ItemID TODO
func (p Pattern) ItemID() uint64 {
	return fnv1a.HashString64(string(p))
}

// MatcherType type
type MatcherType int

// MatcherType enum
const (
	_             MatcherType = iota
	RgMatcher                 // regex matcher
	SimpleMatcher             // simple matcher
)

// IterFunc TODO
type IterFunc func(Pattern, interface{}) bool

type matcherHandler interface {
	Match(string) (Pattern, interface{})
	Load(Pattern, interface{}) (Pattern, interface{}, error)
	Get(Pattern) (Pattern, interface{})
	Delete(Pattern) (Pattern, interface{})
	Iter(IterFunc)
}

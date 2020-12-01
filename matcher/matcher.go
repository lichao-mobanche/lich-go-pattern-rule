package matcher

import (
	"sync"
)

// Matcher TODO
type Matcher struct {
	innerMatcher matcherHandler
	*sync.RWMutex
}

// New TODO
func New(mt MatcherType) *Matcher {
	var mh matcherHandler
	switch mt {
	case RgMatcher:
		mh = newRgMatcher()
	case SimpleMatcher:
		mh = newSpMatcher()
	default:
		return nil
	}
	return &Matcher{mh, &sync.RWMutex{}}
}

// Match TODO
func (m Matcher) Match(rawPath string) (string, interface{}) {
	m.RLock()
	defer m.RUnlock()
	p, r := m.innerMatcher.Match(rawPath)
	return string(p), r
}

// Load TODO
func (m *Matcher) Load(pattern string, rule interface{}) (string, interface{}, error) {
	m.Lock()
	defer m.Unlock()
	p, r, e := m.innerMatcher.Load(Pattern(pattern), rule)
	return string(p), r, e
}

// Get TODO
func (m Matcher) Get(pattern string) (string, interface{}) {
	m.RLock()
	defer m.RUnlock()
	p, r := m.innerMatcher.Get(Pattern(pattern))
	return string(p), r
}

//Delete TODO
func (m *Matcher) Delete(pattern string) (string, interface{}) {
	m.Lock()
	defer m.Unlock()
	p, r := m.innerMatcher.Delete(Pattern(pattern))
	return string(p), r
}

// Iter TODO
func (m Matcher) Iter(f IterFunc) {
	m.RLock()
	defer m.RUnlock()
	m.innerMatcher.Iter(f)
}

// Size TODO
func (m Matcher) Size() int {
	m.RLock()
	defer m.RUnlock()
	return m.innerMatcher.Size()
}

// CheckPattern TODO
func (m Matcher) CheckPattern(pattern string)(string, error) {
	m.RLock()
	defer m.RUnlock()
	return m.innerMatcher.CheckPattern(pattern)
}
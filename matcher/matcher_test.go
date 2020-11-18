package matcher_test

import (
	"testing"
	"github.com/lichao-mobanche/lich-go-pattern-rule/matcher"
	"github.com/stretchr/testify/assert"
)

type PatternAndRule struct {
	k string
	v int
}

func createRgMatcher(data []PatternAndRule) *matcher.Matcher {
	matcher := matcher.New(matcher.RgMatcher)
	for _, d := range data {
		matcher.Load(d.k, d.v)
	}
	return matcher
}

func createSpMatcher(data []PatternAndRule) *matcher.Matcher {
	matcher := matcher.New(matcher.SimpleMatcher)
	for _, d := range data {
		matcher.Load(d.k, d.v)
	}
	return matcher
}

func TestRgMatcher(t *testing.T) {
	matcher := createRgMatcher(
		[]PatternAndRule{
			{"/.*", 123},
			{"/abc/.*", 456},
			
		},
	)

	tests := []struct {
		path      string
		expected int
	}{
		{"/abc/def", 456},
		
	}

	for _, test := range tests {
		_, v := matcher.Match(test.path)
		assert.Equal(t, test.expected, v)
	}
}

func TestSpMatcher(t *testing.T) {
	matcher := createSpMatcher(
		[]PatternAndRule{
			{"/a*", 123},
			{"/abc/def*.gj$", 456},
			
		},
	)

	tests := []struct {
		path      string
		expected int
	}{
		{"/abc/def.gj", 456},
		
	}

	for _, test := range tests {
		_, v := matcher.Match(test.path)
		assert.Equal(t, test.expected, v)
	}
}

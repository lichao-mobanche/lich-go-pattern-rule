package matcher

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/lichao-mobanche/lich-go-pattern-rule/pkg/sortmap"
	"github.com/segmentio/fasthash/fnv1a"
)

type matchUnit struct {
	rule interface{}
	rg   *regexp.Regexp
}

func (mu matchUnit) ItemID() uint64 {
	return fnv1a.HashString64(mu.rg.String())
}
func newRgMatcher() *rgMatcher {
	comparer := func(a interface{}, b interface{}) bool {
		aprefix, _ := a.(*matchUnit).rg.LiteralPrefix()
		bprefix, _ := b.(*matchUnit).rg.LiteralPrefix()
		return len(aprefix) > len(bprefix)
	}
	return &rgMatcher{unit: sortmap.New(comparer)}
}

type rgMatcher struct {
	unit *sortmap.Map
}

func (m rgMatcher) Match(path string) (p Pattern, r interface{}) {
	iterator := sortmap.NewBaseIter(m.unit)
	iterator.Iter(
		func(item sortmap.Item) bool {
			it := item.(*matchUnit)
			if it.rg.FindString(path) == path {
				p = Pattern(it.rg.String())
				r = it.rule
				return false
			}
			return true
		},
	)
	return
}

func (m *rgMatcher) Load(pattern Pattern, rule interface{}) (Pattern, interface{}, error) {
	rg, e := regexp.Compile(string(pattern))
	prefix, _ := rg.LiteralPrefix()
	if e != nil || !strings.HasPrefix(prefix, "/") {
		return "", nil, fmt.Errorf("illegal pattern %s", pattern)

	}
	iid := pattern.ItemID()
	if item := m.unit.Get(iid); item != nil {
		oldRule := item.(*matchUnit).rule
		item.(*matchUnit).rule = rule
		return pattern, oldRule, nil
	}

	mu := &matchUnit{rule, rg}
	m.unit.Add(mu)
	return pattern, rule, nil
}

func (m rgMatcher) Get(pattern Pattern) (p Pattern, r interface{}) {
	iid := pattern.ItemID()
	if item := m.unit.Get(iid); item != nil {
		p = pattern
		r = item.(*matchUnit).rule
	}
	return p, r
}

func (m *rgMatcher) Delete(pattern Pattern) (p Pattern, r interface{}) {
	iid := pattern.ItemID()
	if item := m.unit.Get(iid); item != nil {
		p = pattern
		r = item.(*matchUnit).rule
		m.unit.Delete(iid)
	}
	return p, r
}

func (m rgMatcher) Iter(f IterFunc) {
	iterator := sortmap.NewBaseIter(m.unit)
	iterator.Iter(
		func(item sortmap.Item) bool {
			it := item.(*matchUnit)
			return f(Pattern(it.rg.String()), it.rule)
		},
	)
}

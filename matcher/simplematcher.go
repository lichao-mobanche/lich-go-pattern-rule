package matcher

import (
	"fmt"
	"strings"
	"github.com/lichao-mobanche/lich-go-pattern-rule/pkg/sortmap"
	"github.com/segmentio/fasthash/fnv1a"
)

type simpleMatchUnit struct {
	rule interface{}
	pattern   Pattern
}
func (s simpleMatchUnit) ItemID() uint64 {
	return fnv1a.HashString64(string(s.pattern))
}
func newSpMatcher() *spMatcher {
	comparer:=func(a interface{}, b interface{}) bool {
		return len(a.(*simpleMatchUnit).pattern) > len(b.(*simpleMatchUnit).pattern)
	}
	return &spMatcher{sortmap.New(comparer)}
}

type spMatcher struct {
	unit *sortmap.Map
}

func (s spMatcher) match(pattern, path string) bool {
	
	pathlen := len(path)
	pos:=make([]int,pathlen+1)
	pos[0]=0
	numpos:=1
	for index, pat:=range pattern {
		if pat=='$' && index+1 == len(pattern) {
			return pos[numpos-1] == pathlen
		}
		if pat == '*' {
			numpos=pathlen - pos[0] +1
			for i:=1;i<numpos;i++ {
				pos[i]=pos[i-1]+1
			}
		} else {
			newnumpos:=0
			for i:=0;i<numpos;i++ {
				if pos[i]<pathlen && path[pos[i]]== byte(pat) {
					pos[newnumpos]=pos[i]+1
					newnumpos++
				}
			}
			numpos = newnumpos;
			if numpos==0 {
				return false
			}
		}
	}
	return true
}

func (s *spMatcher) Match(path string) (p Pattern, r interface{}) {
	iterator := sortmap.NewBaseIter(s.unit)
	iterator.Iter(
		func(item sortmap.Item) bool {
			it := item.(*simpleMatchUnit)
			if s.match(string(it.pattern), path) {
				p = Pattern(it.pattern)
				r = it.rule
				return false
			}
			return true
		},
	)
	return
}

func (s *spMatcher) Load(pattern Pattern, rule interface{}) (Pattern, interface{}, error) {
	if !strings.HasPrefix(string(pattern), "/"){
		return "", nil, fmt.Errorf("illegal pattern %s", pattern)
	}
	iid := pattern.ItemID()
	if item := s.unit.Get(iid); item != nil {
		oldRule := item.(*matchUnit).rule
		item.(*matchUnit).rule = rule
		return pattern, oldRule, nil
	}

	sm := &simpleMatchUnit{rule, pattern}
	s.unit.Add(sm)
	return pattern, rule, nil
}

func (s spMatcher) Get(pattern Pattern) (p Pattern, r interface{}) {
	iid := pattern.ItemID()
	if item := s.unit.Get(iid); item != nil {
		p = pattern
		r = item.(*simpleMatchUnit).rule
	}
	return p, r
}

func (s *spMatcher) Delete(pattern Pattern) (p Pattern, r interface{}) {
	iid := pattern.ItemID()
	if item := s.unit.Get(iid); item != nil {
		p = pattern
		r = item.(*simpleMatchUnit).rule
		s.unit.Delete(iid)
	}
	return p, r
}

func (s spMatcher) Iter(f IterFunc) {
	iterator := sortmap.NewBaseIter(s.unit)
	iterator.Iter(
		func(item sortmap.Item) bool {
			it := item.(*simpleMatchUnit)
			return f(Pattern(it.pattern), it.rule)
		},
	)
}

func (s spMatcher) Size() int {
	return s.unit.Size()
}

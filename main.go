package main

import (
	"fmt"
	"github.com/lichao-mobanche/lich-go-pattern-rule/matcher"
)

func main() {
	rematcher := matcher.New(matcher.RgMatcher)
	fmt.Println(rematcher.Load("/.*", 123))
	fmt.Println(rematcher.Load("/abc/.*", 456))
	fmt.Println(rematcher.Match("/abc/def"))

	spmatcher := matcher.New(matcher.SimpleMatcher)
	fmt.Println(spmatcher.Load("/a*", 123))
	fmt.Println(spmatcher.Load("/abc/def*.gj$", 456))
	fmt.Println(spmatcher.Match("/abc/def.gj"))
}

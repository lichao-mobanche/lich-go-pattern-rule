# lich-go-pattern-rule

A common library for pattern rule use case.

## Install

You can get the library with ``go get``

```
go get -u github.com/lichao-mobanche/lich-go-pattern-rule
```

## Usage
matcher.RgMatcher is regexp pattern, comply with standard regular expressions. And pattern should start white '/'.

matcher.SimpleMatcher is simple pattern, Which only '*' and '&' are effective。And pattern should start white '/'.
'*': designates 0 or more instances of any valid character.
'$': designates the end of the URL.
```
package main

import (
	"fmt"
	"github.com/lichao-mobanche/lich-go-pattern-rule/matcher"
)

func main() {
	matcher := matcher.New(matcher.RgMatcher)
	fmt.Println(matcher.Load("/.*", 123))
	fmt.Println(matcher.Load("/abc/.*", 456))
	fmt.Println(matcher.Match("/abc/def"))

	spmatcher := matcher.New(matcher.SimpleMatcher)
	fmt.Println(spmatcher.Load("/a*", 123))
	fmt.Println(spmatcher.Load("/abc/def*.gj$", 456))
	fmt.Println(spmatcher.Match("/abc/def.gj"))
}

```
## License
  MIT licensed.
// Strings manipulation

package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {

	p("Contains:", s.Contains("test", "es"))
	p("Count:", s.Count("test", "t"))
	p("Prefix", s.HasPrefix("test", "te"))
	p("Suffix:", s.HasSuffix("test", "st"))
	p("Index:", s.Index("test", "e"))
	p("Join:", s.Join([]string{"a", "b", "c"}, "-"))
	p("Repeat:", s.Repeat("a", 5))
	p("Replace:", s.Replace("foo", "o", "0", 1))
	p("Replace:", s.Replace("foo", "o", "0", 2))
	p("Split:", s.Split("a,b,c", ","))
	p("ToLower:", s.ToLower("ABC"))
	p("ToUpper:", s.ToUpper("abc"))

}

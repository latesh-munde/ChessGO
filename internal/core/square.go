package core

import "fmt"

type Square struct {
	Rank int
	File int
}

func (s Square) Offset(dr, df int) (Square, bool) {
	r := s.Rank + dr
	f := s.File + df
	if r < 0 || r > 7 || f < 0 || f > 7 {
		return Square{}, false
	}
	return Square{Rank: r, File: f}, true
}

func (s Square) String() string {
	return fmt.Sprintf("%c%d", 'a'+s.File, s.Rank+1)
}

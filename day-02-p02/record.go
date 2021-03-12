package main

type record struct {
	Min int
	Max int
	Letter rune
	Password []rune
}

func (r *record) isValid() bool{
	if r.Password[r.Min-1] == r.Letter && r.Password[r.Max-1] == r.Letter {
		return false
	}

	if r.Password[r.Min-1] == r.Letter || r.Password[r.Max-1] == r.Letter {
		return true
	}
	return false
}

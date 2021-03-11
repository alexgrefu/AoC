package main

type record struct {
	Min int
	Max int
	Letter rune
	Password map[rune]int
}

func (r *record) isValid() bool{
	if n, ok := r.Password[r.Letter]; ok {
		if r.Min <= n && n <= r.Max {
			return true
		}
	}
	return false
}

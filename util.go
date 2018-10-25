package main

func score(s stats, c class) float64 {
	dump := s[c.Dump]
	if dump < s[c.Dump2] {
		dump = s[c.Dump2]
	}
	return (s[c.Primary]*1.5 + s[c.Secondary] - dump) / 42.0
}

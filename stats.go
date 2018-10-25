package main

type stat string
type stats map[stat]float64

// Attributes
const (
	Str  stat = "STR"
	Dex  stat = "DEX"
	Con  stat = "CON"
	Int  stat = "INT"
	Wis  stat = "WIS"
	Cha  stat = "CHA"
	None stat = ""
)

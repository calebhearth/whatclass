package main

type class struct {
	Name      string  `json:"name"`
	Primary   stat    `json:"primary"`
	Secondary stat    `json:"secondary"`
	Dump      stat    `json:"dump"`
	Dump2     stat    `json:"alternate_dump"`
	Score     float64 `json:"score"`
}

type klasses []class

func (s klasses) Len() int {
	return len(s)
}

func (s klasses) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s klasses) Less(i, j int) bool {
	return s[i].Score < s[j].Score
}

var (
	classes klasses = []class{
		class{"Barbarian", Str, Con, Int, None, 0},
		class{"Bard", Cha, Dex, Int, None, 0},
		class{"Cleric", Wis, Con, Int, Cha, 0},
		class{"Druid", Wis, Con, Str, None, 0},
		class{"STR Fighter", Str, Con, Int, Cha, 0},
		class{"DEX Fighter", Dex, Con, Int, Cha, 0},
		class{"Monk", Dex, Wis, Int, None, 0},
		class{"Paladin", Str, Cha, Int, None, 0},
		class{"Ranger", Dex, Con, Cha, None, 0},
		class{"Rogue", Dex, Con, Str, None, 0},
		class{"Sorcerer", Cha, Con, Int, None, 0},
		class{"Warlock", Cha, Con, Int, None, 0},
		class{"Wizard", Int, Dex, Str, None, 0},
	}
)

package main

import (
	"fmt"
	"math"
)

const (
	Boss   = "boss"
	Player = "player"
)

func main() {
	boss := &Combatant{
		HitPoints: 104,
		Damage:    8,
		Armor:     1,
	}

	_ = boss

	weapons := []Item{
		{"Dagger", 8, 4, 0},
		{"Shortsword", 10, 5, 0},
		{"Warhammer", 25, 6, 0},
		{"Longsword", 40, 7, 0},
		{"Greataxe", 74, 8, 0},
	}

	armors := []Item{
		{"Leather", 13, 0, 1},
		{"Chainmail", 31, 0, 2},
		{"Splintmail", 53, 0, 3},
		{"Bandedmail", 75, 0, 4},
		{"Platemail", 102, 0, 5},
		{"None", 0, 0, 0},
	}

	rings := []Item{
		{"Damage +1", 25, 1, 0},
		{"Damage +2", 50, 2, 0},
		{"Damage +3", 100, 3, 0},
		{"Defense +1", 20, 0, 1},
		{"Defense +2", 40, 0, 2},
		{"Defense +3", 80, 0, 3},
		{"None", 0, 0, 0},
		{"None", 0, 0, 0},
	}

	loadouts := [][]Item{}
	for _, weapon := range weapons {
		for _, armor := range armors {
			for i, ring1 := range rings {
				for _, ring2 := range rings[i+1:] {
					// This will have duplicates, but who cares, there's only like 1000 combinations
					loadouts = append(loadouts, []Item{weapon, armor, ring1, ring2})
				}
			}
		}
	}

	mostExpensive := -1
	for _, loadout := range loadouts {
		totalForLoadout := sumLoadout(loadout)
		cost := totalForLoadout.Cost
		player := &Combatant{
			HitPoints: 100,
			Damage:    totalForLoadout.Damage,
			Armor:     totalForLoadout.Armor,
		}

		if getWinner(player, boss) == Boss {
			if cost > mostExpensive {
				mostExpensive = cost
			}
		}
	}

	fmt.Println(mostExpensive)
}

type Item struct {
	Name   string
	Cost   int
	Damage int
	Armor  int
}

func (i Item) String() string {
	return fmt.Sprintf("<%v>", i.Name)
}

type Combatant struct {
	HitPoints int
	Damage    int
	Armor     int
}

func (c Combatant) String() string {
	return fmt.Sprintf("HP: <%d> Damage: <%d>, Armor: <%d>", c.HitPoints, c.Damage, c.Armor)
}

func getWinner(player, boss *Combatant) string {
	bossDamagePerTurn := getDamage(boss.Damage, player.Armor)
	playerDamagePerTurn := getDamage(player.Damage, boss.Armor)

	bossTurnsToWins := ceilDiv(player.HitPoints, bossDamagePerTurn)
	playerTurnsToWins := ceilDiv(boss.HitPoints, playerDamagePerTurn)

	if bossTurnsToWins < playerTurnsToWins {
		return Boss
	}
	return Player
}

func getDamage(damage, armor int) int {
	if armor >= damage {
		return 1
	}
	return damage - armor
}

func ceilDiv(a, b int) int {
	return int(math.Ceil(float64(a) / float64(b)))
}

func sumLoadout(items []Item) Item {
	total := Item{}
	for _, item := range items {
		total.Cost += item.Cost
		total.Damage += item.Damage
		total.Armor += item.Armor
	}
	return total
}

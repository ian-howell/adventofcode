package main

import (
	"fmt"
	"strings"
)

// bossDamage is part of the input for my specific puzzle
const bossDamage = 9

type SpellType int

const (
	None SpellType = iota
	MagicMissile
	Drain
	Shield
	Poison
	Recharge
)

func (s SpellType) String() string {
	switch s {
	case MagicMissile:
		return "MagicMissile"
	case Drain:
		return "Drain"
	case Shield:
		return "Shield"
	case Poison:
		return "Poison"
	case Recharge:
		return "Recharge"
	case None:
		fallthrough
	default:
		return "None"
	}
}

type Mana int

const (
	MagicMissileCost Mana = 53
	DrainCost             = 73
	ShieldCost            = 113
	PoisonCost            = 173
	RechargeCost          = 229
)

type TurnType int

const (
	PlayerTurn TurnType = 0
	BossTurn            = 1
)

func (t TurnType) String() string {
	if t == PlayerTurn {
		return "Player"
	}
	return "Boss"
}

type State struct {
	Turn TurnType

	LastSpell  SpellType
	PlayerMana Mana

	PlayerHP int
	BossHP   int

	ShieldCounter   int
	PoisonCounter   int
	RechargeCounter int
}

func main() {
	initialState := State{
		Turn:       PlayerTurn,
		PlayerMana: 500,
		PlayerHP:   50,
		BossHP:     58,
	}
	path := Dijkstra(initialState)
	for _, step := range path {
		fmt.Println(step)
	}

	fmt.Println(getScore(path))
}

func Weight(s State) int {
	switch s.LastSpell {
	case MagicMissile:
		return int(MagicMissileCost)
	case Drain:
		return int(DrainCost)
	case Shield:
		return int(ShieldCost)
	case Poison:
		return int(PoisonCost)
	case Recharge:
		return int(RechargeCost)
	}
	return 0
}

func Neighbors(u State) []State {
	neighbors := []State{}

	applyEffects(&u)

	// This can happen if the boss is poisoned
	if u.BossHP <= 0 {
		u.Turn = PlayerTurn
		u.LastSpell = None
		return []State{u}
	}

	if u.Turn == BossTurn {
		playerNewHP := u.PlayerHP - bossDamage
		if u.ShieldCounter > 0 {
			playerNewHP += 7
		}

		return []State{
			{
				Turn:            PlayerTurn,
				PlayerMana:      u.PlayerMana,
				PlayerHP:        playerNewHP,
				BossHP:          u.BossHP,
				ShieldCounter:   u.ShieldCounter,
				PoisonCounter:   u.PoisonCounter,
				RechargeCounter: u.RechargeCounter,
			},
		}
	}

	if u.PlayerHP == 1 {
		u.PlayerHP = 0
		return []State{u}
	}

	if u.PlayerMana >= MagicMissileCost {
		neighbors = append(
			neighbors,
			State{
				Turn:            BossTurn,
				PlayerMana:      u.PlayerMana - MagicMissileCost,
				PlayerHP:        u.PlayerHP - 1,
				BossHP:          u.BossHP - 4,
				ShieldCounter:   u.ShieldCounter,
				PoisonCounter:   u.PoisonCounter,
				RechargeCounter: u.RechargeCounter,
				LastSpell:       MagicMissile,
			})
	}
	if u.PlayerMana > DrainCost {
		neighbors = append(
			neighbors,
			State{
				Turn:            BossTurn,
				PlayerMana:      u.PlayerMana - DrainCost,
				PlayerHP:        u.PlayerHP + 2 - 1,
				BossHP:          u.BossHP - 2,
				ShieldCounter:   u.ShieldCounter,
				PoisonCounter:   u.PoisonCounter,
				RechargeCounter: u.RechargeCounter,
				LastSpell:       Drain,
			})
	}
	if u.PlayerMana > ShieldCost && u.ShieldCounter == 0 {
		neighbors = append(
			neighbors,
			State{
				Turn:            BossTurn,
				PlayerMana:      u.PlayerMana - ShieldCost,
				PlayerHP:        u.PlayerHP - 1,
				BossHP:          u.BossHP,
				ShieldCounter:   6,
				PoisonCounter:   u.PoisonCounter,
				RechargeCounter: u.RechargeCounter,
				LastSpell:       Shield,
			})
	}
	if u.PlayerMana > PoisonCost && u.PoisonCounter == 0 {
		neighbors = append(
			neighbors,
			State{
				Turn:            BossTurn,
				PlayerMana:      u.PlayerMana - PoisonCost,
				PlayerHP:        u.PlayerHP - 1,
				BossHP:          u.BossHP,
				ShieldCounter:   u.ShieldCounter,
				PoisonCounter:   6,
				RechargeCounter: u.RechargeCounter,
				LastSpell:       Poison,
			})
	}
	if u.PlayerMana > RechargeCost && u.RechargeCounter == 0 {
		neighbors = append(
			neighbors,
			State{
				Turn:            BossTurn,
				PlayerMana:      u.PlayerMana - RechargeCost,
				PlayerHP:        u.PlayerHP - 1,
				BossHP:          u.BossHP,
				ShieldCounter:   u.ShieldCounter,
				PoisonCounter:   u.PoisonCounter,
				RechargeCounter: 5,
				LastSpell:       Recharge,
			})
	}

	return neighbors
}

func applyEffects(u *State) {
	if u.ShieldCounter > 0 {
		u.ShieldCounter--
	}
	if u.PoisonCounter > 0 {
		u.BossHP -= 3
		u.PoisonCounter--
	}
	if u.RechargeCounter > 0 {
		u.PlayerMana += 101
		u.RechargeCounter--
	}
}

func (s State) String() string {
	sb := strings.Builder{}
	if s.Turn == PlayerTurn {
		sb.WriteString("-- Player turn --\n")
	} else {
		sb.WriteString("-- Boss turn --\n")
	}
	armor := 0
	if s.ShieldCounter > 0 {
		armor = 7
	}
	sb.WriteString(fmt.Sprintf("- Player has %d hit points, %d armor, %d mana\n", s.PlayerHP, armor, s.PlayerMana))
	sb.WriteString(fmt.Sprintf("- Boss has %d hit points\n", s.BossHP))
	return sb.String()
}

func getScore(path []State) int {
	score := 0
	for _, step := range path {
		score += Weight(step)
	}
	return score
}

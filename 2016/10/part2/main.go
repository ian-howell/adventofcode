package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RuleType int

const (
	BotRule = iota
	OutputRule
)

func (rt RuleType) String() string {
	return map[RuleType]string{BotRule: "bot", OutputRule: "output"}[rt]
}

type Rule struct {
	Type RuleType
	I    int
}

func (r Rule) String() string {
	return fmt.Sprintf("%v %v", r.Type, r.I)
}

func main() {
	bots, rules := GetData()
	outputs := Do(bots, rules)
	fmt.Println(outputs[0][0] * outputs[1][0] * outputs[2][0])
}

func Do(bots map[int][]int, rules map[int][2]Rule) map[int][]int {
	q := []int{}
	for bot, values := range bots {
		if len(values) == 2 {
			q = append(q, bot)
		}
	}

	outputs := map[int][]int{}
	for len(q) > 0 {
		u := q[len(q)-1]
		q = q[:len(q)-1]

		rule := rules[u]
		bot := bots[u]

		lowRule, highRule := rule[0], rule[1]
		lowVal, highVal := bot[0], bot[1]
		if lowVal > highVal {
			lowVal, highVal = highVal, lowVal
		}

		if lowRule.Type == OutputRule {
			outputs[lowRule.I] = append(outputs[lowRule.I], lowVal)
		} else {
			bots[lowRule.I] = append(bots[lowRule.I], lowVal)
			if len(bots[lowRule.I]) == 2 {
				q = append(q, lowRule.I)
			}
		}

		if highRule.Type == OutputRule {
			outputs[highRule.I] = append(outputs[highRule.I], highVal)
		} else {
			bots[highRule.I] = append(bots[highRule.I], highVal)
			if len(bots[highRule.I]) == 2 {
				q = append(q, highRule.I)
			}
		}

		// bot u no longer has any chips
		bots[u] = nil

	}

	return outputs
}

func GetData() (bots map[int][]int, rules map[int][2]Rule) {
	bots = map[int][]int{}
	rules = map[int][2]Rule{}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if strings.HasPrefix(s.Text(), "value") {
			parts := strings.Split(s.Text(), " ")
			botNo, value := atoi(parts[5]), atoi(parts[1])
			bots[botNo] = append(bots[botNo], value)
		} else {
			parts := strings.Split(s.Text(), " ")
			givingBot := atoi(parts[1])
			var lowRule Rule
			var highRule Rule

			if parts[5] == "output" {
				lowRule = Rule{OutputRule, atoi(parts[6])}
			} else {
				lowRule = Rule{BotRule, atoi(parts[6])}
			}

			if parts[10] == "output" {
				highRule = Rule{OutputRule, atoi(parts[11])}
			} else {
				highRule = Rule{BotRule, atoi(parts[11])}
			}

			rules[givingBot] = [2]Rule{lowRule, highRule}
		}
	}
	return bots, rules
}

func atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return i
}

func PrintData(bots map[int][]int, rules map[int][2]Rule) {
	for bot, values := range bots {
		fmt.Printf("Bot % 4d has: ", bot)
		for value := range values {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
	for bot, rule := range rules {
		fmt.Printf("Bot %d gives low to %v and high to %v\n", bot, rule[0], rule[1])
	}
}

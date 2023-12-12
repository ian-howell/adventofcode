package main

type CamelCard string

func (c CamelCard) HandType() HandType {
	freq := map[byte]int{}
	for i := 0; i < len(c); i++ {
		freq[c[i]]++
	}
	switch len(freq) {

	case 5:
		return HighCardHandType
	case 4:
		return OnePairHandType
	case 3:
		for _, val := range freq {
			switch val {
			case 2:
				return TwoPairHandType
			case 3:
				return ThreeOfAKindHandType
			}
		}
	case 2:
		for _, val := range freq {
			switch val {
			case 1, 4:
				return FourOfAKindHandType
			case 2, 3:
				return FullHouseHandType
			}
		}
	case 1:
		return FiveOfAKindHandType
	}
	return NoneHandType
}

func (c CamelCard) Less(other CamelCard) bool {
	iType, jType := c.HandType(), other.HandType()
	if iType < jType {
		return true
	} else if iType > jType {
		return false
	}

	value := func(b byte) int {
		switch b {
		case 'T':
			return 10
		case 'J':
			return 11
		case 'Q':
			return 12
		case 'K':
			return 13
		case 'A':
			return 14
		}
		return int(b - '0')
	}

	for k := 0; k < 5; k++ {
		if value(c[k]) < value(other[k]) {
			return true
		} else if value(c[k]) > value(other[k]) {
			return false
		}
	}
	return false
}

type CamelCardSlice []CamelCard

func (cs CamelCardSlice) Len() int      { return len(cs) }
func (cs CamelCardSlice) Swap(i, j int) { cs[i], cs[j] = cs[j], cs[i] }
func (cs CamelCardSlice) Less(i, j int) bool { return cs[i].Less(cs[j]) }

type HandType int

const (
	NoneHandType = iota

	HighCardHandType
	OnePairHandType
	TwoPairHandType
	ThreeOfAKindHandType
	FullHouseHandType
	FourOfAKindHandType
	FiveOfAKindHandType
)

func (h HandType) String() string {
	switch h {
	case HighCardHandType:
		return "HighCardHandType"
	case OnePairHandType:
		return "OnePairHandType"
	case TwoPairHandType:
		return "TwoPairHandType"
	case ThreeOfAKindHandType:
		return "ThreeOfAKindHandType"
	case FullHouseHandType:
		return "FullHouseHandType"
	case FourOfAKindHandType:
		return "FourOfAKindHandType"
	case FiveOfAKindHandType:
		return "FiveOfAKindHandType"
	default:
		return "NoneHandType"
	}
}

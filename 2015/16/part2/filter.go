package main

func Filter(as *AuntSue) bool {
	filterFns := []FilterFn{
		Children,
		Cats,
		Samoyeds,
		Pomeranians,
		Akitas,
		Vizslas,
		Goldfish,
		Trees,
		Cars,
		Perfumes,
	}

	for _, fn := range filterFns {
		if !fn(as) {
			return false
		}
	}
	return true
}

type FilterFn func(*AuntSue) bool

func Children(as *AuntSue) bool {
	return as.Children == -1 || as.Children == 3
}

func Cats(as *AuntSue) bool {
	return as.Cats == -1 || as.Cats > 7
}

func Samoyeds(as *AuntSue) bool {
	return as.Samoyeds == -1 || as.Samoyeds == 2
}

func Pomeranians(as *AuntSue) bool {
	return as.Pomeranians == -1 || as.Pomeranians < 3
}

func Akitas(as *AuntSue) bool {
	return as.Akitas == -1 || as.Akitas == 0
}

func Vizslas(as *AuntSue) bool {
	return as.Vizslas == -1 || as.Vizslas == 0
}

func Goldfish(as *AuntSue) bool {
	return as.Goldfish == -1 || as.Goldfish < 5
}

func Trees(as *AuntSue) bool {
	return as.Trees == -1 || as.Trees > 3
}

func Cars(as *AuntSue) bool {
	return as.Cars == -1 || as.Cars == 2
}

func Perfumes(as *AuntSue) bool {
	return as.Perfumes == -1 || as.Perfumes == 1
}

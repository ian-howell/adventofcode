package main

import (
	"strconv"
	"strings"
)

type OperationType int

const (
	NoneOp OperationType = iota
	HlfOp
	TplOp
	IncOp
	JmpOp
	JieOp
	JioOp
)

func (o OperationType) String() string {
	switch o {
	case HlfOp:
		return "Hlf"
	case TplOp:
		return "Tpl"
	case IncOp:
		return "Inc"
	case JmpOp:
		return "Jmp"
	case JieOp:
		return "Jie"
	case JioOp:
		return "Jio"
	}
	return "None"
}

type Instruction struct {
	Op     OperationType
	Reg    string
	Offset int
}

func newHlf(parts []string) Instruction {
	return Instruction{Op: HlfOp, Reg: strings.ToUpper(parts[1])}
}

func newTpl(parts []string) Instruction {
	return Instruction{Op: TplOp, Reg: strings.ToUpper(parts[1])}
}

func newInc(parts []string) Instruction {
	return Instruction{Op: IncOp, Reg: strings.ToUpper(parts[1])}
}

func newJmp(parts []string) Instruction {
	// Assuming the input has no errors...
	offset, _ := strconv.Atoi(parts[1])
	return Instruction{Op: JmpOp, Offset: offset}
}

func newJie(parts []string) Instruction {
	// Assuming the input has no errors...
	offset, _ := strconv.Atoi(parts[2])
	return Instruction{Op: JieOp, Reg: strings.ToUpper(parts[1]), Offset: offset}
}

func newJio(parts []string) Instruction {
	// Assuming the input has no errors...
	offset, _ := strconv.Atoi(parts[2])
	return Instruction{Op: JioOp, Reg: strings.ToUpper(parts[1]), Offset: offset}
}

func normalize(s string) string {
	return strings.ReplaceAll(s, ",", "")
}

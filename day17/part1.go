package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

type Emulator struct {
	registerA int
	registerB int
	registerC int
	out string
}

func (emulator *Emulator) comboToNum(operand int) int {
	if operand <= 3 {
		return operand
	}

	result := 0
	
	switch operand {
	case 4:
		result = emulator.registerA
	case 5:
		result = emulator.registerB
	case 6:
		result = emulator.registerC
	}

	return result
}

func (emulator *Emulator) output(output int) {
	emulator.out += strconv.Itoa(output) + ","
}

func (emulator *Emulator) runOperation(opcode int, operand int) (int, error) {
	comboValue := emulator.comboToNum(operand)
	jumpedTo := -1

	switch opcode {
	// adv
	case 0:
		denom := math.Pow(2, float64(comboValue))
		emulator.registerA = int(math.Floor(float64(emulator.registerA) / denom))
	// bxl
	case 1:
		emulator.registerB ^= operand
	// bst
	case 2:
		emulator.registerB = comboValue % 8
	// jnz
	case 3:
		if emulator.registerA != 0 { 
			jumpedTo = operand
		}
	// bxc
	case 4:
		emulator.registerB ^= emulator.registerC
	// out
	case 5:
		emulator.output(comboValue % 8)
	// bdv
	case 6:
		denom := math.Pow(2, float64(comboValue))
		emulator.registerB = int(math.Floor(float64(emulator.registerA) / denom))
	// cdv
	case 7:
		denom := math.Pow(2, float64(comboValue))
		emulator.registerC = int(math.Floor(float64(emulator.registerA) / denom))
	default:
		return jumpedTo, fmt.Errorf("Found unexpected opcode: %d", opcode)
	}

	return jumpedTo, nil
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func charToInt(char byte) int {
	return int(char) - int('0')
}

func parseRegisters(fileContents []byte) (int, int, int, error) {
	registerA := 0
	registerB := 0
	registerC := 0

	registerCount := 0

	i := 0
	contentsLen := len(fileContents)
	for i < contentsLen && registerCount < 3 {
		char := fileContents[i]

		if !isDigit(char) {
			i++
			continue
		}	

		strNum := string(char)
		i++
		for i < contentsLen && isDigit(fileContents[i]) {
			strNum += string(fileContents[i])
			i++
		}
		
		num, err := strconv.Atoi(strNum)

		if err != nil {
			return 0, 0, 0, err
		}
		
		if registerCount == 0 {
			registerA = num
		} else if registerCount == 1 {
			registerB = num
		} else if registerCount == 2 {
			registerC = num
		}

		registerCount++
	}

	if registerCount < 3 {
		return 0, 0, 0, fmt.Errorf("Missing one or more registers")
	}

	return registerA, registerB, registerC, nil
}

func parseProgram(fileContents []byte) []int {
	// hardcoded to skip first 3 lines because im lazy
	i := 0
	newLines := 0
	for newLines < 3 {
		if fileContents[i] == '\n' {
			newLines++
		}
		i++
	}

	program := make([]int, 0)
	fileLen := len(fileContents)
	for i < fileLen {
		char := fileContents[i]
		if isDigit(char) { 
			program = append(program, charToInt(char))
		}
		i++
	}

	return program
}

func emulate(fileName string) (string, error) {
	fileContent, err := os.ReadFile(fileName)	
	
	if err != nil {
		return "", err
	}

	emulator := new(Emulator)

	emulator.registerA, emulator.registerB, emulator.registerC, err = parseRegisters(fileContent)

	if err != nil {
		return "", err
	}
	
	program := parseProgram(fileContent)
	
	i := 0
	for i < len(program) {
		opcode := program[i]
		operand := program[i + 1]
		jumpedTo, err := emulator.runOperation(opcode, operand)

		if err != nil {
			return "", err
		}

		if jumpedTo > -1 {
			i = jumpedTo
			continue
		}
		
		i += 2
	}

	return emulator.out, nil
}

func main() {
	result, err := emulate("input.txt")
	
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}


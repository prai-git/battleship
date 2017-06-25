package battleship

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Trim(input string) (output string) {
	output = strings.Replace(input, "\n", "", -1)
	output = strings.Replace(output, "\t", "", -1)
	output = strings.Replace(output, "\r", "", -1)
	output = strings.Replace(output, " ", "", -1)
	return output
}

func GetNumberFromInputAs0to9(message string) (result int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	number, err := reader.ReadString('\n')
	if err != nil || len(Trim(number)) == 0 {
		fmt.Print("Your input is not correct! Please type again.")
		result = GetNumberFromInputAs0to9(message)
		return
	} else {
		var err2 error
		result, err2 = strconv.Atoi(Trim(number))
		if err2 != nil || result < 1 || result > 9 {
			fmt.Print("Your input is not correct! Please type again. \n")
			result = GetNumberFromInputAs0to9(message)
			return
		}
	}

	return
}

func GetNumberFromInputAsAtoZ(message string) (result int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	str, err := reader.ReadString('\n')
	if err != nil || len(Trim(str)) == 0 {
		fmt.Print("Your input is not correct! Please type again. \n")
		result = GetNumberFromInputAsAtoZ(message)
		return
	}

	result = int([]rune(Trim(str))[0])
	A := int([]rune("A")[0])
	Z := int([]rune("Z")[0])

	if result < A || result > Z {
		fmt.Print("Your input is not correct! Please type again. \n")
		result = GetNumberFromInputAsAtoZ(message)
		return
	}

	result = int([]rune(Trim(str))[0]) - int([]rune("A")[0]) + 1

	return
}

func GetYCordFromInputAsAtoZ(message string) (result int, input string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	str, err := reader.ReadString('\n')
	if err != nil || len(Trim(str)) == 0 {
		fmt.Print("Your input is not correct! Please type again. \n")
		result, input = GetYCordFromInputAsAtoZ(message)

		return
	}

	result = int([]rune(Trim(str))[0])
	A := int([]rune("A")[0])
	Z := int([]rune("Z")[0])

	if result < A || result > Z {
		fmt.Print("Your input is not correct! Please type again. \n")
		result, input = GetYCordFromInputAsAtoZ(message)

		return
	}

	result = int([]rune(Trim(str))[0]) - int([]rune("A")[0]) + 1
	input = string(Trim(str)[0])
	return
}

func GetXCordFromInputAs0to9(message string) (result int, input string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	number, err := reader.ReadString('\n')
	if err != nil || len(Trim(number)) == 0 {
		fmt.Print("Your input is not correct! Please type again.")
		result, input = GetXCordFromInputAs0to9(message)
		return
	} else {

		var err2 error
		result, err2 = strconv.Atoi(Trim(number))
		input = number
		if err2 != nil || result < 1 || result > 9 {
			fmt.Print("Your input is not correct! Please type again. \n")
			result, input = GetXCordFromInputAs0to9(message)
			return
		}
	}

	return
}

func GetShipTypeFromInput(message string) (result ShipType) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	str, err := reader.ReadString('\n')

	if err != nil || len(Trim(str)) == 0 {
		fmt.Print("Your input is not correct! Please type again. \n")
		result = GetShipTypeFromInput(message)
		return
	}

	output := int([]rune(Trim(str))[0])
	PChar := int([]rune("P")[0])
	QChar := int([]rune("Q")[0])

	if output < PChar || output > QChar {
		fmt.Print("Your input is not correct! Please type again. \n")
		result = GetShipTypeFromInput(message)
		return
	}

	result = P
	if output != PChar {
		result = Q
	}

	return
}

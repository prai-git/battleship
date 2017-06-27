package battleship

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
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

func GetNumberFrom0to9(number string) (result int, err error) {

	result, err = strconv.Atoi(Trim(number))
	if err != nil || result < 1 || result > 9 {
		err = errors.New("Your input is not correct! Please type again. \n")
		return
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

func GetNumberFromAtoZ(str string) (result int, err error) {

	if len(Trim(str)) == 0 {
		err = errors.New("Your input is not correct! Please type again. \n")
		return
	}

	result = int([]rune(Trim(str))[0])
	A := int([]rune("A")[0])
	Z := int([]rune("Z")[0])

	if result < A || result > Z {
		err = errors.New("Your input is not correct! Please type again. \n")
		return
	}

	result = int([]rune(Trim(str))[0]) - int([]rune("A")[0]) + 1

	return
}

func convertNumberToChar(number int) (result int) {

	result = int([]rune("A")[0]) + number - 1

	return
}

func GetYCordFromAtoZ(str string) (result int, input string, err error) {
	if len(Trim(str)) == 0 {
		err = errors.New("Your input is not correct! Please type again. \n")
		return
	}

	result = int([]rune(Trim(str))[0])
	A := int([]rune("A")[0])
	Z := int([]rune("Z")[0])

	if result < A || result > Z {
		err = errors.New("Your input is not correct! Please type again. \n")
		return
	}

	result = int([]rune(Trim(str))[0]) - int([]rune("A")[0]) + 1
	input = string(Trim(str)[0])
	return
}

func GetXCordFrom0to9(number string) (result int, input string, err error) {

	if err != nil || len(Trim(number)) == 0 {
		err = errors.New("Your input is not correct! Please type again. \n")
		return
	} else {

		var err2 error
		result, err2 = strconv.Atoi(Trim(number))
		input = number
		if err2 != nil || result < 1 || result > 9 {
			err = errors.New("Your input is not correct! Please type again. \n")
			return
		}
	}

	return
}

func GetShipType(str string) (result ShipType, err error) {

	if len(Trim(str)) == 0 {
		err = errors.New("Your input is not correct! Please type again. \n")
		return
	}

	output := int([]rune(Trim(str))[0])
	PChar := int([]rune("P")[0])
	QChar := int([]rune("Q")[0])

	if output < PChar || output > QChar {
		err = errors.New("Your input is not correct! Please type again. \n")
		return
	}

	result = P
	if output != PChar {
		result = Q
	}

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

// Read a whole file into the memory and store it as array of lines
func ReadLines(path string) (lines []string, err error) {
	var (
		file   *os.File
		part   []byte
		prefix bool
	)
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}

func WriteLines(lines []string, path string) (err error) {
	var (
		file *os.File
	)

	if file, err = os.Create(path); err != nil {
		return
	}
	defer file.Close()

	//writer := bufio.NewWriter(file)
	for _, item := range lines {
		//fmt.Println(item)
		_, err := file.WriteString(strings.TrimSpace(item) + "\n")
		//file.Write([]byte(item));
		if err != nil {
			//fmt.Println("debug")
			fmt.Println(err)
			break
		}
	}
	/*content := strings.Join(lines, "\n")
	  _, err = writer.WriteString(content)*/
	return
}

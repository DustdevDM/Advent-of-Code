package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func main() {
	file := "input.txt"
	input := parseInput(file)

	color.Cyan("INPUT:")
	for _, element1 := range input {
		for _, element2 := range element1 {
			fmt.Printf("%c", element2)
		}
		fmt.Printf("\n")
	}
	print("\n")

	positiveCoordinates := generatePositiveCoordinates(file, input)

	color.Cyan("POSITIVE COORDINATES:")
	for _, element1 := range positiveCoordinates {
		for _, element2 := range element1 {
			if element2 == 2 {
				fmt.Printf("🟢")
			}
			if element2 == 1 {
				fmt.Printf("🔵")
			}
			if element2 == 0 {
				fmt.Printf("🔴")
			}
		}
		fmt.Printf("\n")
	}

	result := getResult(input, positiveCoordinates)
	color.Cyan("RESULT:")
	fmt.Println(result)

	result2 := generateResult2(input)
	color.Cyan("RESULT 2:")
	fmt.Println(result2)
}

type intermediateHolder struct {
	stack                 []rune
	hadPositiveCoordinate bool
}

func getResult(input [][]rune, positiveCoordinates [][]int) int {
	returnValue := 0

	for x := range input {
		var hold intermediateHolder
		for y := range input[x] {
			if runeIsADigit(input[x][y]) {
				//is a number
				hold.stack = append(hold.stack, input[x][y])
				if positiveCoordinates[x][y] == 2 {
					hold.hadPositiveCoordinate = true
				}
			} else {
				//is not a number
				if len(hold.stack) != 0 {
					if hold.hadPositiveCoordinate {
						//number endet and is part of the result

						stackToInt, err := strconv.Atoi(string(hold.stack))

						if err != nil {
							log.Panic(err)
						}

						returnValue += stackToInt
						hold.stack = make([]rune, 0)
						hold.hadPositiveCoordinate = false
					} else {
						//number endet but is not part of result
						hold.stack = make([]rune, 0)
						hold.hadPositiveCoordinate = false
					}
				}
			}
		}
	}

	return returnValue
}

func generateResult2(input [][]rune) int {

	returnValue := 0

	for x, colum := range input {
		for y := range colum {
			if input[x][y] == '*' {
				numbers := make([][]rune, 0)
				if x-1 < 0 == false {
					if runeIsADigit(input[x-1][y]) {
						numbers = appendNumberIfNotAlreadyAppended(numbers, getNumberRunes(input, x-1, y))
					}
				}
				if x+1 > len(input) == false {
					if runeIsADigit(input[x+1][y]) {
						numbers = appendNumberIfNotAlreadyAppended(numbers, getNumberRunes(input, x+1, y))
					}
				}
				if y-1 < 0 == false {
					if runeIsADigit(input[x][y-1]) {
						numbers = appendNumberIfNotAlreadyAppended(numbers, getNumberRunes(input, x, y-1))
					}
				}
				if y+1 > len(input[x]) == false {
					if runeIsADigit(input[x][y+1]) {
						numbers = appendNumberIfNotAlreadyAppended(numbers, getNumberRunes(input, x, y+1))
					}
				}
				if x+1 > len(input) == false && y+1 > len(input[x]) == false {
					if runeIsADigit(input[x+1][y+1]) {
						numbers = appendNumberIfNotAlreadyAppended(numbers, getNumberRunes(input, x+1, y+1))
					}
				}
				if x-1 < 0 == false && y-1 < 0 == false {
					if runeIsADigit(input[x-1][y-1]) {
						numbers = appendNumberIfNotAlreadyAppended(numbers, getNumberRunes(input, x-1, y-1))
					}
				}
				if x+1 > len(input) == false && y-1 < 0 == false {
					if runeIsADigit(input[x+1][y-1]) {
						numbers = appendNumberIfNotAlreadyAppended(numbers, getNumberRunes(input, x+1, y-1))
					}
				}
				if x-1 < 0 == false && y+1 > len(input[x]) == false {
					if runeIsADigit(input[x-1][y+1]) {
						numbers = appendNumberIfNotAlreadyAppended(numbers, getNumberRunes(input, x-1, y+1))
					}
				}

				if len(numbers) == 2 {
					runeArrayToInt1, err1 := strconv.Atoi(string(numbers[0]))
					runeArrayToInt2, err2 := strconv.Atoi(string(numbers[1]))

					if err1 != nil {
						log.Panic(err1)
					}

					if err2 != nil {
						log.Panic(err2)
					}

					returnValue += (runeArrayToInt1 * runeArrayToInt2)
				}
			}

		}
	}

	return returnValue
}

func appendNumberIfNotAlreadyAppended(input [][]rune, appendValue []rune) [][]rune {
	shouldNotAppend := false

	for x := range input {
		if reflect.DeepEqual(input[x], appendValue) {
			shouldNotAppend = true
		}
	}

	if shouldNotAppend == false {
		input = append(input, appendValue)
	}
	return input

}

func runeIsADigit(value rune) bool {
	return value == '0' || value == '1' || value == '2' || value == '3' || value == '4' || value == '5' || value == '6' || value == '7' || value == '8' || value == '9'
}

func getNumberRunes(input [][]rune, x int, y int) []rune {
	return collectNumberRunes(make([]rune, 0), input, x, findFirstDigitCoordinate(input, x, y))
}

func collectNumberRunes(returnValue []rune, input [][]rune, x int, y int) []rune {
	if runeIsADigit(input[x][y]) {
		returnValue = append(returnValue, input[x][y])
		return collectNumberRunes(returnValue, input, x, y+1)
	} else {
		return returnValue
	}
}

func findFirstDigitCoordinate(input [][]rune, x int, y int) int {
	if runeIsADigit(input[x][y]) == false {
		return y + 1
	} else if y-1 == -1 {
		return 0
	} else {
		return findFirstDigitCoordinate(input, x, y-1)
	}
}

func parseInput(filepath string) [][]rune {
	var rows, colums int

	input := ReadFile(filepath)

	rows = strings.Count(string(input), "\n") + 1           // Plus one because the last line does not have \n
	colums = len(strings.Split(string(input), "\n")[0]) - 1 //Minus one because the [0] line has a linebreak

	// create two-dimensional array
	returnValue := make([][]rune, rows)
	for i := range returnValue {
		returnValue[i] = make([]rune, colums)
	}

	//set value of two-dimensional array
	for i := 0; i < rows; i++ {
		returnValue[i] = []rune(strings.Split(input, "\n")[i])
	}

	return returnValue
}

func ReadFile(fileinput string) string {
	fileContent, err := os.ReadFile(fileinput)

	if err != nil {
		log.Panic(err)
	}

	return string(fileContent)
}

func generatePositiveCoordinates(filepath string, input [][]rune) [][]int {
	inputtext := ReadFile(filepath)

	var rows, colums int
	rows = strings.Count(string(inputtext), "\n") + 1           // Plus one because the last line does not have \n
	colums = len(strings.Split(string(inputtext), "\n")[0]) - 1 //Minus one because the [0] line has a linebreak

	// create two-dimensional array
	returnValue := make([][]int, rows)
	for i := range returnValue {
		returnValue[i] = make([]int, colums)
	}

	for x, colum := range input {
		for y := range colum {
			if input[x][y] != '.' && input[x][y] != '\n' && input[x][y] != '\r' && (input[x][y] < '0' || input[x][y] > '9') {
				returnValue[x][y] = 1
				if x-1 < 0 == false {
					returnValue[x-1][y] = 2
				}
				if x+1 > len(input) == false {
					returnValue[x+1][y] = 2
				}
				if y-1 < 0 == false {
					returnValue[x][y-1] = 2
				}
				if y+1 > len(input[x]) == false {
					returnValue[x][y+1] = 2
				}
				if x+1 > len(input) == false && y+1 > len(input[x]) == false {
					returnValue[x+1][y+1] = 2
				}
				if x-1 < 0 == false && y-1 < 0 == false {
					returnValue[x-1][y-1] = 2
				}
				if x+1 > len(input) == false && y-1 < 0 == false {
					returnValue[x+1][y-1] = 2
				}
				if x-1 < 0 == false && y+1 > len(input[x]) == false {
					returnValue[x-1][y+1] = 2
				}
			}

		}
	}

	return returnValue
}

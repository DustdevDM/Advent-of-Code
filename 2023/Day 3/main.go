package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	file := "input.txt"
	input := Parse(file)

	positiveCoordinates := InputToNumericPartMap(file, input)

	fmt.Println("POSITIVE COORDINATES:")
	for _, line := range positiveCoordinates {
		for _, partMapping := range line {
			if partMapping == NextToPart {
				fmt.Printf("🟢")
			}
			if partMapping == Part {
				fmt.Printf("🔵")
			}
			if partMapping == NotPart {
				fmt.Printf("🔴")
			}
		}
		fmt.Printf("\n")
	}

	result := GetResultTask1(input, positiveCoordinates)
	fmt.Printf("\nDay 3 Task 1:%s", strconv.Itoa(result))

	result2 := GetResultTask2(input)
	fmt.Printf("\nDay 3 Task 2:%s", strconv.Itoa(result2))

}

// Struct used for stacking and determine number with positive part-type values
type IntermediateHolder struct {
	Stack               []rune
	HadPositivePartType bool
}

// Returns Advent of Code Day 3 result for the first task
func GetResultTask1(input [][]rune, partTypeMapping [][]int) int {
	returnValue := 0

	for x := range input {
		var hold IntermediateHolder
		for y := range input[x] {
			if IsRuneDigit(input[x][y]) {

				hold.Stack = append(hold.Stack, input[x][y])

				if partTypeMapping[x][y] == NextToPart {
					hold.HadPositivePartType = true
				}

			} else {

				if len(hold.Stack) != 0 {

					if hold.HadPositivePartType {
						//number ended and is part of the result

						stackToInt, err := strconv.Atoi(string(hold.Stack))

						if err != nil {
							log.Panic(err)
						}

						returnValue += stackToInt
						hold.Stack = make([]rune, 0)
						hold.HadPositivePartType = false

					} else {
						//number ended but is not part of result
						hold.Stack = make([]rune, 0)
						hold.HadPositivePartType = false
					}
				}
			}
		}
	}

	return returnValue
}

// Returns Advent of Code Day 3 result for the second task
func GetResultTask2(input [][]rune) int {

	returnValue := 0

	for x, colum := range input {
		for y := range colum {
			if input[x][y] == '*' {
				numbers := make([][]rune, 0)
				if x-1 < 0 == false {
					if IsRuneDigit(input[x-1][y]) {
						numbers = AppendUniqueRunes(numbers, DetermineAndGetFullNumber(input, x-1, y))
					}
				}
				if x+1 > len(input) == false {
					if IsRuneDigit(input[x+1][y]) {
						numbers = AppendUniqueRunes(numbers, DetermineAndGetFullNumber(input, x+1, y))
					}
				}
				if y-1 < 0 == false {
					if IsRuneDigit(input[x][y-1]) {
						numbers = AppendUniqueRunes(numbers, DetermineAndGetFullNumber(input, x, y-1))
					}
				}
				if y+1 > len(input[x]) == false {
					if IsRuneDigit(input[x][y+1]) {
						numbers = AppendUniqueRunes(numbers, DetermineAndGetFullNumber(input, x, y+1))
					}
				}
				if x+1 > len(input) == false && y+1 > len(input[x]) == false {
					if IsRuneDigit(input[x+1][y+1]) {
						numbers = AppendUniqueRunes(numbers, DetermineAndGetFullNumber(input, x+1, y+1))
					}
				}
				if x-1 < 0 == false && y-1 < 0 == false {
					if IsRuneDigit(input[x-1][y-1]) {
						numbers = AppendUniqueRunes(numbers, DetermineAndGetFullNumber(input, x-1, y-1))
					}
				}
				if x+1 > len(input) == false && y-1 < 0 == false {
					if IsRuneDigit(input[x+1][y-1]) {
						numbers = AppendUniqueRunes(numbers, DetermineAndGetFullNumber(input, x+1, y-1))
					}
				}
				if x-1 < 0 == false && y+1 > len(input[x]) == false {
					if IsRuneDigit(input[x-1][y+1]) {
						numbers = AppendUniqueRunes(numbers, DetermineAndGetFullNumber(input, x-1, y+1))
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

// Appends a rune-field into a rune-matrix if its not already contained
func AppendUniqueRunes(input [][]rune, appendValue []rune) [][]rune {
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

// Determines if a rune is a ascii-number
func IsRuneDigit(value rune) bool {
	return value == '0' || value == '1' || value == '2' || value == '3' || value == '4' || value == '5' || value == '6' || value == '7' || value == '8' || value == '9'
}

// Determines a full Number by its matrix-coordinates of any digit
func DetermineAndGetFullNumber(input [][]rune, x int, y int) []rune {
	return CollectFullNumberRunes(make([]rune, 0), input, x, DetermineFirstDigitYCoordinate(input, x, y))
}

// Recursion that collects a full number from that starting coordinates of a digit
func CollectFullNumberRunes(returnValue []rune, input [][]rune, x int, y int) []rune {
	if IsRuneDigit(input[x][y]) {
		returnValue = append(returnValue, input[x][y])
		return CollectFullNumberRunes(returnValue, input, x, y+1)
	} else {
		return returnValue
	}
}

// Recursion that converts any matrix-coordinates into the first digit-coordinate of the Number
func DetermineFirstDigitYCoordinate(input [][]rune, x int, y int) int {
	if IsRuneDigit(input[x][y]) == false {
		return y + 1
	} else if y-1 == -1 {
		return 0
	} else {
		return DetermineFirstDigitYCoordinate(input, x, y-1)
	}
}

// parses the input into a matrix
func Parse(filepath string) [][]rune {
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

// Reads the contents of a file and returns its content.
// Will panic if file-reading fails for any reason
func ReadFile(fileInput string) string {
	fileContent, err := os.ReadFile(fileInput)

	if err != nil {
		log.Panic(err)
	}

	return string(fileContent)
}

// part-types
const (
	//Rune is not a part or next to a part
	NotPart int = 0
	//Rune is directly positioned next to a part
	NextToPart = 2
	Part       = 1
)

// Maps input matrix to a part-type matrix
func InputToNumericPartMap(filepath string, input [][]rune) [][]int {
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
				returnValue[x][y] = Part
				if x-1 < 0 == false {
					returnValue[x-1][y] = NextToPart
				}
				if x+1 > len(input) == false {
					returnValue[x+1][y] = NextToPart
				}
				if y-1 < 0 == false {
					returnValue[x][y-1] = NextToPart
				}
				if y+1 > len(input[x]) == false {
					returnValue[x][y+1] = NextToPart
				}
				if x+1 > len(input) == false && y+1 > len(input[x]) == false {
					returnValue[x+1][y+1] = NextToPart
				}
				if x-1 < 0 == false && y-1 < 0 == false {
					returnValue[x-1][y-1] = NextToPart
				}
				if x+1 > len(input) == false && y-1 < 0 == false {
					returnValue[x+1][y-1] = NextToPart
				}
				if x-1 < 0 == false && y+1 > len(input[x]) == false {
					returnValue[x-1][y+1] = NextToPart
				}
			}
		}
	}

	return returnValue
}

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	writer.Write([]byte("Введите наименование входного файла:\n"))
	writer.Flush()

	inputFile := readConsoleString(reader)
	fileLines := readFileLines(inputFile)

	writer.Write([]byte("Введите наименование файла для вывода результатов:\n"))
	writer.Flush()
	outputFile := readConsoleString(reader)
	fileWriter := fileWriterEmpty(outputFile)

	exampleRegexp := regexp.MustCompile(`(^[\d]+)([+\-\/*])([\d]+)=\?$`)
	for i := 0; i < len(fileLines); i++ {
		submatches := exampleRegexp.FindAllStringSubmatch(fileLines[i], -1)
		if len(submatches) == 0 {
			continue
		}
		writeBufferCalculated(fileWriter, submatches[0][1], submatches[0][3], submatches[0][2])
		if i%5 == 0 {
			fileWriter.Flush()
		}
	}

	fileWriter.Flush()
}

func readConsoleString(rd *bufio.Reader) string {
	str, err := rd.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(str)
}

func readFileLines(fileName string) []string {
	fileCont, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic("Входной файл не существует")
	}

	return strings.Split(string(fileCont), "\n")
}

func writeBufferCalculated(writer *bufio.Writer, val1Str string, val2Str string, operand string) {
	val1, err := strconv.Atoi(val1Str)
	if err != nil {
		panic(err)
	}

	val2, err := strconv.Atoi(val2Str)
	if err != nil {
		panic(err)
	}

	var resultVal string
	switch operand {
	case "+":
		resultVal = strconv.Itoa(val1 + val2)
	case "-":
		resultVal = strconv.Itoa(val1 - val2)
	case "*":
		resultVal = strconv.Itoa(val1 * val2)
	default:
		if val2 == 0 {
			panic("Деление на ноль")
		}
		resultVal = fmt.Sprintf("%.2f", float64(val1)/float64(val2))
	}

	resultStr := fmt.Sprintf("%s%s%s=%s\n", val1Str, operand, val2Str, resultVal)
	writer.Write([]byte(resultStr))
}

func fileWriterEmpty(fileName string) *bufio.Writer {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	return bufio.NewWriter(file)
}

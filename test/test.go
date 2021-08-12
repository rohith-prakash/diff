package test

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/rohith-prakash/diff/myer"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func equals(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func process(a []string, instructions []myer.Instruction) ([]string, error) {
	var b = make([]string, 0)
	length := len(instructions)
	i := 0
	j := 0
	var present myer.Instruction
	for j < length {
		present = instructions[j]
		if present.Decision == '=' {
			if a[i] != present.Data {
				return b, errors.New("String a does not aloow said instruction")
			}
			b = append(b, present.Data)
			i++
		} else if present.Decision == '+' {
			b = append(b, present.Data)
		} else if present.Decision == '-' {
			i++
		} else {
			return b, errors.New("String a does not aloow said instruction")
		}
		j++
	}
	return b, nil
}

func Tester(filepath1 string, filepath2 string) (bool, error) {
	file1, err := readLines(filepath1)
	if err != nil {
		fmt.Println(err)
		return false, errors.New("Unable to open file1")
	}
	file2, err := readLines(filepath2)
	if err != nil {
		fmt.Println(err)
		return false, errors.New("Unable to open file2")
	}
	instructions, err := myer.MyersDiff(file1, file2)
	if err == nil {
		b, err := process(file1, instructions)
		if err != nil {
			return false, err
		}
		myer.DisplayInstructions(instructions)
		fmt.Println(file1)
		fmt.Println(file2)
		fmt.Println(b)

		return equals(b, file2), nil
		//myer.DisplayInstructions(instructions)
	} else {
		//fmt.Println(err)
		myer.DisplayInstructions(instructions)
		return false, err
	}
}

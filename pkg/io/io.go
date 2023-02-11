package io

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Input() *string {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("[ERROR] : ", err)
	}

	input = strings.TrimSpace(input)

	return &input
}

func UserInput(prompt string) *string {
	fmt.Print(prompt)
	value := *Input()

	return &value
}

func CreateFile(where, data string) bool {
	err := ioutil.WriteFile(where, []byte(data), 0644)

	if err != nil {
		fmt.Println("my bad, there's an error when creating note")
		return false
	}

	return true
}

func Format(data ...string) string {
	var fs string
	for _, val := range data {
		fs += fmt.Sprintf("%v\n", val)
	}

	return fs
}

func FsExist(where string) bool {
	_, err := os.Stat(where)

	if os.IsNotExist(err) {
		return false
	}

	return true
}

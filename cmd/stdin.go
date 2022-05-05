package cmd

import (
	"bufio"
	"fmt"
	"os"
)

func GetStdin() (string, bool) {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return "", false
	}

	var stdin []byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		stdin = append(stdin, scanner.Bytes()...)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return "", false
	}
	return string(stdin), true
}

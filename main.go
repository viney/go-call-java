package main

import (
	"fmt"
	"os/exec"
)

func command(name string) (string, error) {
	out, err := exec.Command("/usr/bin/java", "Hello", name).Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func main() {
	say, err := command("viney")
	if err != nil {
		fmt.Println("command: ", err)
		return
	}
	fmt.Println(say)
}

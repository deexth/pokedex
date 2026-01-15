package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

)

func commandClear(conf *config) error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}
	if err := cmd.Run(); err != nil {
	fmt.Println(err)
	}
	return nil
}

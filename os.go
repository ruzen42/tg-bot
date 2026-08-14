package main

import "os/exec"

func fastfetch() (string, error) {
	cmd := exec.Command("neofetch", "--stdout")
	out, err := cmd.Output()
	return string(out), err
}

package command

import (
	"fmt"
	"os/exec"
)

func FortuneGet() string {
	stdout, err := exec.Command("fortune", "mythical-linux", "off/mythical-linux").Output()
	if err != nil {
		fmt.Println(err)
	}
	out := string(stdout[:])
	fmt.Println(out)
	return out
}

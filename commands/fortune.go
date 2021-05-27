package command

import (
	"log"
	"os/exec"
)

func FortuneGet() string {
	cmd := exec.Command("fortune", "mythical_linux off/mythical_linux")
	stout, err := cmd.Output()
	if err != nil {
		log.Panic(err)
	}
	out := string(stout)
	return out
}

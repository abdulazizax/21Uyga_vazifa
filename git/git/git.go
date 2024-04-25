package git

import (
	"os/exec"
)

func GetUserName() (string, error) {
	cmd := exec.Command("git", "config", "--get", "user.name")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func GetUserEmail() (string, error) {
	cmd := exec.Command("git", "config", "--get", "user.email")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

package shell

import (
	"fmt"
	"os/exec"
	"time"
)

const (
	BannerWriter = "banner-writer"
	ProjectPath  = "/opt/tmp/" + BannerWriter
)

type Shell struct {
	time time.Time
}

func New(time time.Time) *Shell {
	return &Shell{time: time}
}

func (sh *Shell) NextDay() {
	sh.time = sh.time.Add(24 * time.Hour)
}

func (sh *Shell) RmGit() error {
	err := exec.Command("rm", "-r", ".git").Run()
	if err != nil {
		return fmt.Errorf("failed exec rm -r .git: %v", err)
	}
	return nil
}
func (sh *Shell) SetGitGlobalConfig() error {
	err := exec.Command("git", "config", "--global", "user.email", "anton.t1984@gmail.com").Run()
	if err != nil {
		return fmt.Errorf("failed exec git config --global user.email: %v", err)
	}

	err = exec.Command("git", "config", "--global", "user.name", "Anton Trufanov").Run()
	if err != nil {
		return fmt.Errorf("failed exec git config --global user.name: %v", err)
	}

	return nil
}

func (sh *Shell) GitInit() error {
	dt := sh.time.Format(time.DateTime)
	err := exec.Command("faketime", "-f", "@"+dt+"", "git", "init").Run()
	if err != nil {
		return fmt.Errorf("failed exec git init: %v", err)
	}
	return nil
}

func (sh *Shell) GitCommit(text string) error {
	sh.time = sh.time.Add(time.Second)
	dt := sh.time.Format(time.DateTime)
	err := exec.Command("faketime", "-f", "@"+dt+"", "git", "add", ".").Run()
	if err != nil {
		return fmt.Errorf("failed exec git add . : %v", err)
	}

	err = exec.Command("faketime", "-f", "@"+dt+"", "git", "commit", "-m", text).Run()
	if err != nil {
		return fmt.Errorf("failed exec git commit: %v", err)
	}
	return nil
}

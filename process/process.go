package process

import (
	"fmt"
	"log"
	"os"

	"github.com/antonT001/banner-writer/shell"
	"github.com/antonT001/banner-writer/template"
	"github.com/antonT001/banner-writer/utils"
	"github.com/google/uuid"
	"github.com/schollz/progressbar/v3"
)

func Run(
	tmpl template.Template,
	sh *shell.Shell,
	bar *progressbar.ProgressBar,
) {
	// delete git history from directory
	if err := sh.RmGit(); err != nil {
		log.Fatal(err)
	}

	if err := sh.GitInit(); err != nil {
		log.Fatal(err)
	}

	if err := sh.SetGitGlobalConfig(); err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile("README.md", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("failed open file: %v", err)
	}
	defer f.Close()

	for j := 0; j < tmpl.Width(); j++ {
		for i := 0; i < tmpl.Height(); i++ {
			if tmpl[i][j] == template.COLORED_CELL {
				commits(f, sh, utils.RandInt(1, 5))
			}
			bar.Add(1)
			sh.NextDay()
		}
	}
}

func commits(f *os.File, sh *shell.Shell, n int) {
	for k := 0; k < n; k++ {
		msg := fmt.Sprintln(uuid.New().String())
		_, err := f.Write([]byte(msg))
		if err != nil {
			log.Fatal(err)
		}
		if err := sh.GitCommit(msg); err != nil {
			log.Fatal(err)
		}
	}
}

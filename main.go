package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/antonT001/banner-writer/process"
	"github.com/antonT001/banner-writer/shell"
	"github.com/antonT001/banner-writer/template"
	"github.com/schollz/progressbar/v3"
)

func main() {
	faketime := prepareDateTime(os.Getenv("YEAR"))
	text := os.Getenv("TEXT")
	tmpl := template.New()
	for _, r := range text {
		tmpl.SetDelimeter()
		tmpl.JoinObject(template.ByRune(r))
	}

	fmt.Printf("\n\nThis is how the banner will look on GitHub:\n")
	tmpl.Println()

	if !tmpl.Approve() {
		fmt.Println("Update the TEXT value in the file .env and repeat again")
		return
	}

	sh := shell.New(faketime)
	bar := progressbar.NewOptions(tmpl.Width()*tmpl.Height(),
		progressbar.OptionSetWidth(30),
		progressbar.OptionSetDescription("Creating new commits..."),
	)

	process.Run(tmpl, sh, bar)

	fmt.Println("====Success====")
}

func prepareDateTime(year string) time.Time {
	t, err := time.Parse(time.DateTime, year+"-01-01 15:04:05")
	if err != nil {
		log.Fatalf("failed parse year: %v", err)
	}
	if t.Year() >= time.Now().Year() {
		log.Fatalf("error, you specified the current or future year")
	}

	for {
		if t.Weekday() == time.Sunday {
			return t
		}
		t = t.Add(24 * time.Hour)
	}
}

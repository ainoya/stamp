package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "Stamp"
	app.Usage = "Add build id to iOS icon"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "in",
			Usage: "File path of source icon",
		},
		cli.StringFlag{
			Name:  "out",
			Usage: "File path of output image",
		},
	}

	app.Action = func(c *cli.Context) {
		m := NewMaterial(
			c.String("in"),
			c.String("out"))

		stamp(m)
	}

	app.Run(os.Args)
}

type material struct {
	inputPath  string
	outputPath string
	caption    string
	width      int
}

func NewMaterial(inputPath string, outputPath string) *material {
	material := &material{
		inputPath:  inputPath,
		outputPath: outputPath}

	material.decideWidth()
	material.makeCaption()

	return material
}

func stamp(m *material) {
	size := fmt.Sprintf("%dx40", m.width)

	cmd := exec.Command(
		"convert",
		"-background", "#0008",
		"-fill", "white",
		"-gravity", "center",
		"-size", size,
		m.caption,
		m.inputPath, "+swap",
		"-gravity", "north",
		"-composite", m.outputPath)
	cmd.Stderr = os.Stderr
	data, err := cmd.Output()

	if len(data) == 0 && err != nil {
		fmt.Printf("invoking convert: %v\n", err)
	}
}

func (m *material) decideWidth() {
	cmd := exec.Command(
		"identify",
		"-format", "%w",
		m.inputPath)

	cmd.Stderr = os.Stderr
	data, err := cmd.Output()

	if len(data) == 0 && err != nil {
		fmt.Printf("invoking indentify: %v\n", err)
		os.Exit(1)
	} else {
		m.width, _ = strconv.Atoi(string(data))
	}
}

func (m *material) makeCaption() {
	cmd := exec.Command(
		"git",
		"rev-parse", "--short", "HEAD")

	cmd.Stderr = os.Stderr
	data, err := cmd.Output()

	if len(data) == 0 && err != nil {
		fmt.Printf("invoking git-revparse: %v\n", err)
		os.Exit(1)
	} else {
		m.caption = fmt.Sprintf("caption:%s", strings.Trim(string(data), "\n"))
	}
}

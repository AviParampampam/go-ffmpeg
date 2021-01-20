package ffmpeg

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Input is a input file and options part of the ffmpeg command.
type Input struct {
	Options []string
	File    string
}

// InputGroup is a group input files.
type InputGroup *[]Input

// Output is a output file and options part of the ffmpeg command.
type Output struct {
	Options []string
	File    string
}

// OutputGroup is a group output files.
type OutputGroup *[]Output

// Command is a ffmpeg command.
type Command struct {
	FFmpegPath  string
	InputGroup  InputGroup
	OutputGroup OutputGroup
}

// NewCommand is a function for creating new ffmpeg command.
func (ffmpeg *FFmpeg) NewCommand(inputGroup InputGroup, outputGroup OutputGroup) *Command {
	return &Command{
		FFmpegPath:  ffmpeg.config.FFmpegPath,
		InputGroup:  inputGroup,
		OutputGroup: outputGroup,
	}
}

// String is a function for converting a command to a string.
func (cmd *Command) String() string {
	var commandString string

	commandString += cmd.FFmpegPath + " "

	for _, i := range *cmd.InputGroup {
		for _, j := range i.Options {
			commandString += j + " "
		}
		commandString += "-i " + i.File + " "
	}

	for _, i := range *cmd.OutputGroup {
		for _, j := range i.Options {
			commandString += j + " "
		}
		commandString += i.File + " "
	}

	return commandString
}

// Cmd is a function for creating new exec.Cmd.
func (cmd *Command) Cmd() *exec.Cmd {
	c := strings.Fields(cmd.String())

	return exec.Command(c[0], c[1:]...)
}

// Run starts the specified command and waits for it to complete.
func (cmd *Command) Run(stderr, stdout *os.File) {
	c := cmd.Cmd()

	c.Stderr = stderr
	c.Stdout = stdout

	c.Run()
}

// Start starts the specified command but does not wait for it to complete.
func (cmd *Command) Start(stderr, stdout *os.File) *exec.Cmd {
	c := cmd.Cmd()

	c.Stderr = stderr
	c.Stdout = stdout

	c.Start()

	for {
		fmt.Println(c.ProcessState)

	}

	return c
}

// NewInput is a function for creating new input file for ffmpeg command.
func NewInput(file string, options ...string) Input {
	return Input{
		Options: options,
		File:    file,
	}
}

// NewInputGroup is a function for creating input struct group.
func NewInputGroup(inputs ...Input) InputGroup {
	return &inputs
}

// NewOutput is a function for creating new output file for ffmpeg command.
func NewOutput(file string, options ...string) Output {
	return Output{
		Options: options,
		File:    file,
	}
}

// NewOutputGroup is a function for creating output struct group.
func NewOutputGroup(outputs ...Output) OutputGroup {
	return &outputs
}

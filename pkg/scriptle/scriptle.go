package scriptle

import (
	"fmt"
	"os"
	// "log"
	// "os"
	// "strings"

	"github.com/adamazing/scriptle/pkg/model"
	util "github.com/adamazing/scriptle/pkg/utility"
	tea "github.com/charmbracelet/bubbletea"
	wrap "github.com/mitchellh/go-wordwrap"
)

const (
	blue = "#4776E6"
)

type Flags struct {
  Length string
}

func FromStdin(n int, flagStruct *Flags) error {
  var stdin []byte
  scanner := bufio.NewScanner(os.Stdin)

  return run(text)
}

func FromFile(path string, flagStruct *Flags) error {
	text, err := util.ReadFile(path)
	if err != nil {
		return err
	} else {
    fmt.Println(text)
  }

	text, err = flagStruct.doStuffWithFlags()
	if err != nil {
		return err
	}

	return run(text)
}

func (f *Flags) doStuffWithFlags() (string, error) {
  // var err error
  return "", nil
}

func run(text string) error {
	program := tea.NewProgram(model.Model{
		Text: wrap.WrapString("ASNDJKASNDKJAS", 80),
	})

	return program.Start()
}

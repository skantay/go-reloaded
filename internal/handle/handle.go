package reloaded

import (
	"errors"
	"fmt"
	"os"
)

func Run() error {
	text := new(Text)

	if err := text.getFiles(); err != nil {
		return fmt.Errorf("run error | getFiles: %w", err)
	}

	if err := text.getText(); err != nil {
		return fmt.Errorf("run error | getText: %w", err)
	}

	text.getFlags()
	fmt.Println(text.Text)
	return nil
}

// Main object
type Text struct {
	fileFrom string
	fileTo   string
	Text     string
	flags    []string
}

// Object
func (t *Text) getFlags() {
	text := t.Text
	brackets := []byte{}
	var flagIndex int
	for i := 0; i < len(text); i++ {
		if text[i] == '(' {
			brackets = append(brackets, text[i])
			for j := i; i < len(text); j++ {
				if text[j] == ')' {
					brackets = brackets[:len(brackets)-2]
				} else if text[j] == '(' {
					brackets = append(brackets, text[i])
					flagIndex = j + 1
				}
				if len(brackets) == 0 {
					t.flags = append(t.flags, flagsHelper(flagIndex, j, text))
					i = j
					break
				}
			}
		}
	}
}

// Helper object method - gets the flags
func flagsHelper(flagIndex, j int, text string) string {
	result := ""
	for ; flagIndex < j; flagIndex++ {
		result += string(text[flagIndex])
	}
	return result
}

// Object method - gets only 2 file paths
func (t *Text) getFiles() error {
	args := os.Args[1:]

	if len(args) != 2 {
		return errors.New("enter 2 file paths")
	}

	t.fileFrom = args[0]
	t.fileTo = args[1]

	return nil
}

// Object method - gets text by filepath
func (t *Text) getText() error {
	data, err := os.ReadFile(t.fileFrom)
	if err != nil {
		return errors.New("no such file")
	}

	t.Text = string(data)

	return nil
}

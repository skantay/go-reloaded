package reloaded

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	tool "go-reloaded/internal/handle/functions"
)

// Main object
type text struct {
	fileFrom string
	fileTo   string
	text     string
	flags    []string
}

// Object method - gets only 2 file paths
func (t *text) getFiles() error {
	args := os.Args[1:]

	if len(args) != 2 {
		return errors.New("enter 2 file paths")
	}

	t.fileFrom = args[0]
	t.fileTo = args[1]

	return nil
}

// Object method - validate files
func (t *text) validateFiles() error {
	first := filepath.Ext(t.fileFrom)
	second := filepath.Ext(t.fileTo)
	errorText := "file extension is invalid. Provide ONLY .txt files"

	if first != ".txt" && first != "" {
		return fmt.Errorf("input, %s", errorText)
	} else if second != ".txt" && second != "" {
		return fmt.Errorf("output %s", errorText)
	}

	return nil
}

// Object method - gets text by filepath
func (t *text) getText() error {
	data, err := os.ReadFile(t.fileFrom)
	if err != nil {
		return fmt.Errorf("%s no such file", t.fileFrom)
	}

	t.text = string(data)

	return nil
}

// Object method - gets flags in right order
func (t *text) getFlags() {
	re := regexp.MustCompile(`\([\s\n]*([Hh][Ee][Xx]|[Uu][Pp]|[Ll][Oo][Ww]|[Bb][Ii][Nn]|[Cc][Aa][Pp])\s*\)|\([\s\n]*([Uu][Pp]|[Ll][Oo][Ww]|[Cc][Aa][Pp])[\s\n]*,[\s\n]*\d*[\s\n]*\)`)

	t.flags = re.FindAllString(t.text, len(t.text))
}

// Object method - process and apply flags
func (t *text) processFlags() {
	t.getFlags()
	flags := t.flags
	for indexFlag := 0; indexFlag < len(flags); indexFlag++ {
		flag, _ := trim(flags[indexFlag])
		fmt.Println(indexFlag, flag)
		fmt.Println(t.text)
		fmt.Println("-------------")
		switch flag {
		case "up":
			t.text = tool.Up(t.text)
		case "low":
			t.text = tool.Low(t.text)
		case "cap":
			t.text = tool.Cap(t.text)
		case "bin":
			t.text = tool.Bin(t.text)
		case "hex":
			t.text = tool.Hex(t.text)
		}
		fmt.Println(t.text)
		fmt.Println("-------------------------------------------------")
	}
}

// Helper of processFlags()
func trim(flag string) (string, string) {
	flag = strings.ReplaceAll(flag, " ", "")
	flag = strings.ReplaceAll(flag, "(", "")
	flag = strings.ReplaceAll(flag, ")", "")
	flag = strings.ReplaceAll(flag, ",", "")
	flag = strings.ReplaceAll(flag, "\n", "")

	numFlag := flag
	for i := 'a'; i <= 'z'; i++ {
		numFlag = strings.ReplaceAll(numFlag, string(i), "")
	}

	for i := 'A'; i <= 'Z'; i++ {
		numFlag = strings.ReplaceAll(numFlag, string(i), "")
	}

	for i := '0'; i <= '9'; i++ {
		flag = strings.ReplaceAll(flag, string(i), "")
	}

	flag = strings.ToLower(flag)
	return flag, numFlag
}

// Objecet - method - apply other changes
func (t *text) other() {
	t.text = tool.Punct(t.text)
	t.text = tool.Quotes(t.text)
	t.text = tool.Articles(t.text)
}

// Object method - puts the t.text to a file
func (t *text) putFile() {
	result, _ := os.Create(t.fileTo)
	defer result.Close()

	buffer := []rune(t.text)

	for _, line := range buffer {
		result.WriteString(string(line))
	}
}

// Helper of Error()
func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

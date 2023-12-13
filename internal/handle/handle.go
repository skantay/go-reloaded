package reloaded

import (
	"errors"
	"fmt"
)

var (
	errGetFiles = errors.New("err: run() ---> getFiles()")
	errGetText  = errors.New("err: run() ---> getText()")
	errFileName = errors.New("err: run() ---> validateFiles()")
	errPutFile  = errors.New("err: run() ---> putFile()")
)

func Run() error {
	text := new(text)

	if err := text.getFiles(); err != nil {
		return fmt.Errorf("%s: %w", errGetFiles, err)
	}

	if err := text.validateFiles(); err != nil {
		return fmt.Errorf("%s: %w", errGetText, err)
	}

	if err := text.getText(); err != nil {
		return fmt.Errorf("%s: %w", errFileName, err)
	}

	text.other()

	for i := 0; i < 10; i++ {
		text.processFlags()
	}

	text.other()

	if err := text.putFile(); err != nil {
		return fmt.Errorf("%s: %w", errPutFile, err)
	}

	return nil
}

func Error() error {
	err1 := errors.New("did you README")
	err2 := errors.New("hmmmmm, I think you did not README")
	err3 := errors.New("why would you skip my README")
	err4 := errors.New("bruh, something is wrong")
	err5 := errors.New("i want an ice cream")

	errors := []error{
		err1,
		err2,
		err3,
		err4,
		err5,
	}

	return errors[randInt(0, len(errors))]
}

package reloaded

import (
	"errors"
	"fmt"
)

var (
	errGetFiles = errors.New("err: run() ---> getFiles()")
	errGetText  = errors.New("err: run() ---> getText()")
	errFileName = errors.New("err: run() ---> validateFiles()")
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
	text.putFile()

	return nil
}

func Error() error {
	err1 := errors.New("Did you read README ???????")
	err2 := errors.New("Hmmmmm, I think you did not read README !")
	err3 := errors.New("Why would you skip my README (⁠⊙⁠_⁠◎⁠)")
	err4 := errors.New("bruh, something is wrong")
	err5 := errors.New("I think The weather causes errors ;(")

	errors := []error{
		err1,
		err2,
		err3,
		err4,
		err5,
	}

	return errors[randInt(0, len(errors))]
}

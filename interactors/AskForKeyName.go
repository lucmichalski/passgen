package interactors

import (
	"errors"
	"os"
	"unicode/utf8"

	"gopkg.in/AlecAivazis/survey.v1"
)

// AskForKeyName is asks for the key name for the password you are going to generate
func AskForKeyName() string {
	name := ""
	prompt := &survey.Input{
		Message: "What is key name for password :",
		Help:    "Key name should be at least 3 characters long",
	}

	err := survey.AskOne(prompt, &name, func(val interface{}) error {
		if utf8.RuneCountInString(val.(string)) < 3 {
			return errors.New("value is too short. Min length is 3")
		}
		return nil
	})
	if err != nil {
		os.Exit(1)
	}

	return name
}

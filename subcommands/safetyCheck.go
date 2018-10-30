package subcommands

import (
	"errors"
	"os"
	"passgen/helpers"
	"passgen/safety"

	"github.com/AlecAivazis/survey"
	"github.com/spf13/cobra"
)

// safetyCheck subcommand variable
var safetyCheck = &cobra.Command{
	Use:   "check",
	Short: "Safety check for passwords",
	Long:  "Couple of security checks for your password's safety",
	Run:   safetyCheckRunner,
}

// safetyCheckRunner runs couple of security checks for your password's safety
func safetyCheckRunner(cmd *cobra.Command, args []string) {
	password := ""
	prompt := &survey.Password{
		Message: "What is your password :",
		Help:    "Write any of your current password for security check",
	}

	err := survey.AskOne(prompt, &password, func(val interface{}) error {
		if helpers.ProperCharacterCounter(val.(string)) < 1 {
			return errors.New("value is too short. Min length is 1")
		}
		return nil
	})
	if err != nil {
		os.Exit(1)
	}

	safety.Check(password)
}

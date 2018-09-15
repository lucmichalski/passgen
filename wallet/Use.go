package wallet

import (
	"errors"
	"log"

	"github.com/ozguncagri/passgen/generators"

	"github.com/ozguncagri/passgen/interactors"

	"github.com/ozguncagri/passgen/helpers"

	"gopkg.in/AlecAivazis/survey.v1"
)

// Use is used for generating password using settings defined in wallet item
func Use(memoryWallet *PassgenWallet) {
	if len(memoryWallet.Wallet) == 0 {
		helpers.NegativePrintf("\nThere is no item in your wallet\n\n")
		return
	}

	walletItemKey := ""
	var allKeys []string

	for keys := range memoryWallet.Wallet {
		allKeys = append(allKeys, keys)
	}

	prompt := &survey.Select{
		Message: "Choose wallet item for generate password :",
		Help:    "Pick any wallet item for generate password with pre defined settings",
		Options: allKeys,
	}
	err := survey.AskOne(prompt, &walletItemKey, func(val interface{}) error {
		if helpers.ProperCharacterCounter(val.(string)) == 0 {
			return errors.New("this is not valid selection")
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	masterPassword := interactors.AskForPassword()
	generatedPassword := generators.GeneratePassword(
		memoryWallet.Wallet[walletItemKey].Pool,
		walletItemKey,
		masterPassword,
		memoryWallet.Wallet[walletItemKey].Length,
	)

	helpers.ResultPrintf("\nYour password is : %v\n\n", generatedPassword)
}

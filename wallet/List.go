package wallet

import (
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/ozguncagri/passgen/helpers"
)

// List subcommand for wallet lister
func List(memoryWallet *PassgenWallet) {
	if len(memoryWallet.Wallet) == 0 {
		helpers.NegativePrintf("\nThere is no item in your wallet\n\n")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Key Name", "Character Pool", "Length"})

	for key, v := range memoryWallet.Wallet {
		seperatedPool := strings.Split(v.Pool, "")

		for i := range seperatedPool {
			switch seperatedPool[i] {
			case "U":
				seperatedPool[i] = "Upper"
			case "L":
				seperatedPool[i] = "Lower"
			case "N":
				seperatedPool[i] = "Number"
			case "S":
				seperatedPool[i] = "Symbol"
			}
		}

		table.Append([]string{key, strings.Join(seperatedPool, ", "), fmt.Sprint(v.Length)})
	}
	table.Render()
}

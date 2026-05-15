package topfeel

import (
	"fmt"

	"github.com/spf13/cobra"
)

var signInCmd = &cobra.Command{
	Use:   "sign-in",
	Short: "自动签到",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("OK")
	},
}

func init() {
	TopfeelCmd.AddCommand(signInCmd)
}

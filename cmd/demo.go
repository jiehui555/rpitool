package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "测试",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("OK")
	},
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "测试命令",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("OK")
	},
}

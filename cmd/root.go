package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rpitool",
	Short: "一个多功能命令行工具",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("欢迎使用 rpitool！使用 --help 查看帮助。")
	},
}

// Execute 将所有子命令加入根命令并执行
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

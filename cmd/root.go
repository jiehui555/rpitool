package cmd

import (
	"fmt"
	"os"

	"github.com/jiehui555/rpitool/cmd/topfeel"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	envFile string
)

var rootCmd = &cobra.Command{
	Use:   "rpitool",
	Short: "一个多功能命令行工具",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if envFile != "" {
			if err := godotenv.Load(envFile); err != nil {
				return fmt.Errorf("无法加载环境变量文件 %s: %w", envFile, err)
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("欢迎使用 rpitool！使用 --help 查看帮助。")
	},
}

func init() {
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().StringVarP(&envFile, "env", "e", ".env", "指定环境变量文件")

	rootCmd.AddCommand(topfeel.TopfeelCmd)
	rootCmd.AddCommand(demoCmd)
}

// Execute 将所有子命令加入根命令并执行
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

package cmd

import (
	"fmt"
	"os"
	"strings"

	huaweisms "github.com/jiehui555/rpitool/pkg/huawei_sms"
	"github.com/spf13/cobra"
)

var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "测试华为路由器短信获取功能",
	Long:  `测试命令：从华为路由器获取收件箱短信列表并显示。`,
	Run: func(cmd *cobra.Command, args []string) {
		// 创建客户端
		client := huaweisms.NewHuaweiClient(
			os.Getenv("HUAWEI_SMS_URL"),
			os.Getenv("HUAWEI_SMS_USER"),
			os.Getenv("HUAWEI_SMS_PASSWD"),
		)

		fmt.Println("正在连接路由器并获取短信...")

		// 获取短信列表
		messages, err := client.GetSMSList()
		if err != nil {
			fmt.Printf("获取短信失败: %v\n", err)
			return
		}

		if len(messages) == 0 {
			fmt.Println("当前收件箱没有短信。")
			return
		}

		// 输出短信内容
		fmt.Printf("成功获取 %d 条短信：\n\n", len(messages))

		for i, msg := range messages {
			fmt.Printf("【%d】\n", i+1)
			fmt.Printf("号码: %s\n", msg.Phone)
			fmt.Printf("时间: %s\n", msg.Date)
			fmt.Printf("内容: %s\n", msg.Content)
			fmt.Println(strings.Repeat("-", 60))
		}
	},
}

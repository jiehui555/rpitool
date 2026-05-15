package topfeel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var signInCmd = &cobra.Command{
	Use:   "sign-in",
	Short: "自动签到",
	Run: func(cmd *cobra.Command, args []string) {
		result := executeSignIn()
		if result.Success {
			fmt.Printf("✅ 签到成功: %s\n", result.Message)
		} else {
			fmt.Printf("❌ 签到失败: %s\n", result.Message)
		}
	},
}

type SignInResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// 执行签到逻辑
func executeSignIn() SignInResult {
	// 生成时间戳（毫秒）
	now := time.Now().UnixMilli()
	newTime := now + int64(rand.Intn(4)+3)*1000 // 随机 3~6 秒

	body := map[string]interface{}{
		"oldtime": now,
		"newtime": newTime,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return SignInResult{Success: false, Message: "JSON 序列化失败: " + err.Error()}
	}

	topfeelSignInURL := "https://bbs.topfeel.com/api/gift/day_sign"
	req, err := http.NewRequest("POST", topfeelSignInURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return SignInResult{Success: false, Message: "创建请求失败: " + err.Error()}
	}

	// 设置请求头
	req.Header.Set("Referer", "https://bbs.topfeel.com/h5/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="143", "Chromium";v="143", "Not A(Brand";v="24"`)
	req.Header.Set("token", os.Getenv("TOPFEEL_TOKEN"))

	// 发送请求
	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return SignInResult{Success: false, Message: "网络请求失败: " + err.Error()}
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return SignInResult{Success: false, Message: "读取响应失败: " + err.Error()}
	}

	// 非 200 状态码处理
	if resp.StatusCode != http.StatusOK {
		return SignInResult{
			Success: false,
			Message: fmt.Sprintf("HTTP %d: %s", resp.StatusCode, string(respBody)),
		}
	}

	// 解析响应
	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return SignInResult{Success: false, Message: "JSON 解析失败: " + err.Error()}
	}

	// 输出原始响应（便于调试）
	fmt.Printf("📡 服务器响应: %s\n", string(respBody))

	code := 0
	if c, ok := result["code"].(float64); ok {
		code = int(c)
	}

	msg := ""
	if m, ok := result["msg"].(string); ok {
		msg = m
	}

	if code != 0 {
		return SignInResult{Success: false, Message: msg}
	}

	return SignInResult{Success: true, Message: msg}
}

func init() {
	TopfeelCmd.AddCommand(signInCmd)
}

package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
)

const tgAPIToken = "666666666:xxxxxxxxxxxxx"
const tgUserID = "xxxxxxxxx"

func sendToTelegram(message string, wg *sync.WaitGroup) {
	defer wg.Done()

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", tgAPIToken)
	data := url.Values{}
	data.Set("chat_id", tgUserID)
	data.Set("text", message)

	// 忽略所有错误
	_, _ = http.PostForm(apiURL, data)
}

func main() {
	// 从命令行参数中获取输入
	if len(os.Args) < 2 {
		os.Exit(0)
	}
	input := strings.Join(os.Args[1:], " ")

	var wg sync.WaitGroup
	wg.Add(1)

	// 使用goroutine异步发送到Telegram
	go sendToTelegram(input, &wg)

	// 等待所有的goroutine完成
	wg.Wait()
}

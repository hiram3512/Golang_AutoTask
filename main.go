/*
每天凌晨1点自动更新svn
每天凌晨2点自动载入unity
*/

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/axgle/mahonia"
	"github.com/robfig/cron"
)

var (
	unityPath, projectPath string
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("输入Unity安装路径:")
	text, _ := reader.ReadString('\n')
	unityPath = strings.Replace(text, "\r\n", "", -1) + "/Unity.exe"
	fmt.Printf("输入项目路径:")
	text, _ = reader.ReadString('\n')
	projectPath = strings.Replace(text, "\r\n", "", -1)
	fmt.Println("--------------------------------------\n开始执行定时任务\n每天凌晨1点自动更新svn\n每天凌晨3点自动载入unity\n--------------------------------------")

	c := cron.New()
	c.AddFunc("0 0 3 * * ?", func() {
		//强制关闭unity
		taskkill := exec.Command("taskkill", "/f", "/im", "unity.exe")
		taskkill.Run()

		fmt.Println(time.Now(), "自动更新Svn")
		cmdCleanup := exec.Command("svn", "cleanup", projectPath)
		cmdCleanup.Run()
		cmdUpdate := exec.Command("svn", "update", projectPath, "--accept", "postpone")
		cmdUpdate.Run()
	})
	c.Start()

	c2 := cron.New()
	c2.AddFunc("0 0 3 * * ?", func() {
		fmt.Println(time.Now(), "自动载入Unity")
		out := bytes.NewBuffer(nil)
		cmd := exec.Command(unityPath, "-projectPath", projectPath)
		cmd.Stdout = out
		cmd.Run()
		enc := mahonia.NewDecoder("gb18030")
		goStr := enc.ConvertString(out.String())
		fmt.Println(goStr)
	})
	c2.Start()

	fmt.Scanln()
}

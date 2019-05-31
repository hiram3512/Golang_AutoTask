/*
每天凌晨1点自动更新svn
每天凌晨2点自动载入unity
*/

package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"

	"github.com/axgle/mahonia"
	"github.com/robfig/cron"
)

var (
	unityPath, projectPath string
)

func main() {
	fmt.Printf("输入Unity安装路径:")
	fmt.Scanln(&unityPath)
	unityPath = unityPath + "/Unity.exe"
	fmt.Printf("输入项目路径:")
	fmt.Scanln(&projectPath)
	fmt.Println("开始执行定时任务")
	fmt.Println("每天凌晨1点自动更新svn\n每天凌晨2点自动载入unity")
	c := cron.New()
	c.AddFunc("0 20 11 * * ?", func() {
		fmt.Println(time.Now(), "自动更新Svn")
		out := bytes.NewBuffer(nil)
		cmd := exec.Command("svn", "update", "D:/MySvn/QinUI")
		cmd.Stdout = out
		cmd.Run()

		enc := mahonia.NewDecoder("gb18030")
		goStr := enc.ConvertString(out.String())
		fmt.Println(goStr)
	})
	c.Start()

	c2 := cron.New()
	c2.AddFunc("0 22 11 * * ?", func() {
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

/*
每天凌晨1点自动更新svn
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

func main() {
	// ticker := time.NewTicker(time.Minute * 10)
	// for _ = range ticker.C {
	// 	fmt.Println(time.Now(), "更新Svn")
	// 	cmd := exec.Command("svn", "update", "D:/MySvn/QinUI")
	// 	cmd.Run()
	// }

	c := cron.New()
	//c.AddFunc("0 0 1 * * ?", func() {
	c.AddFunc("0 0 21 27 * ?", func() {
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
}

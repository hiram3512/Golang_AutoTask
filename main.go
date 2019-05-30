/*
每10分钟自动更新Svn
*/

package main

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/axgle/mahonia"
)

func main() {
	//path := "C:/Users/tanxiaoliang.hiram/AppData/Local/Google/Chrome/Application/chrome.exe"

	out := bytes.NewBuffer(nil)
	cmd := exec.Command("svn", "update", "D:/MySvn/QinUI")
	cmd.Stdout = out
	cmd.Run()

	enc := mahonia.NewDecoder("gb18030")
	goStr := enc.ConvertString(out.String())
	fmt.Println(goStr)

	// ticker := time.NewTicker(time.Minute * 10)
	// for _ = range ticker.C {
	// 	fmt.Println(time.Now(), "更新Svn")
	// 	cmd := exec.Command("svn", "update", "D:/MySvn/QinUI")
	// 	cmd.Run()
	// }
}

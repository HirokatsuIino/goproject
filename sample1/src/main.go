package main

import (
	"fmt"
	"github.com/carlescere/scheduler"
	"runtime"
)
import _ "github.com/carlescere/scheduler"
import _ "runtime"

func main() {
	// 5秒に1回出力させる
	scheduler.Every(5).Seconds().Run(printSuccess)
	runtime.Goexit()
	//fmt.Print("hello Wprld")

}

func printSuccess(){
	fmt.Print("success \n")
}
// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/xiaobo9/go-gui/gui"
	"github.com/xiaobo9/go-gui/toastMsg"
)

// init log
func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// type Action func()
// var mapFuncs map[string]Action

var actionMap map[string]func()

// init name to function map
func init() {
	// var mapFuncs = make(map[string]Action)
	// var mapFuncs = make(map[string]func())

	actionMap = map[string]func(){
		"toast": func() { toastMsg.Notification() },
	}

	actionMap["gui"] = func() { gui.Window() }
}

func main() {
	if len(os.Args) <= 1 {
		log.Println("没有指定动作")
		return
	}
	actionName := os.Args[1]
	log.Println("action name: ", actionName)
	action := actionMap[actionName]
	if action != nil {
		(action)()
		return
	}
	log.Println("没有找到指定动作: ", actionName)
}

func gcd(x, y int) int {
	for y != 0 {
		fmt.Printf("%d %d %d\n", x, y, x%y)
		x, y = y, x%y
	}
	return x
}

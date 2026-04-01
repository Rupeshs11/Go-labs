package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
	"runtime"
	"syscall"
	"unsafe"
)

const(
	markName = "GOLANG_CLI_REMINDER"
	markValue = "1"
)

func main(){
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <time> <message>\n", os.Args[0])
		os.Exit(1)
	}
	now := time.Now()

	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	t,err := w.Parse(os.Args[1],now)

	if err != nil{
		fmt.Println(err)
		os.Exit(2)
	}

	if t == nil {
		fmt.Println("Unable to parse time!")
		os.Exit(2)
	}
	if now.After(t.Time){
		fmt.Println("set a future time!")
		os.Exit(3)
	}
	diff := t.Time.Sub(now)
	if os.Getenv(markName) == markValue {
		time.Sleep(diff)
		message := strings.Join(os.Args[2:], " ")
		beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)

		if runtime.GOOS == "windows" {
			user32 := syscall.NewLazyDLL("user32.dll")
			messageBox := user32.NewProc("MessageBoxW")
			titlePtr, _ := syscall.UTF16PtrFromString("Reminder")
			msgPtr, _ := syscall.UTF16PtrFromString(message)
			messageBox.Call(0, uintptr(unsafe.Pointer(msgPtr)), uintptr(unsafe.Pointer(titlePtr)), 0x00010040)
		} else {
			err := beeep.Alert("Reminder", message, "")
			if err != nil {
				fmt.Println(err)
				os.Exit(4)
			}
		}
	} else {
		cmd := exec.Command(os.Args[0], os.Args[1:]...)
		cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", markName, markValue))
		if err := cmd.Start(); err != nil {
			fmt.Println(err)
			os.Exit(5)
		}
		fmt.Println("Reminder will be displayed after",diff.Round(time.Second))
		os.Exit(0)
	}
}


package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var (
	raspivid *exec.Cmd
)

func startRecord(command exec.Cmd, output string) {
	starterr := command.Start()
	if starterr != nil {
		fmt.Println("Can't start raspivid")
		fmt.Println(starterr.Error())
	}

}

func stopRecord(command exec.Cmd) {
	killerr := command.Process.Kill()
	if killerr != nil {
		fmt.Println("Can't kill raspivid")
		fmt.Println(killerr)
	}
}

func init() {
}

func main() {
	fmt.Println(os.Args[1])
	raspivid = exec.Command("raspivid", "-o", os.Args[1]+".h264", "-t", "1000000000")
	startRecord(*raspivid, "video2.h264")
	time.Sleep(10)
	stopRecord(*raspivid)

}

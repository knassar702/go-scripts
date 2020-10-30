package main 
// by Mohamed Dief @DEMON1A
import (
	"fmt"
	"os/exec"
	"runtime"
	"flag"
)

func main() {
	OS := runtime.GOOS
	if OS != "linux" {
		fmt.Println("This Simple Shell Only Works With Linux")
		fmt.Println("You Current OS:" , OS)
	}

	Command := flag.String("cmd" , "" , "The Command You Want To Exec On Your Machine")
	flag.Parse()

	if *Command == "" {
		fmt.Println("You Didn't Provide The Command.")
	} else {
		Output , error := exec.Command(*Command).Output()
		if error != nil {
			fmt.Println(error)
		}

		CMDOutput := string(Output[:])
		fmt.Println("The Command Has Been Executed. Here's Your Results:")
		fmt.Println(CMDOutput)
	}
}

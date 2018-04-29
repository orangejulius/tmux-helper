package main

import "fmt"
import "os"
import "os/exec"
import "strings"
import "syscall"

func main() {
	env := os.Environ()
	fmt.Println("hello world")

	fmt.Println(os.Args)

	binary, lookErr := exec.LookPath("tmux")
	if lookErr != nil {
		panic(lookErr)
	}

	if len(os.Args) == 1 {
		pwd, _ := os.Getwd()

		directories := strings.Split(pwd, "/")
		dLen := len(directories)

		if directories[dLen-3] == "repos" {
			args := []string{"tmux", "new-session", "-s", directories[dLen-2] + "/" + directories[dLen-1]}
			syscall.Exec(binary, args, env)
		} else {
			args := []string{"tmux", "attach"}
			syscall.Exec(binary, args, env)
		}
	} else {
		args := []string{"tmux", "new-session", "-s", os.Args[1]}
		syscall.Exec(binary, args, env)
	}
}

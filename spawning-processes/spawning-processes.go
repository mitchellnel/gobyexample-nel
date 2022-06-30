package main

import (
	"fmt"
	"io"
	"os/exec"
)

// Sometimes our Go programs need to spawn other, non-Go processes

func main() {
	// we'll start with a simple command that takes no arguments or input, and just prints something
	// to stdout

	// exec.Command helper creates an object to represent this external process
	dateCmd := exec.Command("date")

	// the Output method runs the command, waits for it to finish, and collects its standard output
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(">date")
	fmt.Println(string(dateOut))
	// if there were not errors, dateOut will hold bytes with the date info

	// Output and other methods of Command will return *exec.Error if there was a problem executing
	// the command (e.g. wrong path), and *exec.ExitError if the command ran but exited with a
	// nonzero return code
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}

	// next we'll look at a slightly more involved case where we pipe date to the external process
	// on its stdin, and collect the results from stdout
	grepCmd := exec.Command("grep", "hello")

	// here, we explicitly grab input/output pipes, start the process, write some input to it, read
	// the resulting output, and finally wait for the process to exit
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// we omitted error checks in the above example, but use the usual if err != nil pattern
	// we also only collect the StdoutPipe results, but you could collect the StderrPipe in exactly
	// the same way

	// note that when spawning commands we need to provide an explicitly delinated command and
	// argument array vs just being able to pass in one command line string

	// if we want to just pass in one command line string, we can use bash's -c option
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {
	defer log.Println("Exiting main application")

	setpgid := flag.Bool("new-group", false, "Specifies whether child process should be in a new process group")
	flag.Parse()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	cmd := exec.Command("./system/signal/assets/trapper.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	// Force child process into their own process group, instead of being in the same process group as the main process.
	//
	// "In a POSIX-conformant operating system, a process group denotes a collection of one or more processes.
	//  Among other things, a process group is used to control the distribution of a signal; when a signal is directed to a process group,
	//  the signal is delivered to each process that is a member of the group.(..)"
	//
	// With a new group, we can preserve the running child process on SIGINT and SIGTERM sent to main process.
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: *setpgid,
		Pgid:    0,
	}
	exitOnErr(cmd.Start())

	cmdDone := make(chan struct{})
	go func() {
		exitOnErr(cmd.Wait())
		close(cmdDone)
	}()

	id, err := syscall.Getpgid(os.Getpid())
	exitOnErr(err)
	log.Printf("Main process PID:%d in group %d\n", os.Getpid(), id)

	id, err = syscall.Getpgid(cmd.Process.Pid)
	exitOnErr(err)
	log.Printf("Child process PID %d in group %d\n", cmd.Process.Pid, id)

	select {
	case <-sigCh:
		os.Exit(2)
	case <-cmdDone:
		log.Println("Child process finished, exiting with child exit code")
		os.Exit(cmd.ProcessState.ExitCode())
	}
}

func exitOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

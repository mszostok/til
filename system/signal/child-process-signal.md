## Send signal to a child process

By default, when the `exec.Cmd` is used, the child process is assigned the main proces group. As a result, it receives signal sent to the main process too.

> In a POSIX-conformant operating system, a process group denotes a collection of one or more processes.
> Among other things, a process group is used to control the distribution of a signal; when a signal is directed to a process group,
> the signal is delivered to each process that is a member of the group.
<!--more-->

We can change that behavior with `setpgid` syscall:

```go
cmd.SysProcAttr = &syscall.SysProcAttr{
	Setpgid: true,
	Pgid:    0,
}
```

Other option, is to make the child process aware of a OS signal:

```bash
_term() {
  echo "child: Ignoring signal for 30 sec"
  sleep 30
}

trap _term SIGTERM SIGINT
```

## Testing

1. Same group - signal received:

   ```bash
   go run system/signal/assets/child-proc-sig.go -new-group
   # press ctrl-c
   ```

   Output:

   ```log
   2022/04/03 01:17:45 Main process PID:39465 in group 39442
   2022/04/03 01:17:45 Child process PID 39468 in group 39442
   Starting child process - sleep 60s
   ^Cchild: received signal, ignoring it for 30 sec
   exit status 2
   ```

   As you can see, child process received signal, and decided to ignore it fo 30sec. If the child did not pick up the signal, it would be shut down too.

2. New child group - signal not received:

   ```bash
   go run system/signal/assets/child-proc-sig.go
   # press ctrl-c
   ```

   Output:

   ```log
   2022/04/03 01:17:56 Main process PID:39537 in group 39515
   2022/04/03 01:17:56 Child process PID 39538 in group 39538
   Starting child process - sleep 60s
   ^Cexit status 2
   ```

   As you can see, child process doesn't receive signal, and it's still running. It will be even if did not handle OS signals.
